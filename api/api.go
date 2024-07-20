package api

import (
	"example/personal-website-app-backend/api/routes"

	"github.com/gin-gonic/gin"
)

func ApiPath(r *gin.Engine) {

	routes.Get(r)
	// routes.Post(r)
	// routes.Put(r)
	// routes.Patch(r)
	// routes.Delete(r)

}
