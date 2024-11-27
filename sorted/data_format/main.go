package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

// User 展示基本的 JSON tags 用法
type User struct {
	ID       int    `json:"id"`                    // 自定義欄位名稱
	Name     string `json:"name,omitempty"`        // 空值省略
	Password string `json:"-"`                     // 完全忽略此欄位
	Age      int    `json:"age,string"`            // 數字轉字串
	Address  string `json:"addr,omitempty,string"` // 組合多個選項
}

// Book 展示基本的 XML tags 用法
type Book struct {
	XMLName   xml.Name `xml:"book"`       // 指定 XML 元素名稱
	Title     string   `xml:"title"`      // 基本映射
	Author    string   `xml:"author"`     // 基本映射
	Price     float64  `xml:"price,attr"` // 作為屬性
	ISBN      string   `xml:",omitempty"` // 空值省略
	Publisher string   `xml:"-"`          // 忽略此欄位
}

// Library 展示進階的 XML 結構控制
type Library struct {
	XMLName  xml.Name `xml:"library"`    // 指定根元素名稱
	Name     string   `xml:"name,attr"`  // 作為屬性
	Location string   `xml:"location"`   // 純文字內容
	Books    []Book   `xml:"books>book"` // 巢狀元素
}

func main() {
	// JSON Demo
	fmt.Println("=== JSON Demo ===")
	user := User{
		ID:       1,
		Name:     "John Doe",
		Password: "secret123",
		Age:      30,
		Address:  "123 Main St",
	}

	// JSON 序列化
	jsonData, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON Output:\n%s\n\n", string(jsonData))

	// XML Demo
	fmt.Println("=== XML Demo ===")
	library := Library{
		Name:     "City Library",
		Location: "Downtown",
		Books: []Book{
			{
				Title:     "Go Programming",
				Author:    "John Smith",
				Price:     29.99,
				ISBN:      "123-456-789",
				Publisher: "Tech Books",
			},
			{
				Title:     "Cloud Native",
				Author:    "Jane Doe",
				Price:     39.99,
				ISBN:      "987-654-321",
				Publisher: "Cloud Pub",
			},
		},
	}

	// XML 序列化
	xmlData, err := xml.MarshalIndent(library, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("XML Output:\n%s\n", string(xmlData))

	// 示範反序列化
	fmt.Println("\n=== Deserialization Demo ===")

	// JSON 反序列化
	var newUser User
	if err := json.Unmarshal(jsonData, &newUser); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deserialized User: %+v\n", newUser)

	// XML 反序列化
	var newLibrary Library
	if err := xml.Unmarshal(xmlData, &newLibrary); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deserialized Library Name: %s\n", newLibrary.Name)
	fmt.Printf("First Book: %+v\n", newLibrary.Books[0])
}
