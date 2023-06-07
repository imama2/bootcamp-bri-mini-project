package middleware

import (
	"github.com/gin-gonic/gin"
	token2 "github.com/imama2/bootcamp-bri-mini-project/utils/token"
	"log"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		auth_token := token2.SplitBearer(tokenString)
		_, err := token2.VerifyAccessToken(auth_token)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		accountData, err := token2.GetDataUserFromToken(auth_token)
		if err != nil {
			log.Println(err)
		}
		if !accountData.IsVerified || !accountData.IsActive {
			c.JSON(401, gin.H{"error": "User admin must be verified and active"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		auth_token := token2.SplitBearer(tokenString)
		_, err := token2.VerifyAccessToken(auth_token)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		accountData, err := token2.GetDataUserFromToken(auth_token)
		if err != nil {
			log.Println(err)
		}
		if !accountData.IsVerified || !accountData.IsActive {
			c.JSON(401, gin.H{"error": "User admin must be verified and active"})
			c.Abort()
			return
		}
		if accountData.RoleID != 2 {
			c.JSON(401, gin.H{"error": "Only for super admin role"})
			c.Abort()
		}
		c.Next()
	}
}
