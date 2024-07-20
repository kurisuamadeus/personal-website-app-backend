package routes

import (
	"example/personal-website-app-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine) {

	r.GET("/getpagedata", controllers.GetPageData)
}
