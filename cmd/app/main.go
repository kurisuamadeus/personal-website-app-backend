package main

import (
	"example/personal-website-app-backend/api"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	api.ApiPath(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
