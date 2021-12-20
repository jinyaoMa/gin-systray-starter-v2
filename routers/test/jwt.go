package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": "",
		})
	}
}
