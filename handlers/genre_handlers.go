package handlers

import (
	"cinema/models"
	"cinema/services"
	"cinema/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getGenres(c *gin.Context) {
	genres := services.FindGenres()

	utils.RespondWithData(c, http.StatusOK, genres)
}

func getGenre(c *gin.Context) {
	id, err := utils.ParseId(c)
	genre, err := services.FindGenreByID(id)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithData(c, http.StatusOK, genre)
}

func createGenre(c *gin.Context) {
	var input models.CreateGenreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	genre, err := services.CreateGenre(input)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithData(c, http.StatusOK, genre)
}

func updateGenre(c *gin.Context) {
	var input models.UpdateGenreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := utils.ParseId(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	genre, err := services.UpdateGenre(id, input)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithData(c, http.StatusOK, genre)
}

func deleteGenre(c *gin.Context) {
	id, err := utils.ParseId(c)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	err = services.DeleteGenre(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithData(c, http.StatusOK, true)
}

func MapGenreHandlers(genres *gin.RouterGroup) {
	genres.GET("/", getGenres)
	genres.GET("/:id", getGenre)
	genres.POST("/", createGenre)
	genres.PUT("/:id", updateGenre)
	genres.DELETE("/:id", deleteGenre)
}
