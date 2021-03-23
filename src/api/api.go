package api

import (
	"fmt"
	"net/http"

	conf "github.com/NODO-UH/uh-email-quota/src/config"
	"github.com/NODO-UH/uh-email-quota/src/quota"
	"github.com/gin-gonic/gin"
)

func StartAPI() error {
	server := gin.New()
	server.Use(gin.Logger())

	server.Use(authorizeAPIKey())
	server.GET("/quota", getUserQuota)

	if err := server.Run(); err != nil {
		return err
	}
	return nil
}

func authorizeAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get API Key from header
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Check API Key
		configuration := conf.GetConfiguration()
		if apiKey != *configuration.APIKey {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func getUserQuota(c *gin.Context) {
	if userEmail := c.Query("userEmail"); userEmail == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, APIError{
			Code:    ErrUserMissing,
			Message: "userEmail is required",
		})
	} else if quotaInfo, err := quota.GetUserQuota(userEmail); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, APIError{
			Code:    ErrUserNotFound,
			Message: fmt.Sprintf("user %s not found", userEmail),
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, quotaInfo)
	}
}
