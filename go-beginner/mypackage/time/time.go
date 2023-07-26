package mypackage

import (
	"time"

	g "github.com/lexyu/go-beginner/mypackage/global"
)

func Time() {
	now := time.Now()
	g.PL(now.Local())
	g.PL(now.Year(), now.Month(), now.Day())
	g.PL(now.Hour(), now.Minute(), now.Second())
}
