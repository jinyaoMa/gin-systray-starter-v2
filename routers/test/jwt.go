package test

import (
	"App/utils"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// @Description Test JWT
// @Tags Test
// @accept plain
// @Produce json
// @Success 200 "{ token }"
// @Failure 404 "{ error }"
// @Router /test/jwt [get]
func TestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		key, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		privateBytes := x509.MarshalPKCS1PrivateKey(key)
		publicBytes := x509.MarshalPKCS1PublicKey(&key.PublicKey)
		x509.ParsePKCS1PrivateKey(privateBytes)
		x509.ParsePKCS1PublicKey(publicBytes)

		token, err := utils.EncryptClaims(&utils.Claims{
			RegisteredClaims: jwt.RegisteredClaims{},
		}, key)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
