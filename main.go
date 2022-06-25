package main

import (
	"the-movie-backend/example"
	"the-movie-backend/get"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", example.GetAlbums)
	router.GET("/albums/:id", example.GetAlbumByID)
	router.POST("/albums", example.PostAlbums)

	router.GET("/getTopRated/:page", get.GetTopRated)

	router.Run("localhost:8080")
}
