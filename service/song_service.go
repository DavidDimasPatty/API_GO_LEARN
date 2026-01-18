package service

import (
	"database/sql"
	"example/projectIseng1/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// declare variabel
type SongService struct {
	db *sql.DB
}

// constructor
func NewSongService(db *sql.DB) *SongService {
	return &SongService{db: db}
}

// method
func (s *SongService) GetAllSongs(c *gin.Context) {
	rows, err := s.db.Query(`select * from songs`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer rows.Close()
	var songs []model.Song

	for rows.Next() {
		var song model.Song
		if err := rows.Scan(
			&song.ID,
			&song.Nama,
			&song.ReleaseDate,
			&song.AddTime,
			&song.AddID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		songs = append(songs, song)
	}

	c.JSON(http.StatusOK, songs)
}
