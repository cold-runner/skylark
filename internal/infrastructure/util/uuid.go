package util

import (
	"time"
)

func GenUnixTimeNano() int64 {
	return time.Now().UnixNano()
}
