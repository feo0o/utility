package utility

import (
	"fmt"
	"time"
)

// TimeNowReadable return the current local time with format "2006-01-02 15:04:05" as string
func TimeNowReadable() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimeNowUnixString return the current local time with Unix second as string
func TimeNowUnixString() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// TimeNowUnixNano return the current local time with Unix nano second as string
func TimeNowUnixNanoString() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// TimeAddToUnixString add "t" into base time, return result as Unix second string
func TimeAddUnixString(base time.Time, t time.Duration) string {
	return fmt.Sprintf("%d", base.Add(t).Unix())
}
