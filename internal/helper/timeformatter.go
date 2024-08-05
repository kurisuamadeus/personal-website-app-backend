package helper

import (
	"fmt"
	"time"
)

func FormatTime(hour int, minute int, second int) string {
	timeZone, _ := time.Now().Zone()
	return fmt.Sprintf("%02d:%02d:%02d %s", hour, minute, second, timeZone)
}
