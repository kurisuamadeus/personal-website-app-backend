package helper

import (
	"fmt"
	"strconv"
	"time"
)

func FormatMessageId(inquiryType string, messageCount int64) string {

	return inquiryType + fmt.Sprintf("%02d", int(time.Now().Month())) + formatYear(time.Now().Year()) + fmt.Sprintf("%06d", messageCount)

}

func formatYear(year int) string {
	return strconv.Itoa(year)[2:]
}
