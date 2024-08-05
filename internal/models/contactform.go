package models

type ContactForm struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Inquiry string `json:"inquiry"`
	Message string `json:"message"`
}
