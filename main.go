package main

import (
	"example/projectIseng1/config"
	"example/projectIseng1/router"
	"example/projectIseng1/service"
	"log"
)

// Model
// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// Data
// var albums = []album{
// 	{ID: "1", Title: "Blue", Artist: "David", Price: 12.5},
// 	{ID: "2", Title: "Red", Artist: "Dimas", Price: 22.0},
// 	{ID: "3", Title: "Yellow", Artist: "Patty", Price: 16.1},
// }

// // Service
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, albums)
// }

// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

// handler
// func main() {
// 	router := gin.Default()
// 	router.GET("/albums", getAlbums)
// 	router.POST("/albums", postAlbums)
// 	router.GET("/albums/:id", getAlbumByID)
// 	router.Run("localhost:8080")
// }

func main() {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	songService := service.NewSongService(db)
	albumService := service.NewAlbumService()
	artistService := service.NewArtistService()
	r := router.SetUpRouter(albumService, artistService, songService)
	r.Run(":8080")
}
