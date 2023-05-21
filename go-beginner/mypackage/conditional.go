package mypackage

import "strings"

func Condtional() {
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an Important Birthday")
	}

	pl("!true = ", !true)

	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl("sV2 = ", sV2)
	pl("sV2 Length :", len(sV2))
	pl("sV2 Contains Another :", strings.Contains(sV2, "Another"))
	pl("sV2 o index :", strings.Index(sV2, "o"))
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words\n" // \t \" \\
	sV3 = strings.TrimSpace(sV3)
	pl("sV3 = ", sV3)
	pl("Split :", strings.Split("a-b-c-d", "-"))
	pl("sV2 Lower :", strings.ToLower(sV2))
	pl("sV2 Upper :", strings.ToUpper(sV2))
	pl("sV2 Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("sV2 Prefix :", strings.HasSuffix("tacocat", "cat"))

	pl("sV3 Lower :", strings.ToLower(sV3))
	pl("sV3 Upper :", strings.ToUpper(sV3))
	pl("sV3 Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("sV3 Prefix :", strings.HasSuffix("tacocat", "cat"))
}
