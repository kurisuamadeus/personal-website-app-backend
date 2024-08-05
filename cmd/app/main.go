package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kurisuamadeus/personal-website-app-backend/api"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/db"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalln(".env not loaded properly")
	}

	r := gin.Default()

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	// corsConfig.AllowMethods = []string{"GET"}
	// corsConfig.AllowHeaders = []string{"Origin"}
	// corsConfig.AllowCredentials = false
	middleware.CorsRouterConfig(r)
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://localhost:3000"
	// 	},
	// }))
	api.ApiPath(r)
	db.MongoDB()
	defer func() {
		if err := db.DB.Disconnect(context.TODO()); err != nil {
			fmt.Println(err)
		}
	}()
	r.Run() // listen and serve on 0.0.0.0:8080
}
