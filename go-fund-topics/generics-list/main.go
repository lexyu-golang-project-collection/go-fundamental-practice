package main

import "fmt"

func main() {
	glist := NewGemericList[string]()

	glist.Insert("Bug")
	glist.Insert("Foo")
	glist.Insert("Bar")
	glist.Insert("Unknown")

	glist.RemoveByValue("Foo")
	glist.Remove(2)

	fmt.Printf("%+v\n", glist)
}
