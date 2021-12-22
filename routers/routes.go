package routers

import (
	"App/routers/test"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "root path",
		})
	})

	_test := r.Group("/test")
	{
		_test.GET("/jwt", test.TestJWT())
	}

	if withSwag {

	}
}
