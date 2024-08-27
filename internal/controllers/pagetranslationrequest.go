package controllers

import (
	"context"
	"os"

	"github.com/kurisuamadeus/personal-website-app-backend/internal/db"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPageData(c *gin.Context) {

	// middleware.CorsConfig(c, "GET")
	var res models.SucceesResponse
	var errRes models.RequestError

	if c.Query("lang") == "" || c.Query("dataname") == "" {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(400, errRes)
		return
	}
	opt := options.FindOne()
	opt.SetProjection(bson.D{{"_id", 0}})
	coll := db.DB.Database(os.Getenv("MONGODB_DB_NAME")).Collection(os.Getenv("MONGODB_DB_PAGE_COLLECTION_NAME")).FindOne(context.TODO(), bson.D{{"lang", c.Query("lang")}, {"dataName", c.Query("dataname")}}, opt)
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
