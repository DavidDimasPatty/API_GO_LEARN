package service

import (
	"example/projectIseng1/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArtistService struct {
	artist []model.Artist
}

// constructor
func NewArtistService() *ArtistService {
	return &ArtistService{
		artist: []model.Artist{
			{ID: "1", Nama: "David", Rating: 4},
			{ID: "2", Nama: "Dimas", Rating: 2},
		},
	}
}

// method
func (s *ArtistService) GetArtist(c *gin.Context) {
	c.JSON(http.StatusOK, s.artist)
}

func (s *ArtistService) AddArtist(c *gin.Context) {
	var NewArtist model.Artist

	if err := c.ShouldBindJSON(&NewArtist); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messages": err.Error()})
		return
	}

	s.artist = append(s.artist, NewArtist)
	c.JSON(http.StatusCreated, gin.H{"messages": "success"})
}

func (s *ArtistService) GetArtistByID(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(s.artist); i++ {
		if s.artist[i].ID == id {
			c.JSON(http.StatusOK, s.artist[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"messages": "Not Found"})
}
