package handlers

import (
	middleware "cinema/middlewares"
	"cinema/services"
	"cinema/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		utils.RespondWithError(c, http.StatusUnprocessableEntity, "Failed to read body")
		return
	}

	if err := services.Signup(body.Email, body.Password); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	utils.RespondWithData(c, http.StatusOK, "User created successfully")
}

func login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		utils.RespondWithError(c, http.StatusUnprocessableEntity, "Failed to read body")
		return
	}

	token, err := services.Login(body.Email, body.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid email or password")
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	utils.RespondWithData(c, http.StatusOK, "Successfully logged in")
}

func logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	utils.RespondWithData(c, http.StatusOK, "Successfully logged out")
}

func MapAuthHandlers(genres *gin.RouterGroup) {
	genres.POST("/sign-up", signUp)
	genres.POST("/login", login)
	genres.POST("/logout", middleware.RequireAuth, logout)
}
