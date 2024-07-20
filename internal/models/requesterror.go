package models

type RequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
