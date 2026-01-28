package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func EncodeDemo1() {
	config := Config{
		BaseEntity: BaseEntity{
			Version: "1.0",
		},
		LastUpdate: "2024-01-01",
		Database: DatabaseConfig{
			Host: "localhost",
			Port: 5432,
		},
		Description: "This is a description with <special> characters",
		Properties: Properties{
			Items: []Property{
				{Key: "timeout", Value: "30"},
				{Key: "retries", Value: "3"},
			},
		},
		Settings: Settings{
			Theme: "dark",
			Mode:  "prod",
		},
		Servers: []Server{
			{Name: "main", IP: "192.168.1.1", Status: "active"},
			{Name: "backup", IP: "192.168.1.2"},
		},
	}

	output, err := xml.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("編碼錯誤: %v\n", err)
		return
	}

	xmlData := []byte(xml.Header + string(output))
	fmt.Printf("基本 XML 輸出:\n%s\n", string(xmlData))
}

func EncodeDemo2() {
	nsConfig := NamespacedConfig{
		NamespacedEntity: NamespacedEntity{
			Xmlns: "http://example.com/ns",
		},
		Settings: NamespacedSettings{
			Theme: "dark",
			Mode:  "production",
		},
	}

	fmt.Println("使用 Encoder 輸出帶有命名空間的 XML:")
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	fmt.Print(xml.Header)

	if err := enc.Encode(nsConfig); err != nil {
		fmt.Printf("編碼錯誤: %v\n", err)
	}
}

func EncodeDemo3() {
	config := ExtraConfig{
		Value: "直接內容",
		Deep: DeepConfig{
			Deeper: struct {
				Deepest string `xml:"deepest"`
			}{
				Deepest: "最深層值",
			},
		},
		Special: "<customXML>保留原始內容</customXML>",
	}

	output, err := xml.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("編碼錯誤: %v\n", err)
		return
	}

	fmt.Printf("特殊內容 XML:\n%s\n", string(output))
}
