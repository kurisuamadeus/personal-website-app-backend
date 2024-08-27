package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kurisuamadeus/personal-website-app-backend/api"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/db"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("DOCKER") != "true" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatalln(".env not loaded properly")
		}
	}
	r := gin.Default()

	middleware.CorsRouterConfig(r)

	api.ApiPath(r)
	db.MongoDB()
	defer func() {
		if err := db.DB.Disconnect(context.TODO()); err != nil {
			fmt.Println(err)
		}
	}()
	r.Run(getDomain() + os.Getenv("PORT"))
}

func getDomain() string {
	if os.Getenv("DOCKER") == "true" {
		return "0.0.0.0:"
	} else {
		return "localhost:"
	}
}
