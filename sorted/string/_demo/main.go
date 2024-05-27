package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	str2 := replacer.Replace(str1)

	fmt.Println("str2 = ", str2)
	fmt.Println("sV2 Length :", len(str2))
	fmt.Println("sV2 Contains Another :", strings.Contains(str2, "Another"))
	fmt.Println("sV2 o index :", strings.Index(str2, "o"))
	fmt.Println("Replace :", strings.Replace(str2, "o", "0", -1))

	str3 := "\nSome Words\n" // \t \" \\
	str3 = strings.TrimSpace(str3)

	fmt.Println("str3 = ", str3)
	fmt.Println("Split :", strings.Split("a-b-c-d", "-"))
	fmt.Println("str3 Lower :", strings.ToLower(str2))
	fmt.Println("str3 Upper :", strings.ToUpper(str2))
	fmt.Println("str3 Prefix :", strings.HasPrefix("tacocat", "taco"))
	fmt.Println("str3 Prefix :", strings.HasSuffix("tacocat", "cat"))
	fmt.Println("str3 Lower :", strings.ToLower(str3))
	fmt.Println("str3 Upper :", strings.ToUpper(str3))
	fmt.Println("str3 Prefix :", strings.HasPrefix("tacocat", "taco"))
	fmt.Println("str3 Prefix :", strings.HasSuffix("tacocat", "cat"))
}
