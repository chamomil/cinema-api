package main

import (
	"cinema/db"
	"cinema/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	initialize()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			handlers.MapAuthHandlers(auth)
		}
		feedbacks := v1.Group("/feedbacks")
		{
			handlers.MapFeedbackHandlers(feedbacks)
		}
		movies := v1.Group("/movies")
		{
			handlers.MapMovieHandlers(movies)
		}
		genres := v1.Group("/genres")
		{
			handlers.MapGenreHandlers(genres)
		}
	}

	port := os.Getenv("PORT")
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	db.Start()
}
