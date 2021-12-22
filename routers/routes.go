package routers

import (
	"App/routers/test"
	_ "App/swagger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title gin-systray-starter-v2
// @version 2.0.0
// @description "Template for Go application with Gin and System Tray"

// @contact.name Github Issues
// @contact.url https://github.com/jinyaoMa/gin-systray-starter-v2/issues

// @license.name MIT
// @license.url https://github.com/jinyaoMa/gin-systray-starter-v2/blob/main/LICENSE

// @securityDefinitions.apikey BearerIdAuth
// @in header
// @name Authorization

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
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
