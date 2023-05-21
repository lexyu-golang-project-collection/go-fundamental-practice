package mypackage

import "time"

func Time() {
	now := time.Now()
	pl(now.Local())
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}
