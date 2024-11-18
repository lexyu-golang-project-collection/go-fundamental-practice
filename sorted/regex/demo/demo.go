package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 【字元匹配】測試
	fmt.Println("=== 字元匹配測試 ===")

	// . 測試
	dotPattern := regexp.MustCompile(`h.t`)
	fmt.Printf("h.t -> hot: %v\n", dotPattern.MatchString("hot"))
	fmt.Printf("h.t -> hit: %v\n", dotPattern.MatchString("hit"))
	fmt.Printf("h.t -> h@t: %v\n", dotPattern.MatchString("h@t"))

	// \d 測試
	digitPattern := regexp.MustCompile(`\d{3}`)
	fmt.Printf("\\d{3} -> 123: %v\n", digitPattern.MatchString("123"))
	fmt.Printf("\\d{3} -> 456: %v\n", digitPattern.MatchString("456"))
	fmt.Printf("\\d{3} -> 12: %v\n", digitPattern.MatchString("12")) // false

	// \D 測試
	nonDigitPattern := regexp.MustCompile(`\D+`)
	fmt.Printf("\\D+ -> abc: %v\n", nonDigitPattern.MatchString("abc"))
	fmt.Printf("\\D+ -> xyz: %v\n", nonDigitPattern.MatchString("xyz"))

	// \w 測試
	wordPattern := regexp.MustCompile(`\w+`)
	fmt.Printf("\\w+ -> abc123: %v\n", wordPattern.MatchString("abc123"))
	fmt.Printf("\\w+ -> test_123: %v\n", wordPattern.MatchString("test_123"))

	// \W 測試
	nonWordPattern := regexp.MustCompile(`a\W+b`)
	fmt.Printf("a\\W+b -> a@#b: %v\n", nonWordPattern.MatchString("a@#b"))
	fmt.Printf("a\\W+b -> a  b: %v\n", nonWordPattern.MatchString("a  b"))

	// \s 測試
	spacePattern := regexp.MustCompile(`a\sb`)
	fmt.Printf("a\\sb -> a b: %v\n", spacePattern.MatchString("a b"))
	fmt.Printf("a\\sb -> a\tb: %v\n", spacePattern.MatchString("a\tb"))

	// \S 測試
	nonSpacePattern := regexp.MustCompile(`\S+`)
	fmt.Printf("\\S+ -> abc123: %v\n", nonSpacePattern.MatchString("abc123"))
	fmt.Printf("\\S+ -> test@123: %v\n", nonSpacePattern.MatchString("test@123"))

	// 【位置匹配】測試
	fmt.Println("\n=== 位置匹配測試 ===")

	// ^ 測試
	startPattern := regexp.MustCompile(`^test`)
	fmt.Printf("^test -> test123: %v\n", startPattern.MatchString("test123"))
	fmt.Printf("^test -> 123test: %v\n", startPattern.MatchString("123test"))

	// $ 測試
	endPattern := regexp.MustCompile(`test$`)
	fmt.Printf("test$ -> 123test: %v\n", endPattern.MatchString("123test"))
	fmt.Printf("test$ -> test123: %v\n", endPattern.MatchString("test123"))

	// \b 測試
	wordBoundaryPattern := regexp.MustCompile(`\bcat\b`)
	fmt.Printf("\\bcat\\b -> the cat: %v\n", wordBoundaryPattern.MatchString("the cat"))
	fmt.Printf("\\bcat\\b -> category: %v\n", wordBoundaryPattern.MatchString("category"))

	// 【重複次數】測試
	fmt.Println("\n=== 重複次數測試 ===")

	// * 測試
	starPattern := regexp.MustCompile(`ab*c`)
	fmt.Printf("ab*c -> ac: %v\n", starPattern.MatchString("ac"))
	fmt.Printf("ab*c -> abc: %v\n", starPattern.MatchString("abc"))
	fmt.Printf("ab*c -> abbc: %v\n", starPattern.MatchString("abbc"))

	// + 測試
	plusPattern := regexp.MustCompile(`ab+c`)
	fmt.Printf("ab+c -> abc: %v\n", plusPattern.MatchString("abc"))
	fmt.Printf("ab+c -> abbc: %v\n", plusPattern.MatchString("abbc"))
	fmt.Printf("ab+c -> ac: %v\n", plusPattern.MatchString("ac"))

	// ? 測試
	questionPattern := regexp.MustCompile(`colou?r`)
	fmt.Printf("colou?r -> color: %v\n", questionPattern.MatchString("color"))
	fmt.Printf("colou?r -> colour: %v\n", questionPattern.MatchString("colour"))

	// 【字元類】測試
	fmt.Println("\n=== 字元類測試 ===")

	// [abc] 測試
	charClassPattern := regexp.MustCompile(`[abc]at`)
	fmt.Printf("[abc]at -> bat: %v\n", charClassPattern.MatchString("bat"))
	fmt.Printf("[abc]at -> cat: %v\n", charClassPattern.MatchString("cat"))
	fmt.Printf("[abc]at -> dat: %v\n", charClassPattern.MatchString("dat"))

	// [^abc] 測試
	negCharClassPattern := regexp.MustCompile(`[^abc]at`)
	fmt.Printf("[^abc]at -> dat: %v\n", negCharClassPattern.MatchString("dat"))
	fmt.Printf("[^abc]at -> bat: %v\n", negCharClassPattern.MatchString("bat"))

	// 【分組和邏輯運算】測試
	fmt.Println("\n=== 分組和邏輯運算測試 ===")

	// (pattern) 測試
	groupPattern := regexp.MustCompile(`(ab)+`)
	fmt.Printf("(ab)+ -> ab: %v\n", groupPattern.MatchString("ab"))
	fmt.Printf("(ab)+ -> abab: %v\n", groupPattern.MatchString("abab"))

	// | 測試
	orPattern := regexp.MustCompile(`cat|dog`)
	fmt.Printf("cat|dog -> cat: %v\n", orPattern.MatchString("cat"))
	fmt.Printf("cat|dog -> dog: %v\n", orPattern.MatchString("dog"))
	fmt.Printf("cat|dog -> rat: %v\n", orPattern.MatchString("rat"))
}
