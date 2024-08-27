package routes

import (
	"github.com/kurisuamadeus/personal-website-app-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine) {

	r.GET("/getpagedata", controllers.GetPageData)
	r.GET("/gettoolsdata", controllers.GetAllToolsList)
	r.GET("/getshowcase", controllers.GetShowcaseData)
	r.GET("/getproject/search", controllers.GetAllProjectData)
	r.GET("/getproject/details", controllers.GetProjectDataById)
}
func Post(r *gin.Engine) {

	r.POST("/sendmessage", controllers.PostNewMessage)
}
