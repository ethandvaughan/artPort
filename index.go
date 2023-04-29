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
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// database connection
const (
	host   = "arfol.cobkxdytb6pv.us-west-2.rds.amazonaws.com"
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
	Artist_Id         uuid.UUID      `json:"artist_id"`
}

type Artist struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	First_Name    string    `json:"first_name"`
	Last_Name     string    `json:"last_name"`
	Crated        time.Time `json:"created_at"`
	Profile_Photo string    `json:"profile_photo"`
}

type InvalidPasswordError struct {
	error
}

type NoUserError struct {
	error
}

func NewInvalidPasswordError() error {
	return InvalidPasswordError{fmt.Errorf("invalid password")}
}

func NewNoUserError() error {
	return NoUserError{fmt.Errorf("user not found")}
}

// main function to handle the routing of CRUD actions
func main() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		user, os.Getenv("POSTPASS"), host, dbname)

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
		log.Fatalf("Error selecting database: %v", err)
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

	router.GET("/user/:user_id", func(c *gin.Context) {
		userId := c.Param("user_id")
		user, err := getUser(db, userId)
		if err != nil {
			log.Fatalf("Error querying database: %v", err)
		}
		json, err := json.Marshal(user)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusOK, "application/json", json)
	})

	router.GET("/pieces/:user_id", func(c *gin.Context) {
		userId := c.Param("user_id")
		pieces, err := getPieceByUserID(db, userId)
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

	router.GET("/ceramic", func(c *gin.Context) {
		ceramics, err := getCeramics(db)
		if err != nil {
			log.Fatalf("Error getting ceramics %v: ", err)
		}
		json, err := json.Marshal(ceramics)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusOK, "application/json", json)
	})

	router.POST("/auth", func(c *gin.Context) {
		type Auth struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var auth Auth
		c.ShouldBindJSON(&auth)
		// validate user credentials
		artist, err := AuthenticateUser(db, auth.Username, auth.Password)
		if err != nil {
			if _, noUser := err.(NoUserError); noUser {
				fmt.Println("Error: no user")
				c.JSON(http.StatusBadRequest, gin.H{
					"token": "no user",
				})
				return
			}
			if _, invalPass := err.(InvalidPasswordError); invalPass {
				fmt.Println("Error: invalid password")
				c.JSON(http.StatusNotFound, gin.H{
					"token": "invalid password",
				})
				return
			}
		}

		// generate JWT token
		token, err := GenerateToken(artist)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to generate token",
			})
			return
		}
		// return token in response body
		c.JSON(http.StatusOK, gin.H{
			"token":   token,
			"user_id": artist.ID,
		})
	})

	router.POST("/user", func(c *gin.Context) {
		var artist Artist
		if err := c.BindJSON(&artist); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := postUser(db, artist)
		if err != nil {
			log.Fatalf("Error inserting user: %v", err)
		}
		json, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}
		c.Data(http.StatusCreated, "application/json", json)
	})

	router.POST("/piece", func(c *gin.Context) {
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

	router.PUT("/piece/:id", func(c *gin.Context) {
		var piece Piece
		if err := c.ShouldBindJSON(&piece); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := alterPiece(db, piece)
		if err != nil {
			log.Fatalf("Error altering piece: %v", err)
		}
		json, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusCreated, "application/json", json)

	})

	router.DELETE("/piece/:id", func(c *gin.Context) {
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
	rows, err := db.Query("SELECT * FROM piece ORDER BY category ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pieces []Piece
	for rows.Next() {
		var piece Piece
		var imageTemp []uint8
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Category, &piece.Description, &piece.Size, &imageTemp, &piece.Artist_Id)
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

// getPieceByUserID locates the piece whose user_id value matches the id
// parameter sent by the client, then returns that piece as a response.
func getPieceByUserID(db *sql.DB, id string) ([]Piece, error) {
	rows, err := db.Query("SELECT * FROM piece WHERE artist_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pieces []Piece
	var imageTemp []uint8
	for rows.Next() {
		var piece Piece
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Category, &piece.Description, &piece.Size, &imageTemp, &piece.Artist_Id)
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

// getUser responds with the user with given id as JSON.
func getUser(db *sql.DB, userId string) (Artist, error) {
	rows, err := db.Query("SELECT * FROM artist WHERE user.id = $1", userId)
	var artist Artist
	if err != nil {
		return artist, err
	}
	defer rows.Close()
	rows.Next()
	scanErr := rows.Scan(&artist.ID, &artist.Username, &artist.Password, &artist.First_Name, &artist.Last_Name, &artist.Crated, &artist.Profile_Photo)
	if scanErr != nil {
		log.Fatalf("Error scanning result: %v", err)
	}

	return artist, nil
}

func getCeramics(db *sql.DB) ([]Piece, error) {
	rows, err := db.Query("SELECT * FROM piece WHERE piece.category = 'Ceramic'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pieces []Piece
	for rows.Next() {
		var piece Piece
		var imageTemp []uint8
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Category, &piece.Description, &piece.Size, &imageTemp, &piece.Artist_Id)
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

// postPieces adds an piece from JSON received in the request body.
func postPiece(db *sql.DB, piece Piece) ([]Piece, error) {
	id := uuid.New().String()
	_, err := db.Exec("INSERT INTO piece VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", id, piece.Title, piece.Artist, piece.Glaze_Description, piece.Clay, piece.Bisque_Cone, piece.Glaze_Cone, piece.Date, piece.Category, piece.Description, piece.Size, pq.Array(piece.Images), piece.Artist_Id)
	var pieces []Piece
	if err != nil {
		log.Fatalf("Error inserting piece: %v", err)
	}

	pieces = append(pieces, piece)

	return pieces, nil
}

func postUser(db *sql.DB, artist Artist) (Artist, error) {
	artist.ID = uuid.New()
	artist.Crated = time.Now()
	_, err := db.Exec("INSERT INTO artist VALUES ($1, $2, $3, $4, $5, $6, $7)", artist.ID, artist.Username, artist.Password, artist.First_Name, artist.Last_Name, artist.Crated, artist.Profile_Photo)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	return artist, err

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

func alterPiece(db *sql.DB, piece Piece) (Piece, error) {
	_, err := db.Exec("UPDATE piece SET title=$1, artist=$2, glaze_description=$3, clay=$4, bisque_cone=$5, glaze_cone=$6, date=$7, category=$8, description=$9, size=$10 WHERE id=$11", piece.Title, piece.Artist, piece.Glaze_Description, piece.Clay, piece.Bisque_Cone, piece.Glaze_Cone, piece.Date, piece.Category, piece.Description, piece.Size, piece.ID)
	if err != nil {
		log.Fatalf("Error updating piece: %v", err)
	}
	return piece, nil
}

func deletePieceById(db *sql.DB, id string, sess *session.Session) ([]Piece, error) {
	svc := s3.New(sess)
	resp, er := db.Query("SELECT * FROM piece WHERE id = $1", id)
	if er != nil {
		log.Fatalf("Error getting piece in delete: %v", er)
	}
	resp.Next()
	var piece Piece
	var urls []string
	var imageTemp []uint8
	scanErr := resp.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Category, &piece.Description, &piece.Size, &imageTemp, &piece.Artist_Id)
	if scanErr != nil {
		fmt.Printf("Scann err: %v", scanErr)
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

func AuthenticateUser(db *sql.DB, username string, password string) (*Artist, error) {

	var artist Artist
	rows, err := db.Query("SELECT id, username, password FROM artist WHERE username = $1", username)
	if err != nil {
		if err == pgx.ErrNoRows {
			// User not found
			return &artist, NewNoUserError()
		} else {
			log.Fatalf("Failed to query database in AuthenticateUser: %v", err)
		}
	}
	rows.Next()
	rows.Scan(&artist.ID, &artist.Username, &artist.Password)

	if artist.Username != username {
		// User not found
		return &artist, NewNoUserError()
	}
	// Compare the password hash stored in the database with the hash of the password entered by the user
	var correctPass bool = artist.Password == password
	if correctPass != true {
		// Password does not match
		return &artist, NewInvalidPasswordError()
	}

	// Authentication successful
	return &artist, nil

}

func GenerateToken(artist *Artist) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": artist.ID.String(),
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte("secret-key")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
