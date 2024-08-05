package controllers

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/kurisuamadeus/personal-website-app-backend/internal/db"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/helper"
	"github.com/kurisuamadeus/personal-website-app-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
)

type NewMessageFormat struct {
	Id       string
	DateTime time.Time
	Date     string
	Time     string
	Inquiry  string
	Email    string
	Name     string
	Message  string
}

var inquiryMap = map[string]string{
	"general":  "G",
	"question": "Q",
	"business": "B",
	"other":    "O",
}

func PostNewMessage(c *gin.Context) {

	//middleware.CorsConfig(c, "POST")
	var res models.SucceesResponse
	var errRes models.RequestError
	var data models.ContactForm
	rawData, err := c.GetRawData()
	if err != nil {
		errRes.Code = 500
		errRes.Message = "internal server error"
		c.JSON(400, errRes)
		return
	}
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(400, errRes)
		return
	}

	if !validateRequest(data) {
		errRes.Code = 400
		errRes.Message = "bad request"
		c.JSON(400, errRes)
		return
	}
	messageCount, err := db.DB.Database(os.Getenv("MONGODB_DB_NAME")).Collection(os.Getenv("MONGODB_DB_MESSAGE_COLLECTION_NAME")).CountDocuments(context.TODO(), bson.D{{"date", time.Now().Format("2006 January 02")}, {"inquiry", data.Inquiry}})
	if err != nil {
		errRes.Code = 500
		errRes.Message = "internal server error"
		c.JSON(500, errRes)
		return
	}
	var newMessage NewMessageFormat = NewMessageFormat{
		Id:       helper.FormatMessageId(inquiryMap[strings.ToLower(data.Inquiry)], messageCount+1),
		DateTime: time.Now(),
		Date:     time.Now().Format("2006 January 02"),
		Time:     helper.FormatTime(time.Now().Clock()),
		Inquiry:  data.Inquiry,
		Email:    data.Email,
		Name:     data.Name,
		Message:  data.Message,
	}
	_, err = db.DB.Database(os.Getenv("MONGODB_DB_NAME")).Collection(os.Getenv("MONGODB_DB_MESSAGE_COLLECTION_NAME")).InsertOne(context.TODO(), newMessage)
	if err != nil {
		errRes.Code = 500
		errRes.Message = "internal server error"
		c.JSON(500, errRes)
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USERNAME"))
	m.SetHeader("To", os.Getenv("EMAIL_USERNAME"))
	m.SetHeader("Subject", "[PersonalWebsiteMessage] "+newMessage.Inquiry+" #"+newMessage.Id)
	m.SetBody("text/html", helper.GetFormattedHTMLMessage("NEW MESSAGE", newMessage.Id, newMessage.Inquiry, newMessage.Name, newMessage.Email, newMessage.Date, newMessage.Time, newMessage.Message))

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP_SERVER"), 587, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		errRes.Code = 500
		errRes.Message = "internal server error"
		c.JSON(500, errRes)
		return
	}
	res.Code = 200
	res.Message = "Message is sent successfully"
	res.Data = ""
	c.JSON(200, res)
}

func validateRequest(data models.ContactForm) bool {

	if data.Email == "" || data.Name == "" || data.Inquiry == "" || data.Message == "" || inquiryMap[strings.ToLower(data.Inquiry)] == "" {
		return false
	}
	if helper.ValidateEmail(data.Email) != nil {
		return false
	}
	return true

}
