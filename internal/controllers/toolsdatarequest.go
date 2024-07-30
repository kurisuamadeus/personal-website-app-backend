package controllers

import (
	"context"
	"example/personal-website-app-backend/internal/db"
	"example/personal-website-app-backend/internal/middleware"
	"example/personal-website-app-backend/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllToolsList(c *gin.Context) {

	middleware.CorsConfig(c, "GET")
	var res models.SucceesResponse
	var errRes models.RequestError
	if c.Query("lang") == "" {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(400, errRes)
		return
	}
	coll := db.DB.Database("db_personalwebsiteapp").Collection("stack_data").FindOne(context.TODO(), bson.D{{"lang", c.Query("lang")}})
	collRaw, err := coll.Raw()

	if err != nil {
		errRes.Code = 404
		errRes.Message = "data not found"
		c.JSON(404, errRes)
		return
	}
	var data map[string]interface{}
	bson.Unmarshal(collRaw, &data)
	res.Code = 200
	res.Message = "Success getting page data"
	res.Data = data
	c.JSON(200, res)
}
