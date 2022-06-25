package main

import (
	"the-movie-backend/movies"
	"the-movie-backend/tv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/movies/getTopRated/:page", movies.GetTopRated)
	router.GET("/movies/getUpcoming/:page", movies.GetUpcoming)
	router.GET("/tv/getTopRated/:page", tv.GetTopRated)
	router.GET("/tv/getPopular/:page", tv.GetPopular)

	router.Run("localhost:8080")
}
