package mypackage

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"

	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"
)

func Array() {
	var arr1 [5]int
	arr1[0] = 1
	g.PL("arr1 Index 0 = ", arr1[0])

	arr2 := [5]int{1, 2, 3, 4, 5}
	g.PL("arr2 Index 0 = ", arr2[0])
	g.PL("arr2 Length 0 =", len(arr2))
	for i := 0; i < len(arr2); i++ {
		g.PF("arr2[%d] = %d\n", i, arr2[i])
	}
	for i, v := range arr2 {
		g.PF("arr2 index %d : value %d\n", i, v)
	}

	arr3 := [2][2]int{
		{5, 6},
		{7, 8},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			g.PF("arr3[%d][%d] = %d\n", i, j, arr3[i][j])
		}
	}

	g.PL("----------------------------------------------------------------")

	p := []int{1, 2, 3}
	qwe := 0
	qwe, p[qwe] = 1, 2
	g.PL("Exchange ", p, qwe)

	g.PL("----------------------------------------------------------------")
	var arr []int
	fmt.Printf("&arr = %p\n", &arr)
	b := []int{1, 2}
	fmt.Printf("&b = %p\n", &b)
	arr = append(arr, 100)
	arr = append(arr, b...) // append another array
	fmt.Println("arr = ", arr)
	fmt.Printf("&arr = %p\n", &arr)

	fmt.Println("--------------------------------------------------------------")

	var c [10]int // default value 0
	fmt.Println("c = ", c)

	fmt.Println("--------------------------------------------------------------")

	var d []int
	e := []int{1, 2, 3, 4}
	d = append(e[:1], e[2:]...) // [:1] new slice without values after index 1, [2:] include index 2
	fmt.Println("d = ", d)

	fmt.Println("--------------------------------------------------------------")

	// slice
	var f [][]int
	g := []int{1, 2}
	f = append(f, g)
	fmt.Println("f = ", f)

	fmt.Println("--------------------------------------------------------------")

	// array
	var h [][2]int
	j := [2]int{1, 2}
	h = append(h, j)
	fmt.Println("h = ", h)

	fmt.Println("--------------------------------------------------------------")

	var i [][3]int
	// k := [...]int{1, 2}
	// i = append(i, k) // error
	fmt.Println("i = ", i)

	fmt.Println("--------------------------------------------------------------")

	var x [][]int
	fmt.Println("x = ", x)
	fmt.Printf("&x = %p\n", &x)
	y := []int{1, 2}
	x = append(x, y)
	fmt.Println("x = ", x)
	y[1] = 3
	fmt.Printf("&x = %p\n", &x)

	fmt.Println("--------------------------------------------------------------")

	var l [][]int
	fmt.Printf("&l = %p\n", &l)
	q := []int{1, 2}
	fmt.Printf("&q = %p\n", &q)
	// m := make([]int, 2)
	m := q
	fmt.Printf("&m = %p\n", &m)
	// copy(m, q)
	l = append(l, m)
	fmt.Println("l = ", l)
	fmt.Printf("&l = %p\n", &l)
	q[1] = 3
	fmt.Println("l = ", l)
	fmt.Printf("&l = %p\n", &l)
	fmt.Println("m = ", m)
	fmt.Printf("&m = %p\n", &m)

	fmt.Println("--------------------------------------------------------------")

	// 1st Choice
	var a []int
	timeStart := time.Now()
	for i := 0; i < 1_000_000; i++ {
		a = append(a, i)
	}
	fmt.Println(time.Since(timeStart))

	// 2nd Choice
	b = make([]int, 0, 1_000_000)
	timeStart = time.Now()
	for i := 0; i < 1_000_000; i++ {
		b = append(b, i)
	}
	fmt.Println(time.Since(timeStart))

	// 3rd Choice
	z := make([]int, 1_000_000)
	timeStart = time.Now()
	for i := 0; i < 1_000_000; i++ {
		z = append(z, i)
	}
	fmt.Println(time.Since(timeStart))

	fmt.Println("--------------------------------------------------------------")

	scores := []int{1, 2, 3, 4, 5}
	fmt.Printf("&scores = %p\n", &scores)
	scoresArrayPtr := reflect.ValueOf(scores).Pointer()
	fmt.Printf("Memory address of scores' underlying array: %p\n", unsafe.Pointer(scoresArrayPtr))
	// newSlice := scores
	// newSlice := scores[0:] // `:` newSclice point to same address array
	newSlice := scores[1:]
	newSliceArrayPtr := reflect.ValueOf(newSlice).Pointer()
	fmt.Printf("Memory address of newSlice's underlying array: %p\n", unsafe.Pointer(newSliceArrayPtr))
	fmt.Printf("&newSlice = %p\n", &newSlice)
	// Printing the capacity of both slices
	fmt.Println("scores capacity:", cap(scores))
	fmt.Println("newSlice capacity:", cap(newSlice))
	fmt.Println(newSlice) // 3, 4
	newSlice[0] = 999     // 把原本 scores 中 index 值為 3 的元素改成 999
	fmt.Println(scores)   // 1, 2, 999, 4, 5
}

/* Pass by Referenece
type slice struct{
	p *pointer --> array
	len int64
	cap int64
}
*/
