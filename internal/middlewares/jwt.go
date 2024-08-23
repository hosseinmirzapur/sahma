package middlewares

import (
	"sahma/internal/config"
	"sahma/internal/database/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func NewJWT() *jwt.GinJWTMiddleware {
	conf := config.JWTConfig()

	// Add more configuration to jwt config
	conf.PayloadFunc = payloadFunc
	conf.IdentityHandler = identityHandler

	return conf
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			"id": v.ID,
		}
	}

	return nil
}

func identityHandler(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return &models.User{
		ID: claims["id"].(uint),
	}
}
