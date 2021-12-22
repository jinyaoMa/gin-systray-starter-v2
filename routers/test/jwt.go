package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description Test JWT
// @Tags Test
// @accept plain
// @Produce json
// @Success 200 "{ token }"
// @Router /test/jwt [get]
func TestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": "",
		})
	}
}
