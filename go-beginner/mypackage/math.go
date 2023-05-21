package mypackage

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Math() {
	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4)

	pl("Float Precision = ", 0.1111111111111111+0.1111111111111111) // 16

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)
	pl("Abs(-10)", math.Abs(-10))
	pl("Pow(4,2)", math.Pow(4, 2))
	pl("Sqrt(16)", math.Sqrt(16))
	pl("Cbrt(8)", math.Cbrt(8))
	pl("Ceil(4.4)", math.Ceil(4.4))
	pl("Floor(4.4)", math.Floor(4.4))
	pl("Round(4.4)", math.Round(4.4))
	pl("Log2(8)", math.Log2(8))
	pl("Log10(100)", math.Log10(100))

	// log of e to the power of 2
	pl("Log(7.389)", math.Log(7.389))
	pl("Max(5,4)", math.Max(5, 4))
	pl("Min(5,4)", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	pf("%.3f radians = %.3f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))
	// Cos, Tan, Acos, Asin, Atan, Asinh, Atanh, Atan2, Cosh, Sinh, Sincos, Hypot

	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value
	pf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	pf("%9f\n", 3.14)
	pf("%.2f\n", 3.141592)
	pf("%9.f\n", 3.141592)
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}
