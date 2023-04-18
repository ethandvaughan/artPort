package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// database connection
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "arfol"
)

// piece represents data about a record piece.
type Piece struct {
	ID                uuid.UUID      `json:"id"`
	Title             string         `json:"title"`
	Artist            string         `json:"artist"`
	Category          string         `json:"category"`
	Clay              sql.NullString `json:"clay"`
	Bisque_Cone       sql.NullString `json:"bisque_cone"`
	Glaze_Description sql.NullString `json:"glaze_description"`
	Glaze_Cone        sql.NullString `json:"glaze_cone"`
	Size              sql.NullString `json:"size"`
	Date              time.Time      `json:"date"`
	Description       sql.NullString `json:"description"`
	Images            []string       `json:"images"`
}

// main function to handle the routing of CRUD actions
func main() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTPASS"), user, host, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT current_database()")
	if err != nil {
		// handle error
	}
	defer rows.Close()

	var dbName string
	for rows.Next() {
		err := rows.Scan(&dbName)
		if err != nil {
			// handle error
		}
	}

	if err := rows.Err(); err != nil {
		// handle error
	}

	fmt.Println("Connected to database:", dbName)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		log.Fatalf("Error connecting to AWS: %v", err)
	}

	uploader := s3manager.NewUploader(sess)

	router := gin.Default()

	router.Use(cors())

	router.GET("/pieces", func(c *gin.Context) {
		pieces, err := getPieces(db)
		if err != nil {
			log.Fatalf("Error querying database: %v", err)
		}
		json, err := json.Marshal(pieces)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusOK, "application/json", json)
	})

	router.GET("/pieces/:id", func(c *gin.Context) {
		id := c.Param("id")
		pieces, err := getPieceByID(db, id)
		if err != nil {
			log.Fatalf("Error querying database: %v", err)
		}
		json, err := json.Marshal(pieces)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusOK, "application/json", json)

	})

	router.GET("/clays", func(c *gin.Context) {
		clay, err := getClayOptions(db)
		if err != nil {
			log.Fatalf("Error getting clays: %v", err)
		}

		// convert categories to a JSON string
		jsonStr, err := json.Marshal(clay)
		if err != nil {
			// handle error
		}

		// convert the JSON string to a []byte value
		data := []byte(jsonStr)

		c.Data(http.StatusOK, "application/json", data)
	})

	router.GET("/cones", func(c *gin.Context) {
		cone, err := getConeOptions(db)
		if err != nil {
			log.Fatalf("Error getting cones: %v", err)
		}

		// convert categories to a JSON string
		jsonStr, err := json.Marshal(cone)
		if err != nil {
			// handle error
		}

		// convert the JSON string to a []byte value
		data := []byte(jsonStr)

		c.Data(http.StatusOK, "application/json", data)
	})

	// router.GET("/images", func(c *gin.Context) {
	// 	images, err := getImages(db)
	// 	if err != nil {
	// 		log.Fatalf("Error querying database: %v", err)
	// 	}
	// 	json, err := json.Marshal(images)
	// 	if err != nil {
	// 		log.Fatalf("Error encoding JSON: %v", err)
	// 	}

	// 	c.Data(http.StatusOK, "application/json", json)
	// })

	router.POST("/pieces", func(c *gin.Context) {
		var piece Piece
		if err := c.ShouldBindJSON(&piece); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := postPiece(db, piece)
		if err != nil {
			log.Fatalf("Error inserting piece: %v", err)
		}
		json, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusCreated, "application/json", json)
	})

	router.POST("/images", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatalf("Error receiving image: %v", err)
		}
		res, err := postImages(uploader, file)
		if err != nil {
			log.Fatalf("Error inserting image: %v", err)
		}

		// json, err := json.Marshal(res)

		c.Data(http.StatusCreated, "text/plain", []byte(res))
	})

	router.DELETE("/pieces/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := deletePieceById(db, id, sess)
		if err != nil {
			log.Fatalf("Error querying database: %v", err)
		}
	})

	router.Run("localhost:8080")
}

// getpieces responds with the list of all pieces as JSON.
func getPieces(db *sql.DB) ([]Piece, error) {
	rows, err := db.Query("SELECT * FROM piece")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pieces []Piece
	for rows.Next() {
		var piece Piece
		var imageTemp []uint8
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Category, &piece.Description, &piece.Size, &imageTemp)
		if err != nil {
			log.Fatalf("Error scanning result: %v", err)
		}
		var imageArray []string
		for _, v := range bytes.Split(imageTemp, []byte(",")) {
			imageArray = append(imageArray, strings.Trim(string(bytes.TrimSpace(v)), "{}"))
		}
		piece.Images = imageArray
		pieces = append(pieces, piece)
	}

	return pieces, nil
}

