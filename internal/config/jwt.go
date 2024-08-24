package config

import (
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
)

func JWTConfig() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Key:           []byte(os.Getenv("JWT_KEY")),
		Timeout:       time.Hour * 3, // token expire time
		MaxRefresh:    time.Hour * 3,
		IdentityKey:   "id", // unique key
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
	}
}
