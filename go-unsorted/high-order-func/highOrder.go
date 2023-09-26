package main

import "fmt"

func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

func printResult(radius float64, calFunc func(r float64) float64) {
	result := calFunc(radius)
	fmt.Println("Result: ", result)
}

func getFunction(selected int) func(r float64) float64 {
	selected_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return selected_to_func[selected]
}

func main() {

	var selected int
	var radius float64

	fmt.Print("Enter the radius: ")
	fmt.Scanf("%f", &radius)
	fmt.Scanln()
	fmt.Print("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
	fmt.Scanf("%d\n", &selected)

	printResult(radius, getFunction(selected))

}
