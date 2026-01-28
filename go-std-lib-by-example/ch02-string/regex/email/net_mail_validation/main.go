package main

import (
	"fmt"
	"net/mail"
)

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main() {
	emails := []string{
		"user@example.com",            // 基本格式
		"user.name@example.com",       // 本地部分含點
		"user+tag@example.com",        // 本地部分含加號
		"\"quoted name\"@example.com", // 引號包含的本地部分
		"user@localhost",              // 本地域名
		"user@example..com",           // 域名含連續點
		"user name@example.com",       // 本地部分含空格
		"user@",                       // 缺少域名
		"@example.com",                // 缺少本地部分
		"not-an-email",                // 完全不是郵件格式
	}

	for _, email := range emails {
		valid := validateEmail(email)
		if valid {
			fmt.Printf("%-30s: 有效\n", email)
		} else {
			fmt.Printf("%-30s: 無效\n", email)
		}
	}

	// 更詳細的錯誤處理示例
	email := "invalid@@example.com"
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Printf("\n詳細錯誤信息: %s\n", err)
	}
}
