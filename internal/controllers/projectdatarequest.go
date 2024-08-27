package controllers

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/db"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get list of projects
func GetAllProjectData(c *gin.Context) {

	// middleware.CorsConfig(c, "GET")
	var res models.SucceesResponse
	var errRes models.RequestError

	if c.Query("lang") == "" && c.Query("category") == "" {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(400, errRes)
		return
	}
	opt := options.Find()
	opt.SetProjection(bson.D{{"_id", 0}, {"projectId", 1}, {"url", 1}, {"title", 1}, {"data", bson.D{{c.Query("lang"), bson.D{{"title", 1}}}}}, {"thumbnailImageUrl", 1}})
	opt.SetSort(bson.D{{"_id", 1}})
	coll := db.DB.Database(os.Getenv("MONGODB_DB_NAME")).Collection(os.Getenv("MONGODB_DB_PROJECT_COLLECTION_NAME"))
	fetchedData, err := coll.Find(context.TODO(), bson.D{{"category", c.Query("category")}}, opt)
	// collRaw, err := data.Raw()

	if err != nil {
		errRes.Code = 404
		errRes.Message = "data not found"
		c.JSON(404, errRes)
		return
	}
	var data []map[string]interface{}
	fetchedData.All(context.TODO(), &data)
	// bson.Unmarshal(collRaw, &data)
	res.Code = 200
	res.Message = "Success getting page data"
	res.Data = data
	c.JSON(200, res)
}

// Get single project details
func GetProjectDataById(c *gin.Context) {

	// middleware.CorsConfig(c, "GET")
	var res models.SucceesResponse
	var errRes models.RequestError

	if c.Query("lang") == "" && c.Query("projectId") == "" {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(400, errRes)
		return
	}
	opt := options.FindOne()
	opt.SetProjection(bson.D{{"_id", 0}, {"projectId", 1}, {"url", 1}, {"thumbnailImageUrl", 1}, {"title", 1}, {"data", bson.D{{c.Query("lang"), 1}}}})
	coll := db.DB.Database(os.Getenv("MONGODB_DB_NAME")).Collection(os.Getenv("MONGODB_DB_PROJECT_COLLECTION_NAME"))
	fetchedData := coll.FindOne(context.TODO(), bson.D{{"projectId", c.Query("projectId")}}, opt)
	dataRaw, err := fetchedData.Raw()

	if err != nil {
		errRes.Code = 404
		errRes.Message = "data not found"
		c.JSON(404, errRes)
		return
	}
	var data map[string]interface{}
	bson.Unmarshal(dataRaw, &data)
	res.Code = 200
	res.Message = "Success getting page data"
	res.Data = data
	c.JSON(200, res)
}
