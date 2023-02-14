package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ev77916!"
	dbname   = "artport"
)

// piece represents data about a record piece.
type Piece struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Category string `json:"category"`
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
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Category)
		if err != nil {
			log.Fatalf("Error scanning result: %v", err)
		}
		pieces = append(pieces, piece)
	}

	return pieces, nil
}

// postPieces adds an piece from JSON received in the request body.
func postPieces(db *sql.DB, piece Piece) ([]Piece, error) {
	_, err := db.Exec("INSERT INTO piece VALUES ($1, $2, $3, $4)", piece.ID, piece.Title, piece.Artist, piece.Category)
	var pieces []Piece
	if err != nil {
		return pieces, err
	}

	pieces = append(pieces, piece)

	return pieces, nil
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
		err := rows.Scan(&piece.ID, &piece.Title, &piece.Artist, &piece.Category)
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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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

	router.POST("/pieces", func(c *gin.Context) {
		var piece Piece
		if err := c.ShouldBindJSON(&piece); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := postPieces(db, piece)
		if err != nil {
			log.Fatalf("Error inserting piece: %v", err)
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
