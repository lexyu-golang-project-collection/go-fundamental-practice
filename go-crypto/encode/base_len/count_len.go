package main

import (
	"fmt"
	"math"
)

// L: 原始資料長度（以 bytes 計算）
func EstimateBase16Length(L int) int {
	return L * 2
}

func EstimateBase32Length(L int) int {
	return int(math.Ceil(float64(L) * 8 / 5))
}

func EstimateBase64Length(L int) int {
	return ((L + 2) / 3) * 4 // padding
}

func EstimateBase62Length(L int) int {
	// 8 bits / byte，估計轉 base62 長度
	// log₂(62) ≈ 5.954，log₂(256) = 8
	// 所以每 byte ≈ 8 / 5.954 ≈ 1.34 個 base62 字元
	return int(math.Ceil(float64(L) * 8 / math.Log2(62)))
}

func main() {
	dataLen := 3

	fmt.Println("原始資料長度:", dataLen, "bytes")
	fmt.Println("Base16 編碼長度:", EstimateBase16Length(dataLen))
	fmt.Println("Base32 編碼長度:", EstimateBase32Length(dataLen))
	fmt.Println("Base64 編碼長度:", EstimateBase64Length(dataLen))
	fmt.Println("Base62 編碼長度:", EstimateBase62Length(dataLen))
}
