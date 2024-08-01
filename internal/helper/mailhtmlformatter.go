package helper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetFormattedHTMLMessage(header string, messageId string, inquiry string, name string, email string, date string, clock string, message string) string {
	path, err := filepath.Abs("./internal/mailformathtml/mailformat.html")
	if err != nil {

		return ""
	}
	file, err := os.Open(path)
	if err != nil {

		return ""
	}
	defer file.Close()
	byteData, _ := io.ReadAll(file)
	var result string = string(byteData[:])
	return fmt.Sprintf(ParseHTML(result), header, messageId, inquiry, name, email, date, clock, message)
}
