package mypackage

func Slice() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		pf("Rune Array : %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("bStr = ", bStr)

	// var name []datatype
	sl1 := make([]string, 5)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	pl("sl1 Slice = ", sl1)
	pl("sl1 Slice Size = ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pf("sl1[%d] = %s\n", i, sl1[i])
	}
	for _, x := range sl1 {
		pl("x = ", x)
	}
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl("sl3 :", sl3)
	pl("First 3 sArr[:3] = ", sArr[:3])
	pl("Last 3 sArr[2:] = ", sArr[2:])

	sArr[0] = 10
	pl("sArr = ", sArr)
	pl("sl3 = ", sl3)

	sl3[0] = 1
	pl("sArr = ", sArr)
	pl("sl3 = ", sl3)

	sl3 = append(sl3, 12)
	pl("sl3 = ", sl3)
	pl("sArr = ", sArr)

	sl4 := make([]string, 6)
	pl("sl4 = ", sl4)
	pl("sl4[0] = ", sl4[0])
}
