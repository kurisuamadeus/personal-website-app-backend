package controllers

import (
	"encoding/json"
	"example/personal-website-app-backend/internal/models"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetPageData(c *gin.Context) {

	var res models.SucceesResponse
	var errRes models.RequestError

	if c.Query("lang") == "" || c.Query("pagename") == "" {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(500, errRes)
		return
	}
	path, err := filepath.Abs("internal/localizationdata/locales/" + c.Query("lang") + "/" + c.Query("pagename") + "page.json")
	if err != nil {

		errRes.Code = 404
		errRes.Message = "data not found"
		c.JSON(500, errRes)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		errRes.Code = 500
		errRes.Message = "internal server error"
		c.JSON(500, errRes)
		return
	}
	defer file.Close()
	byteData, _ := io.ReadAll(file)
	var data map[string]interface{}
	json.Unmarshal(byteData, &data)
	res.Code = 200
	res.Message = "Success getting page data" + c.Query("test")
	res.Data = data
	c.JSON(200, res)
}
