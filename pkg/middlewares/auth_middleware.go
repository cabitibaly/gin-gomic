package middlewares

import (
	"net/http"
	"time"

	"github.com/cbitbaly/internal/services"
	"github.com/cbitbaly/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  "Impossible de vous authentifier",
				"status": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		jwtClaims, err := utils.ValidateToken(cookie)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  err.Error(),
				"status": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		if jwtClaims.ExpiresAt.Unix() < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  "Votre session a expiré",
				"status": http.StatusUnauthorized,
			})
			return
		}

		c.Set("userID", jwtClaims.UserID)
		c.Set("email", jwtClaims.Email)

		c.Next()
	}
}
