package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
}

type Image struct {
	ID       uuid.UUID `json:"id"`
	Piece_ID uuid.UUID `json:"piece_id"`
	Filename string    `json:"filename"`
	Data     []byte    `json:"data"`
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

	router.GET("/images", func(c *gin.Context) {
		images, err := getImages(db)
		if err != nil {
			log.Fatalf("Error querying database: %v", err)
		}
		json, err := json.Marshal(images)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusOK, "application/json", json)
	})

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
		var image Image
		if err := c.ShouldBindJSON(&image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := postImages(db, image)
		if err != nil {
			log.Fatalf("Error inserting image: %v", err)
		}
		json, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}

		c.Data(http.StatusCreated, "application/json", json)
	})

	router.DELETE("/pieces/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := deletePieceById(db, id)
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
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Category, &piece.Description, &piece.Size)
		if err != nil {
			log.Fatalf("Error scanning result: %v", err)
		}
		pieces = append(pieces, piece)
	}

	return pieces, nil
}

func getImages(db *sql.DB) ([]Image, error) {
	rows, err := db.Query("SELECT * FROM image")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var image Image
		err := rows.Scan(&image.ID, &image.Piece_ID, &image.Filename, &image.Data)
		if err != nil {
			log.Fatalf("Error scanning result: %v", err)
		}
		images = append(images, image)
	}

	return images, nil
}

// postPieces adds an piece from JSON received in the request body.
func postPiece(db *sql.DB, piece Piece) ([]Piece, error) {
	id := uuid.New().String()
	_, err := db.Exec("INSERT INTO piece VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", id, piece.Title, piece.Artist, piece.Glaze_Description, piece.Clay, piece.Bisque_Cone, piece.Glaze_Cone, piece.Date, piece.Category, piece.Description, piece.Size)
	var pieces []Piece
	if err != nil {
		return pieces, err
	}

	pieces = append(pieces, piece)

	return pieces, nil
}

func postImages(db *sql.DB, image Image) ([]Image, error) {
	id := uuid.New().String()
	_, err := db.Exec("INSERT INTO image VALUES ($1, $2, $3, $4)", id, image.Piece_ID, image.Filename, image.Data)
	var images []Image
	if err != nil {
		return images, err
	}

	images = append(images, image)

	return images, nil
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
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Glaze_Description, &piece.Category, &piece.Clay, &piece.Bisque_Cone, &piece.Glaze_Cone, &piece.Date, &piece.Description, &piece.Size)
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

func deletePieceById(db *sql.DB, id string) ([]Piece, error) {
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
