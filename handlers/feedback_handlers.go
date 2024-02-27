package handlers

import (
	middleware "cinema/middlewares"
	"cinema/services"
	"cinema/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createFeedback(c *gin.Context) {
	userID := utils.GetUserId(c)
	var body struct {
		Message string
	}

	if err := c.Bind(&body); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	if err := services.CreateFeedback(userID, body.Message); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create feedback")
		return
	}

	utils.RespondWithData(c, http.StatusOK, "Feedback created successfully")
}

func getFeedbacks(c *gin.Context) {
	userID := utils.GetUserId(c)

	feedbacks, err := services.GetFeedbacksByUserID(userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch feedbacks")
		return
	}

	utils.RespondWithData(c, http.StatusOK, feedbacks)
}

func MapFeedbackHandlers(genres *gin.RouterGroup) {
	genres.Use(middleware.RequireAuth)
	genres.GET("/", getFeedbacks)
	genres.POST("/", createFeedback)
}
