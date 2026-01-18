package model

import "time"

type Song struct {
	ID          int       `json:"id"`
	Nama        string    `json:"nama"`
	ReleaseDate time.Time `json:"release_date"`
	AddTime     time.Time `json:"addtime"`
	AddID       int       `json:"addid"`
}
