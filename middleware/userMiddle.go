package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	myjwt "github.com/hewo233/hdu-cxsj1/utils/jwt"
	"log"
	"net/http"
)

func JWTAuth(audience string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			log.Println("No token")
			c.JSON(http.StatusBadRequest, gin.H{
				"errno": 40000,
				"msg":   "Bad Request, no token",
			})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		token, err := jwt.ParseWithClaims(tokenString, &myjwt.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return myjwt.JWTKey, nil
		})
		if err != nil || !token.Valid {
			log.Println("Parse token error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errno": 50000,
				"msg":   "status internal server error, token parse error",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*myjwt.Claims); ok {
			if claims.Audience != audience {
				log.Println("Audience error")
				c.JSON(http.StatusUnauthorized, gin.H{
					"errno": 40100,
					"msg":   "Unauthorized, audience error",
				})
				c.Abort()
				return
			}

			c.Set("userEmail", claims.Email)
			c.Set("uid", claims.StandardClaims.Id)
		}
	}
}