// func getImages(db *sql.DB) ([]Image, error) {
// 	rows, err := db.Query("SELECT * FROM image")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var images []Image
// 	for rows.Next() {
// 		var image Image
// 		err := rows.Scan(&image.ID, &image.Piece_ID, &image.Filename, &image.Data)
// 		if err != nil {
// 			log.Fatalf("Error scanning result: %v", err)
// 		}
// 		images = append(images, image)
// 	}

// 	return images, nil
// }

// postPieces adds an piece from JSON received in the request body.
func postPiece(db *sql.DB, piece Piece) ([]Piece, error) {
	id := uuid.New().String()
	_, err := db.Exec("INSERT INTO piece VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", id, piece.Title, piece.Artist, piece.Glaze_Description, piece.Clay, piece.Bisque_Cone, piece.Glaze_Cone, piece.Date, piece.Category, piece.Description, piece.Size, pq.Array(piece.Images))
	var pieces []Piece
	if err != nil {
		log.Fatalf("Error inserting piece: %v", err)
	}

	pieces = append(pieces, piece)

	return pieces, nil
}

func postImages(uploader *s3manager.Uploader, image *multipart.FileHeader) (string, error) {
	file, err := image.Open()
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("arfol-images"),
		Key:    aws.String(uuid.NewString()),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("Error uploading image: %v", err)
	}
	return result.Location, nil
}

// getPieceByID locates the piece whose ID value matches the id
// parameter sent by the client, then returns that piece as a response.
func getPieceByID(db *sql.DB, id string) ([]Piece, error) {
	rows, err := db.Query("SELECT * FROM piece WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pieces []Piece
	for rows.Next() {
		var piece Piece
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Category, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Description, &piece.Size, &piece.Images)
		if err != nil {
			log.Fatalf("Error scanning result: %v", err)
		}
		pieces = append(pieces, piece)
	}

	if len(pieces) == 0 {
		log.Fatalf("No Piece with id: %v", id)
	}
	if len(pieces) > 1 {
		log.Fatalf("Multiple pieces with id: %v", id)
	}

	return pieces, nil
}

func deletePieceById(db *sql.DB, id string, sess *session.Session) ([]Piece, error) {
	svc := s3.New(sess)
	resp, er := db.Query("SELECT * FROM piece WHERE id = $1", id)
	if er != nil {
		log.Fatalf("Error getting urls in delete: %v", er)
	}
	resp.Next()
	var piece Piece
	var urls []string
	var imageTemp []uint8
	scanErr := resp.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Category, &piece.Description, &piece.Size, &imageTemp)
	if scanErr != nil {
		log.Fatalf("Error scanning piece for urls %v: ", scanErr)
	}
	for _, v := range bytes.Split(imageTemp, []byte(",")) {
		urls = append(urls, strings.Trim(string(bytes.TrimSpace(v)), "{}"))
	}
	for _, url := range urls {
		parts := strings.SplitN(url, "/", 4)
		if len(parts) != 4 {
			return nil, fmt.Errorf("invalid S3 URL format: %s", url)
		}

		if parts[3] == "noImage.jpg" {
			break
		}

		params := &s3.DeleteObjectInput{
			Bucket: aws.String("arfol-images"),
			Key:    aws.String(parts[3]),
		}
		print(parts[3])
		print("\n")

		var _, deleteErr = svc.DeleteObject(params)
		if deleteErr != nil {
			log.Fatalf("Error deleting image %v: ", deleteErr)
		}

	}
	_, err := db.Query("DELETE FROM piece WHERE id = $1", id)
	return nil, err
}

func getClayOptions(db *sql.DB) ([]string, error) {
	categories, err := getOptions(db, "clay_type")
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func getConeOptions(db *sql.DB) ([]string, error) {
	categories, err := getOptions(db, "cone_type")
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func getOptions(db *sql.DB, option string) ([]string, error) {
	rows, err := db.Query("SELECT enumlabel FROM pg_enum WHERE enumtypid = (SELECT oid FROM pg_type WHERE typname = $1)ORDER BY enumsortorder;", option)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []string
	for rows.Next() {
		var option string
		if err := rows.Scan(&option); err != nil {
			return nil, err
		}
		options = append(options, option)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return options, nil
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
