package main

import (
	"context"
	"example/personal-website-app-backend/api"
	"example/personal-website-app-backend/internal/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	// corsConfig.AllowMethods = []string{"GET"}
	// corsConfig.AllowHeaders = []string{"Origin"}
	// corsConfig.AllowCredentials = false
	api.ApiPath(r)
	db.MongoDB()
	defer func() {
		if err := db.DB.Disconnect(context.TODO()); err != nil {
			fmt.Println(err)
		}
	}()
	r.Run() // listen and serve on 0.0.0.0:8080
}
