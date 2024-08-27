package middleware

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfig(c *gin.Context, method string) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_DOMAIN"))
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", method)
	// c.Writer.Header().Set("Access-Control-Max-Age", "20")
	c.Writer.Header().Set("Content-Type", "application/json")
}

func CorsRouterConfig(r *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{os.Getenv("FRONTEND_DOMAIN")}
	corsConfig.AllowMethods = []string{"GET, POST"}
	r.Use(cors.New(corsConfig))
}
