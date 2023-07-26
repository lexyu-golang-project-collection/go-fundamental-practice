package mypackage

import (
	"strings"

	g "github.com/lexyu/go-beginner/mypackage/global"
)

func Condtional() {
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		g.PL("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		g.PL("Important Birthday")
	} else if iAge >= 65 {
		g.PL("Important Birthday")
	} else {
		g.PL("Not an Important Birthday")
	}

	g.PL("!true = ", !true)

	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	g.PL("sV2 = ", sV2)
	g.PL("sV2 Length :", len(sV2))
	g.PL("sV2 Contains Another :", strings.Contains(sV2, "Another"))
	g.PL("sV2 o index :", strings.Index(sV2, "o"))
	g.PL("Replace :", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words\n" // \t \" \\
	sV3 = strings.TrimSpace(sV3)
	g.PL("sV3 = ", sV3)
	g.PL("Split :", strings.Split("a-b-c-d", "-"))
	g.PL("sV2 Lower :", strings.ToLower(sV2))
	g.PL("sV2 Upper :", strings.ToUpper(sV2))
	g.PL("sV2 Prefix :", strings.HasPrefix("tacocat", "taco"))
	g.PL("sV2 Prefix :", strings.HasSuffix("tacocat", "cat"))

	g.PL("sV3 Lower :", strings.ToLower(sV3))
	g.PL("sV3 Upper :", strings.ToUpper(sV3))
	g.PL("sV3 Prefix :", strings.HasPrefix("tacocat", "taco"))
	g.PL("sV3 Prefix :", strings.HasSuffix("tacocat", "cat"))
}
