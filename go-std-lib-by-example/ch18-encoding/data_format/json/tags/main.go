package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// User 展示各種 struct tag 的使用方式
type User struct {
	// json tag 展示
	ID        int64     `json:"id,omitempty"`        // 省略空值
	Username  string    `json:"username"`            // 基本重命名
	Password  string    `json:"-"`                   // 完全忽略此欄位
	LastLogin time.Time `json:"last_login,string"`   // 轉換為字串輸出
	Temp      *string   `json:"temporary,omitempty"` // 指針類型的省略

	// form tag 展示
	Email string `form:"email" binding:"required,email"` // 表單驗證
	Age   int    `form:"age" binding:"gte=0,lte=130"`    // 數值範圍驗證

	// validate tag 展示
	Phone   string `validate:"required,e164"` // 電話格式驗證
	Website string `validate:"url"`           // URL 格式驗證

	// gorm tag 展示
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"` // 資料庫欄位定義
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"` // 自動更新時間
	DeletedAt time.Time `gorm:"column:deleted_at;index"`          // 軟刪除索引
}

// Demo1 展示 json tag 的基本使用
func Demo1() {
	user := User{
		ID:       1,
		Username: "demo_user",
		Password: "secret",
	}

	jsonBytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("JSON Output:\n%s\n", string(jsonBytes))
}

// Demo2 展示 omitempty 的效果
func Demo2() {
	// 展示空值被省略的情況
	emptyUser := User{
		Username: "", // 空字串不會被省略
		ID:       0,  // 零值會被省略
	}

	jsonBytes, _ := json.MarshalIndent(emptyUser, "", "  ")
	fmt.Printf("Empty User JSON:\n%s\n", string(jsonBytes))
}

// Demo3 展示時間格式化
func Demo3() {
	user := User{
		Username:  "time_user",
		LastLogin: time.Now(),
	}

	jsonBytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("Time Format JSON:\n%s\n", string(jsonBytes))
}

func main() {
	fmt.Println("=== Demo1: Basic JSON Tags ===")
	Demo1()

	fmt.Println("\n=== Demo2: Omitempty Effect ===")
	Demo2()

	fmt.Println("\n=== Demo3: Time Format ===")
	Demo3()
}
