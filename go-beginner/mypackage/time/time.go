package mypackage

import (
	"time"

	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/mypackage/global"
)

func Time() {
	now := time.Now()
	g.PL(now.Local())
	g.PL(now.Year(), now.Month(), now.Day())
	g.PL(now.Hour(), now.Minute(), now.Second())
}
