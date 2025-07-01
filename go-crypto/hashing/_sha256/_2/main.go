package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// HeaderSignature 包含所有必要的 header 欄位
type HeaderSignature struct {
	ClientId   string
	SystemKind string
	Platform   string
	Timestamp  int64
	Source     string
	SecretKey  string
}

// GenerateSignature 生成 SHA256 簽名
func (h *HeaderSignature) GenerateSignature() string {
	// 按照指定順序組合字串：X-ClientId + X-System-Kind + X-Platform + X-Timestamp + X-Source + SecretKey
	data := fmt.Sprintf("%s%s%s%d%s%s",
		h.ClientId,
		h.SystemKind,
		h.Platform,
		h.Timestamp,
		h.Source,
		h.SecretKey,
	)

	// 計算 SHA256 雜湊值
	hash := sha256.Sum256([]byte(data))

	// 轉換為十六進位字串
	return fmt.Sprintf("%x", hash)
}

// VerifySignature 驗證傳入的簽名是否正確
func (h *HeaderSignature) VerifySignature(providedSignature string) bool {
	expectedSignature := h.GenerateSignature()
	return expectedSignature == providedSignature
}

// HTTPHeaders 模擬 HTTP headers 結構
type HTTPHeaders struct {
	XClientId   string `header:"X-ClientId"`
	XSystemKind string `header:"X-System-Kind"`
	XPlatform   string `header:"X-Platform"`
	XTimestamp  int64  `header:"X-Timestamp"`
	XSource     string `header:"X-Source"`
	XSignature  string `header:"X-Signature"`
}

// ValidateHeaders 驗證 HTTP headers 中的簽名
func ValidateHeaders(headers HTTPHeaders, secretKey string) (bool, error) {
	headerSig := HeaderSignature{
		ClientId:   headers.XClientId,
		SystemKind: headers.XSystemKind,
		Platform:   headers.XPlatform,
		Timestamp:  headers.XTimestamp,
		Source:     headers.XSource,
		SecretKey:  secretKey,
	}

	isValid := headerSig.VerifySignature(headers.XSignature)
	if !isValid {
		return false, fmt.Errorf("invalid signature")
	}

	return true, nil
}

func main() {
	secretKey := "87a78d62c99b44b3ca1309004000fe3d66eec505" // 40 characters; 87a78d62c99b44b3ca1309004000fe3d66eec505; abcdefghijklmnopqrstuvwxyz1234567890123456

	headerSig := HeaderSignature{
		ClientId:   "12345678",
		SystemKind: "CLIENT",
		Platform:   "SERVER",
		Timestamp:  1751341000,
		Source:     "CLIENT",
		SecretKey:  secretKey,
	}

	signature := headerSig.GenerateSignature()
	fmt.Printf("Generated Signature: %s\n", signature)

	headers := HTTPHeaders{
		XClientId:   "12345678",
		XSystemKind: "CLIENT",
		XPlatform:   "SERVER",
		XTimestamp:  1751341000,
		XSource:     "CLIENT",
		XSignature:  signature,
	}

	isValid, err := ValidateHeaders(headers, secretKey)
	if err != nil {
		fmt.Printf("Validation Error: %v\n", err)
		return
	}

	fmt.Printf("Signature Valid: %t\n", isValid)

	currentTimestamp := time.Now().Unix()
	fmt.Printf("Current Timestamp: %d\n", currentTimestamp)
}
