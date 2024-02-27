package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func RespondWithError(c *gin.Context, statusCode int, errorMessage string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"error": errorMessage})
}

func RespondWithData(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}

func ParseId(c *gin.Context) (uint, error) {
	val, err := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(val), err
}

func GetUserId(c *gin.Context) uint {
	return c.GetUint("user")
}
