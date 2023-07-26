package mypackage

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	g "github.com/lexyu/go-beginner/mypackage/global"
)

func Math() {
	g.PL("5 + 4 = ", 5+4)
	g.PL("5 - 4 = ", 5-4)
	g.PL("5 * 4 = ", 5*4)
	g.PL("5 / 4 = ", 5/4)
	g.PL("5 % 4 = ", 5%4)

	g.PL("Float Precision = ", 0.1111111111111111+0.1111111111111111) // 16

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	g.PL("Random :", randNum)
	g.PL("Abs(-10)", math.Abs(-10))
	g.PL("Pow(4,2)", math.Pow(4, 2))
	g.PL("Sqrt(16)", math.Sqrt(16))
	g.PL("Cbrt(8)", math.Cbrt(8))
	g.PL("Ceil(4.4)", math.Ceil(4.4))
	g.PL("Floor(4.4)", math.Floor(4.4))
	g.PL("Round(4.4)", math.Round(4.4))
	g.PL("Log2(8)", math.Log2(8))
	g.PL("Log10(100)", math.Log10(100))

	// log of e to the power of 2
	g.PL("Log(7.389)", math.Log(7.389))
	g.PL("Max(5,4)", math.Max(5, 4))
	g.PL("Min(5,4)", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	g.PF("%.3f radians = %.3f degrees\n", r90, d90)
	g.PL("Sin(90) =", math.Sin(r90))
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
	g.PF("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	g.PF("%9f\n", 3.14)
	g.PF("%.2f\n", 3.141592)
	g.PF("%9.f\n", 3.141592)
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	g.PL(sp1)
}
