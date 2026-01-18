package service

import (
	"example/projectIseng1/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumService struct {
	albums []model.Album
}

// constructor
func NewAlbumService() *AlbumService {
	return &AlbumService{
		albums: []model.Album{
			{ID: "1", Title: "Blue", Artist: "David", Price: 12.5},
			{ID: "2", Title: "Red", Artist: "Dimas", Price: 22.0},
			{ID: "3", Title: "Yellow", Artist: "Patty", Price: 16.1},
		},
	}
}

// functions
func (s *AlbumService) GetAlbum(c *gin.Context) {
	c.JSON(http.StatusOK, s.albums)
}

func (s *AlbumService) CreateAlbum(c *gin.Context) {
	var newAlbum model.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.albums = append(s.albums, newAlbum)
	c.JSON(http.StatusCreated, s.albums)
}

func (s *AlbumService) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range s.albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
