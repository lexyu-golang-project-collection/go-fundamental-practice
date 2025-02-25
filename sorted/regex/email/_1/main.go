package main

import (
	"fmt"
	"regexp"
)

var (
	BasicEmailPattern      = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\.[a-zA-Z]{2,})*$`)
	RFC5322EmailPattern    = regexp.MustCompile(`(?i)^(?:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$`)
	SimplifiedEmailPattern = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	UnicodeEmailPattern    = regexp.MustCompile(`(?i)^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	HTML5EmailPattern      = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
)

func IsBasicEmail(email string) bool {
	return BasicEmailPattern.MatchString(email)
}

func IsRFC5322Email(email string) bool {
	return RFC5322EmailPattern.MatchString(email)
}

func IsSimplifiedEmail(email string) bool {
	return SimplifiedEmailPattern.MatchString(email)
}

func IsUnicodeEmail(email string) bool {
	return UnicodeEmailPattern.MatchString(email)
}

func IsHTML5Email(email string) bool {
	return HTML5EmailPattern.MatchString(email)
}

func testPattern(name string, pattern *regexp.Regexp, validEmails, invalidEmails []string) {
	fmt.Printf("===== %s Pattern =====\n", name)

	fmt.Println("有效的電子郵件:")
	for _, email := range validEmails {
		result := pattern.MatchString(email)
		fmt.Printf("  %-40s: %v\n", email, result)
	}

	fmt.Println("\n無效的電子郵件:")
	for _, email := range invalidEmails {
		result := pattern.MatchString(email)
		fmt.Printf("  %-40s: %v\n", email, result)
	}
	fmt.Println()
}

func main() {
	patterns := map[string]*regexp.Regexp{
		"Basic":      BasicEmailPattern,
		"RFC5322":    RFC5322EmailPattern,
		"Simplified": SimplifiedEmailPattern,
		"Unicode":    UnicodeEmailPattern,
		"HTML5":      HTML5EmailPattern,
	}

	basicValid := []string{
		"simple@example.com",
		"very.common@example.com",
		"disposable.style.email.with+tag@example.com",
		"other.email-with-hyphen@example.com",
		"fully-qualified-domain@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"example-indeed@strange-example.com",
	}

	basicInvalid := []string{
		"üser@example.com",                            // Unicode 字元
		"Abc.example.com",                             // 沒有 @ 符號
		"A@b@c@example.com",                           // 多個 @ 符號
		"a\"b(c)d,e:f;g<h>i[j\\k]l@example.com",       // 特殊字元
		"just\"not\"right@example.com",                // 引號
		"this is\"not\\allowed@example.com",           // 含空格
		"i.like.underscores@but_they.are.not.allowed", // 下劃線在域名部分
		"user@example",                                // 缺少 TLD
		"user@example.c",                              // TLD 太短
		"user@[123.123.123.123]",                      // IP 地址形式
	}

	rfc5322Valid := []string{
		"simple@example.com",
		"very.common@example.com",
		"disposable.style.email.with+tag@example.com",
		"other.email-with-hyphen@example.com",
		"x@example.com",
		"\"much.more unusual\"@example.com",
		"\"very.unusual.@.unusual.com\"@example.com",
		"\"very.(),:;<>[]\".VERY.\"very@\\ \"very\".unusual\"@strange.example.com",
		"example-indeed@strange-example.com",
		"admin@mailserver1",                // 無頂級域名
		"user@[IPv6:2001:DB8::1]",          // IPv6 地址
		"user@[192.168.2.1]",               // IPv4 地址
		"\"escaped\\\"quote\"@example.com", // 轉義字元
		"user@example.c",                   // 單字元 TLD
	}

	rfc5322Invalid := []string{
		"Abc.example.com",                            // 沒有 @ 符號
		"A@b@c@example.com",                          // 多個 @ 符號
		"a\"b(c)d,e:f;g<h>i[j\\k]l@example.com",      // 特殊字元組合
		"this\\ still\\\"not\\\\allowed@example.com", // 轉義字元錯誤使用
		"1234567890123456789012345678901234567890123456789012345678901234+x@example.com",                   // 本地部分過長
		"i_like_underscore@but_this_is_too_long_for_a_domain_part_as_it_exceeds_sixty_four_characters.com", // 域名部分過長
		"user@example..com", // 連續的點
		"@example.com",      // 缺少本地部分
		"user@",             // 缺少域名部分
	}

	simplifiedValid := []string{
		"simple@example.com",
		"very.common@example.com",
		"üser@example.com",                            // Unicode 字元
		"user@example.c",                              // 單字元 TLD
		"user@[123.123.123.123]",                      // IP 形式
		"i_like_underscores@but_they.are.not.allowed", // 特殊格式
		"user@localhost",                              // 本地主機
		"user@123.123.123.123",                        // IP 地址
		"user@sub.sub.sub.example.org",                // 多層子域名
	}

	simplifiedInvalid := []string{
		"Abc.example.com",       // 沒有 @ 符號
		"A@b@c@example.com",     // 多個 @ 符號
		"user name@example.com", // 含空格
		"user@example",          // 缺少點
		"@example.com",          // 缺少本地部分
		"user@",                 // 缺少域名部分
	}

	unicodeValid := []string{
		"simple@example.com",
		"very.common@example.com",
		"üser@example.com",       // Unicode 字元
		"用戶@例子.com",              // 中文字元
		"ユーザ@例.jp",               // 日文字元
		"θσερ@εχαμπλε.ψομ",       // 希臘字元
		"user@[123.123.123.123]", // IP 形式
		"\"quoted\"@example.com", // 引號
		"user@example.c",         // 單字元 TLD
	}

	unicodeInvalid := []string{
		"Abc.example.com",       // 沒有 @ 符號
		"A@b@c@example.com",     // 多個 @ 符號
		"user name@example.com", // 含空格
		"user@example..com",     // 連續的點
		"@example.com",          // 缺少本地部分
		"user@",                 // 缺少域名部分
	}

	html5Valid := []string{
		"simple@example.com",
		"very.common@example.com",
		"disposable.style.email.with+tag@example.com",
		"other.email-with-hyphen@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"example-indeed@strange-example.com",
		"admin@mailserver1",        // 無頂級域名
		"user@example.c",           // 單字元 TLD
		"user@example.photography", // 長 TLD
	}

	html5Invalid := []string{
		"üser@example.com",                      // Unicode 字元
		"Abc.example.com",                       // 沒有 @ 符號
		"A@b@c@example.com",                     // 多個 @ 符號
		"a\"b(c)d,e:f;g<h>i[j\\k]l@example.com", // 特殊字元組合
		"just\"not\"right@example.com",          // 引號
		"this is\"not\\allowed@example.com",     // 含空格
		"user@[123.123.123.123]",                // IP 地址形式
		"user@example..com",                     // 連續的點
		"@example.com",                          // 缺少本地部分
		"user@",                                 // 缺少域名部分
	}

	testPattern("Basic", patterns["Basic"], basicValid, basicInvalid)
	testPattern("RFC5322", patterns["RFC5322"], rfc5322Valid, rfc5322Invalid)
	testPattern("Simplified", patterns["Simplified"], simplifiedValid, simplifiedInvalid)
	testPattern("Unicode", patterns["Unicode"], unicodeValid, unicodeInvalid)
	testPattern("HTML5", patterns["HTML5"], html5Valid, html5Invalid)
}
