package middleware

import (
	"cinema/services"
	"cinema/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Token not provided")
		return
	}

	userID, err := services.ParseToken(tokenString)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	user, err := services.GetUserByID(userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not found")
		return
	}

	c.Set("user", user.ID)
	c.Next()
}
