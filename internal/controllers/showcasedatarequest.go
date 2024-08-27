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

func GetShowcaseData(c *gin.Context) {

	// middleware.CorsConfig(c, "GET")
	var res models.SucceesResponse
	var errRes models.RequestError

	if c.Query("lang") == "" {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(400, errRes)
		return
	}
	opt := options.Find()
	opt.SetLimit(3)
	opt.SetProjection(bson.D{{"_id", 0}, {"showcaseId", 1}, {"showcaseUrl", 1}, {"showcaseImageUrl", 1}, {"showcaseTitle", bson.D{{c.Query("lang"), 1}}}})
	opt.SetSort(bson.D{{"showcaseId", 1}})
	coll := db.DB.Database(os.Getenv("MONGODB_DB_NAME")).Collection(os.Getenv("MONGODB_DB_SHOWCASE_COLLECTION_NAME"))
	fetchedData, err := coll.Find(context.TODO(), bson.D{{"showcaseId", bson.D{{"$ne", ""}}}}, opt)

	if err != nil {
		errRes.Code = 404
		errRes.Message = "data not found"
		c.JSON(404, errRes)
		return
	}

	var data []map[string]interface{}
	fetchedData.All(context.TODO(), &data)
	res.Code = 200
	res.Message = "Success getting page data"
	res.Data = data
	c.JSON(200, res)
}
