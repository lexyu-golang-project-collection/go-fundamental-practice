package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	a := []string{"a", "b", "c"}
	// +
	res1 := a[0] + a[1] + a[2]
	fmt.Println("+ concat:", res1)

	// fmt.Sprintf
	res2 := fmt.Sprintf("%s%s%s", a[0], a[1], a[2])
	fmt.Println("fmt.Sprintf concat:", res2)

	// strings.Builder
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	res3 := sb.String()
	fmt.Println("strings.Builder concat:", res3)

	// bytes.Buffer
	buf := new(bytes.Buffer)
	buf.Write([]byte(a[0]))
	buf.Write([]byte(a[1]))
	buf.Write([]byte(a[2]))
	res4 := buf.String()
	fmt.Println("bytes.Buffer concat:", res4)

	// strings.Join
	res5 := strings.Join(a, "")
	fmt.Println("strings.Join concat:", res5)

}
