package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/package/tokens"
)

// method for working with jwt based authentication
func JwtAuthMiddleWare(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "missing of header"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "missing of bearer token"})
			c.Abort()
			return
		}
		ok, _ := tokens.VerifyToken(tokenString, secret)
		if !ok {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "in valid token"})
			c.Abort()
			return
		}
		claims, err := tokens.GetUserClaims(tokenString, secret)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "in valid token"})
			c.Abort()
			return
		}
		c.Set("username", claims["username"])
		c.Set("id", claims["id"])
		c.Next()
	}
}
