package main

import (
	"database/sql"
	"fmt"
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
type piece struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Category string `json:"category"`
}

// pieces slice to seed record piece data.
var pieces = []piece{
	{ID: "1", Title: "Ice Cream Bowls", Artist: "LeeAnn Vaughan", Category: "Ceramic"},
	{ID: "2", Title: "Mt. Shasta", Artist: "Ashley Rosenbaum", Category: "Water Color"},
	{ID: "3", Title: "Mt. St. Helens", Artist: "Ashley Rosenbaum", Category: "Water Color"},
}

// getpieces responds with the list of all pieces as JSON.
func getPieces(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pieces)
}

// postPieces adds an piece from JSON received in the request body.
func postPieces(c *gin.Context) {
	var newPiece piece

	// Call BindJSON to bind the received JSON to
	// newpiece.
	if err := c.BindJSON(&newPiece); err != nil {
		return
	}

	// Add the new piece to the slice.
	pieces = append(pieces, newPiece)
	c.IndentedJSON(http.StatusCreated, newPiece)
}

// getPieceByID locates the piece whose ID value matches the id
// parameter sent by the client, then returns that piece as a response.
func getPieceByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of pieces, looking for
	// an piece whose ID value matches the parameter.
	for _, a := range pieces {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "piece not found"})
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
	router.GET("/pieces", getPieces)
	router.GET("/pieces/:id", getPieceByID)
	router.POST("/pieces", postPieces)

	router.Run("localhost:8080")
}
