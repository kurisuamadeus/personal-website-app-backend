package routes

import (
	"github.com/kurisuamadeus/personal-website-app-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine) {

	r.GET("/getpagedata", controllers.GetPageData)
}
func Post(r *gin.Engine) {

	r.POST("/sendmessage", controllers.PostNewMessage)
}
