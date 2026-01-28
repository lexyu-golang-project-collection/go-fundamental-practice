package main

import (
	"fmt"
	"regexp"
)

func main() {
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)
	fmt.Println(match)

	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("[crmfp]at")

	fmt.Println("MatchString =", r.MatchString(reStr2))
	fmt.Println("FindString =", r.FindString(reStr2))
	fmt.Println("Index =", r.FindStringIndex(reStr2))
	fmt.Println("All String =", r.FindAllString(reStr2, -1))
	fmt.Println("1st 2 Strings =", r.FindAllString(reStr2, 2))
	fmt.Println("All Submatch Index =", r.FindAllStringSubmatchIndex(reStr2, -1))
	fmt.Println(r.ReplaceAllString(reStr2, "Dog"))

}
