package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Song struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

var Songs = []Song{
	{ID: 1, Name: "Alabaster", Artist: "Foals", Album: "Total Life Forever"},
	{ID: 2, Name: "Bravery", Artist: "Human Tetris", Album: "River Pt. 1"},
	{ID: 3, Name: "Lately", Artist: "Metronomy", Album: "Metronomy Forever"},
	{ID: 4, Name: "Paranoid Android", Artist: "Radiohead", Album: "OK Computer"},
}

func ListSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Songs)
}

func main() {
	r := gin.Default()
	r.GET("/songs", ListSongs)

	r.Run("0.0.0.0:8080")
}
