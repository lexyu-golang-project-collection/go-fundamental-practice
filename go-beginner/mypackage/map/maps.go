package mypackage

import (
	"fmt"
	g "github.com/lexyu/go-beginner/mypackage/global"
)

func Maps() {
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
	g.PL("Chip :", superPets[3])
	_, ok := superPets[3]
	g.PL("Is there a 3rd pet :", ok)
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
	delete(heroes, "The Flash")
}
