package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func DecodeDemo1() {
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<config version="1.0">
    <database>
        <host>localhost</host>
        <port>5432</port>
    </database>
    <description><![CDATA[This is a CDATA example]]></description>
    <properties>
        <property key="timeout">30</property>
        <property key="retries">3</property>
    </properties>
</config>`

	var config Config
	if err := xml.Unmarshal([]byte(xmlData), &config); err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
		return
	}

	fmt.Printf("解碼結果:\n")
	fmt.Printf("Version: %s\n", config.Version)
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
	fmt.Printf("Description: %s\n", config.Description)

	fmt.Println("Properties:")
	for _, prop := range config.Properties.Items {
		fmt.Printf("  %s: %s\n", prop.Key, prop.Value)
	}
}

func DecodeDemo2() {
	xmlStream := strings.NewReader(`<?xml version="1.0"?>
    <config>
        <server name="server1" status="active">
            <ip>192.168.1.1</ip>
        </server>
        <server name="server2">
            <ip>192.168.1.2</ip>
        </server>
    </config>`)

	decoder := xml.NewDecoder(xmlStream)

	var server Server

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("解碼錯誤: %v\n", err)
			return
		}

		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "server":
				server = Server{} // 重置 server
				for _, attr := range t.Attr {
					switch attr.Name.Local {
					case "name":
						server.Name = attr.Value
					case "status":
						server.Status = attr.Value
					}
				}
			case "ip":
				if ipToken, err := decoder.Token(); err == nil {
					if text, ok := ipToken.(xml.CharData); ok {
						server.IP = string(text)
					}
				}
			}
		case xml.EndElement:
			if t.Name.Local == "server" {
				fmt.Printf("找到伺服器: %s (狀態: %s, IP: %s)\n",
					server.Name, server.Status, server.IP)
			}
		}
	}
}

func DecodeDemo2Simple1() {
	xmlStream := strings.NewReader(`<?xml version="1.0"?>
    <config>
        <server name="server1" status="active">
            <ip>192.168.1.1</ip>
        </server>
        <server name="server2">
            <ip>192.168.1.2</ip>
        </server>
    </config>`)

	// 定義用於解析的結構
	type ServerConfig struct {
		Servers []Server `xml:"server"`
	}

	decoder := xml.NewDecoder(xmlStream)
	var config ServerConfig

	// 直接解碼整個文檔
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("解碼錯誤: %v\n", err)
			return
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "config" {
				if err := decoder.DecodeElement(&config, &se); err != nil {
					fmt.Printf("解碼錯誤: %v\n", err)
					return
				}
				// 解碼完成後輸出結果
				for _, server := range config.Servers {
					fmt.Printf("找到伺服器: %s (狀態: %s, IP: %s)\n",
						server.Name, server.Status, server.IP)
				}
			}
		}
	}
}

func DecodeDemo2Simple2() {
	xmlData := `<?xml version="1.0"?>
    <config>
        <server name="server1" status="active">
            <ip>192.168.1.1</ip>
        </server>
        <server name="server2">
            <ip>192.168.1.2</ip>
        </server>
    </config>`

	type ServerConfig struct {
		Servers []Server `xml:"server"`
	}

	var config ServerConfig
	if err := xml.Unmarshal([]byte(xmlData), &config); err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
		return
	}

	for _, server := range config.Servers {
		fmt.Printf("找到伺服器: %s (狀態: %s, IP: %s)\n",
			server.Name, server.Status, server.IP)
	}
}

func DecodeDemo3() {
	xmlData := `<?xml version="1.0"?>
    <root>
        直接內容
        <deep>
            <deeper>
                <deepest>最深層值</deepest>
            </deeper>
        </deep>
        <special><customXML>保留原始內容</customXML></special>
    </root>`

	var config ExtraConfig
	decoder := xml.NewDecoder(bytes.NewReader([]byte(xmlData)))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose

	if err := decoder.Decode(&config); err != nil {
		fmt.Printf("解碼錯誤: %v\n", err)
		return
	}

	fmt.Printf("解碼特殊內容:\n")
	fmt.Printf("Value: %s\n", config.Value)
	fmt.Printf("Deepest: %s\n", config.Deep.Deeper.Deepest)
	fmt.Printf("Special: %s\n", config.Special)
}
