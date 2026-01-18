package router

import (
	"example/projectIseng1/service"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(albumService *service.AlbumService, artistService *service.ArtistService, songService *service.SongService) *gin.Engine {
	r := gin.Default()
	//albums
	r.GET("/albums", albumService.GetAlbum)
	r.POST("/albums", albumService.CreateAlbum)
	r.GET("/albums/:id", albumService.GetAlbumByID)

	//artist
	r.GET("/artist", artistService.GetArtist)
	r.POST("/artist", artistService.AddArtist)
	r.GET("/artist/id", artistService.GetArtistByID)

	//song
	r.GET("/songs", songService.GetAllSongs)

	return r
}
