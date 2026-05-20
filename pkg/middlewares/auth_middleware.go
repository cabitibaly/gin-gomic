package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/cbitbaly/internal/services"
	"github.com/cbitbaly/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  "Aucun token d'accès n'a été fourni",
				"status": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		part := strings.Split(authorization, " ")
		if len(part) != 2 || part[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  "Le token d'accès n'est pas valide",
				"status": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		jwtClaims, err := utils.ValidateAccessToken(part[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  err.Error(),
				"status": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		location, _ := time.LoadLocation("Africa/Ouagadougou")

		if jwtClaims.ExpiresAt.Unix() < time.Now().In(location).Unix() {
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
