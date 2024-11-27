package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Config struct {
	// 使用更常見的特殊字元模式
	Filter   string `json:"$filter"`   // MongoDB 風格的查詢
	Select   string `json:"$select"`   // GraphQL 風格的選擇器
	Metadata string `json:"@metadata"` // XML 風格的詮釋資料
}

// DynamicMessage 展示更實用的動態 JSON 處理場景
type DynamicMessage struct {
	Action string `json:"action"`
	// Payload 可以容納任何合法的 JSON 結構
	Payload json.RawMessage `json:"payload"`
}

func main() {
	// 特殊字元示範
	config := Config{
		Filter:   "status eq 'active'",
		Select:   "id, name, email",
		Metadata: "version=1.0",
	}

	configJSON, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=== 特殊字元示範 ===\n%s\n\n", string(configJSON))

	// 動態 JSON 處理示範
	userPayload := json.RawMessage(`{
        "user": {
            "name": "張小明",
            "age": 25,
            "preferences": {
                "theme": "dark",
                "language": "zh-TW"
            }
        }
    }`)

	message := DynamicMessage{
		Action:  "UPDATE_USER",
		Payload: userPayload,
	}

	messageJSON, err := json.MarshalIndent(message, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=== 動態 JSON 示範 ===\n%s\n", string(messageJSON))
}
