package main

import (
	"fmt"
	"math"
)

// 計算編碼空間大小
func calculateSpace(base, digits int) int64 {
	return int64(math.Pow(float64(base), float64(digits)))
}

// 簡單計算50%碰撞閾值 (開根號)
func simpleCollisionThreshold(totalSpace int64) int64 {
	return int64(math.Sqrt(float64(totalSpace)))
}

// 格式化大數字顯示
func formatNumber(n int64) string {
	str := fmt.Sprintf("%d", n)
	result := ""
	for i, char := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}
	return result
}

// 分析單個編碼方案
func analyzeScheme(name string, base, digits int) {
	space := calculateSpace(base, digits)
	threshold := simpleCollisionThreshold(space)

	fmt.Printf("\n📊 %s\n", name)
	fmt.Printf("   輸出空間: %s\n", formatNumber(space))
	fmt.Printf("   50%% 碰撞閾值: %s\n", formatNumber(threshold))

	// 安全性檢查
	analyzeUserSafety(threshold)
}

// 分析用戶數量安全性
func analyzeUserSafety(threshold int64) {
	userCounts := []int64{1000, 10000, 50000, 100000}
	fmt.Printf("   用戶數量安全性:\n")

	for _, users := range userCounts {
		status := getUserSafetyStatus(users, threshold)
		fmt.Printf("     %s 用戶: %s\n", formatNumber(users), status)
	}
}

// 取得用戶數量的安全狀態
func getUserSafetyStatus(users, threshold int64) string {
	if users > threshold {
		return "🔴 超過50%風險"
	} else if users > threshold/2 {
		return "🟡 需要注意"
	}
	return "🟢 安全"
}

// 印出方案比較表格
func printComparisonTable(schemes []EncodingScheme) {
	fmt.Println("\n📈 方案比較")
	fmt.Println("=======================================")
	fmt.Printf("%-12s %-15s %-15s\n", "方案", "輸出空間", "50%碰撞閾值")
	fmt.Println("--------------------------------------")

	for _, scheme := range schemes {
		space := calculateSpace(scheme.base, scheme.digits)
		threshold := simpleCollisionThreshold(space)
		fmt.Printf("%-12s | %-15s   | %-15s\n", scheme.name, formatNumber(space), formatNumber(threshold))
	}
}

func printLearningPoints() {
	fmt.Println("\n💡 ")
	fmt.Println("====================================")
	fmt.Println("🔸 生日悖論公式: 50% 碰撞 ≈ √(輸出空間)")
	fmt.Println("🔸 實務建議: 碰撞閾值要遠大於預期用戶數")
}

// 編碼方案
type EncodingScheme struct {
	name   string
	base   int
	digits int
}

func main() {
	fmt.Println("🎯 簡單碰撞分析工具")
	fmt.Println("====================================")

	schemes := []EncodingScheme{
		// Base16 (十六進制)
		{"Base16-4位", 16, 4},
		{"Base16-6位", 16, 6},
		{"Base16-8位", 16, 8},

		// Base32 (RFC 4648)
		{"Base32-5位", 32, 5},
		{"Base32-6位", 32, 6},
		{"Base32-8位", 32, 8},

		// Base62 (常用於短網址)
		{"Base62-3位", 62, 3},
		{"Base62-4位", 62, 4},
		{"Base62-5位", 62, 5},
		{"Base62-6位", 62, 6},
		{"Base62-7位", 62, 7},
		{"Base62-8位", 62, 8},

		// Base64 (包含+/=字符)
		{"Base64-4位", 64, 4},
		{"Base64-5位", 64, 5},
		{"Base64-6位", 64, 6},
		{"Base64-8位", 64, 8},
	}

	for _, scheme := range schemes {
		analyzeScheme(scheme.name, scheme.base, scheme.digits)
	}

	printComparisonTable(schemes)

	printLearningPoints()
}
