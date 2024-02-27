package handlers

import (
	"cinema/models"
	"cinema/services"
	"cinema/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMovies(c *gin.Context) {
	movies, err := services.GetAllMovies()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithData(c, http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	id, err := utils.ParseId(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	movie, err := services.GetMovieByID(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Record not found!")
		return
	}
	utils.RespondWithData(c, http.StatusOK, movie)
}

func createMovie(c *gin.Context) {
	var input models.CreateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	movie, err := services.CreateMovie(input)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithData(c, http.StatusOK, movie)
}

func updateMovie(c *gin.Context) {
	id, err := utils.ParseId(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	var input models.UpdateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	movie, err := services.UpdateMovie(id, input)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Record not found!")
		return
	}
	utils.RespondWithData(c, http.StatusOK, movie)
}

func deleteMovie(c *gin.Context) {
	id, err := utils.ParseId(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}
	err = services.DeleteMovie(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Record not found!")
		return
	}
	utils.RespondWithData(c, http.StatusOK, true)
}

func MapMovieHandlers(genres *gin.RouterGroup) {
	genres.GET("/", getMovies)
	genres.GET("/:id", getMovie)
	genres.POST("/", createMovie)
	genres.PUT("/:id", updateMovie)
	genres.DELETE("/:id", deleteMovie)
}
