package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

// album represents data about a record album.
/*type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
*/
type simpleJson struct {
	//testa string `json:"testa"`
}

// albums slice to seed record album data.
/*var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
*/
// getAlbums responds with the list of all albums as JSON.

func albums(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GİRİŞ EKRANI"))
}
func getAlbums(c *gin.Context) {
	//fmt.Printf("Ulooo", c.Request)
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	//var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	//if err := c.BindJSON(&newAlbum); err != nil {
	return
}

// Add the new album to the slice.
//	albums = append(albums, newAlbum)
//c.IndentedJSON(http.StatusCreated, newAlbum)
//}

func testaPost(c *gin.Context) {
	var newTesta simpleJson

	if err := c.BindJSON(&newTesta); err != nil {
		return
	}

	fmt.Printf("%s", c.Request.Body)

	c.IndentedJSON(http.StatusCreated, c.Request.Body)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	http.HandleFunc("/albums", albums)
	router.GET("/albums", getAlbums)
	router.POST("/new", testaPost)

	router.Run(":" + port)
}
