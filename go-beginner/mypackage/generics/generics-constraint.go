package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/mypackage/global"

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func Generics() {
	g.PL("5 + 4 =", getSumGen(5, 4))
	g.PL("5 + 4 =", getSumGen(5.6, 4.7))
}
