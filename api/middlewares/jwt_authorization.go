package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"working.com/bank_dash/package/tokens"
)

// method for working with authorization

func JwtAuthoMiddleWare(role string, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "header not found"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		ok, _ := tokens.VerifyToken(tokenString, secret)
		if !ok {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}
		claims, err := tokens.GetUserClaims(tokenString, secret)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err})
			c.Abort()
			return
		}
		userRole, ok := claims["role"].(string)
		if !ok || userRole != role {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "error of authority"})
			c.Abort()
			return
		}
		c.Set("username", claims["username"])
		c.Set("id", claims["id"])
		c.Next()
	}
}
