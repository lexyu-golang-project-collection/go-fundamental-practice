package mypackage

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func Generics() {
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5 + 4 =", getSumGen(5.6, 4.7))
}
