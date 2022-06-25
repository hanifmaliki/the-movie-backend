package main

import (
	"net/http"
	// "os"
	"the-movie-backend/movies"
	"the-movie-backend/tv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/movies/getTopRated/:page", movies.GetTopRated)
	router.GET("/movies/getUpcoming/:page", movies.GetUpcoming)
	router.GET("/tv/getTopRated/:page", tv.GetTopRated)
	router.GET("/tv/getPopular/:page", tv.GetPopular)

	// port := os.Getenv("PORT")
	// router.Run(":" + port)
	router.Run(":8080")
}

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to Hanif's Api"})
}
