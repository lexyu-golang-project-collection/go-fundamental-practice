package main

import (
	"fmt"
)

func main() {
	// var myMap map [keyType]valueType
	var heroes map[string]string
	heroes = make(map[string]string)
	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"

	villians["Lex"] = "Luthor"

	superPets := map[int]string{1: "Krypto", 2: "Bat Cow"}
	fmt.Printf("Batman is %v\n", heroes["Batman"])
	fmt.Println("Chip :", superPets[3])
	_, ok := superPets[3]
	fmt.Println("Is there a 3rd pet :", ok)
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
	delete(heroes, "The Flash")
}
