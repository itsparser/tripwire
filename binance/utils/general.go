package utils

import "time"

func CurrentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}
