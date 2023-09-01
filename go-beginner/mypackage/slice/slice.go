package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"

func Slice() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		g.PF("Rune Array : %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	g.PL("bStr = ", bStr)

	// var name []datatype
	sl1 := make([]string, 5)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	g.PL("sl1 Slice = ", sl1)
	g.PL("sl1 Slice Size = ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		g.PF("sl1[%d] = %s\n", i, sl1[i])
	}
	for _, x := range sl1 {
		g.PL("x = ", x)
	}
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	g.PL("sl3 :", sl3)
	g.PL("First 3 sArr[:3] = ", sArr[:3])
	g.PL("Last 3 sArr[2:] = ", sArr[2:])

	sArr[0] = 10
	g.PL("sArr = ", sArr)
	g.PL("sl3 = ", sl3)

	sl3[0] = 1
	g.PL("sArr = ", sArr)
	g.PL("sl3 = ", sl3)

	sl3 = append(sl3, 12)
	g.PL("sl3 = ", sl3)
	g.PL("sArr = ", sArr)

	sl4 := make([]string, 6)
	g.PL("sl4 = ", sl4)
	g.PL("sl4[0] = ", sl4[0])
}
