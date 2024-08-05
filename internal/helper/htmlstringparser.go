package helper

import "regexp"

func ParseHTML(htmlString string) string {
	firstStatement := regexp.MustCompile(`%`)
	secondStatement := regexp.MustCompile(`%s`)

	htmlString = firstStatement.ReplaceAllString(htmlString, "%%")
	htmlString = secondStatement.ReplaceAllString(htmlString, "s")
	return htmlString
}
