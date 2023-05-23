package mypackage

import "fmt"

type Tsp float64
type Tbs float64
type ML float64

/*
1 Tbs/Tbsp (大匙) ≒ 15 ml
1 tsp (茶匙) ≒ 5 ml
*/

func tspToML(tsp Tsp) ML {
	return ML(tsp * 5)
}

func TbToML(tsp Tbs) ML {
	return ML(tsp * 15)
}
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 5)
}
func (tbs Tbs) ToMLs() ML {
	return ML(tbs * 15)
}

func Defined_Types_Associate_Methods() {
	ml1 := ML(Tsp(3) * 5)
	fmt.Printf("3 tsps = %.2f ML\n", ml1)
	ml2 := ML(Tbs(3) * 15)
	fmt.Printf("3 Tbs = %.2f ML\n", ml2)
	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp + 4 tsp =", Tsp(2) > Tsp(4))
	fmt.Printf("3 tsp = %.2f ml\n", tspToML(3))
	fmt.Printf("3 tsp = %.2f ml\n", TbToML(3))

	// Associate Methods
	tsp := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f ml\n", tsp, tsp.ToMLs())
}
