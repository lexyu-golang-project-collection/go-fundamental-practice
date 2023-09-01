package regex

import (
	"regexp"

	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"
)

func Regex() {
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)

	g.PL(match)
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("[crmfp]at")
	g.PL("MatchString =", r.MatchString(reStr2))
	g.PL("FindString =", r.FindString(reStr2))
	g.PL("Index =", r.FindStringIndex(reStr2))
	g.PL("All String =", r.FindAllString(reStr2, -1))
	g.PL("1st 2 Strings =", r.FindAllString(reStr2, 2))
	g.PL("All Submatch Index =", r.FindAllStringSubmatchIndex(reStr2, -1))
	g.PL(r.ReplaceAllString(reStr2, "Dog"))
}
