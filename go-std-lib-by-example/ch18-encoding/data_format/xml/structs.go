package main

import "encoding/xml"

// BaseEntity 提供基本的 XML 實體屬性
type BaseEntity struct {
	XMLName xml.Name `xml:"config"`
	Version string   `xml:"version,attr"`
}

// DatabaseConfig 資料庫配置
type DatabaseConfig struct {
	Host     string `xml:"host"`
	Port     int    `xml:"port"`
	Username string `xml:"username,omitempty"`
	Password string `xml:"-"`
}

// Server 服務器配置
type Server struct {
	Name   string `xml:"name,attr"`
	IP     string `xml:"ip"`
	Status string `xml:"status,attr,omitempty"`
}

// Property Key-Value 屬性
type Property struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",chardata"`
}

// Properties 屬性集合
type Properties struct {
	Items []Property `xml:"property"`
}

// Settings 設定
type Settings struct {
	Theme string `xml:"theme"`
	Mode  string `xml:"mode"`
}

// Metrics 指標數據
type Metrics struct {
	Count   int     `xml:"count"`
	Average float64 `xml:"average"`
	Active  bool    `xml:"active"`
}

// OptionalData 可選數據
type OptionalData struct {
	Value string `xml:"value"`
}

// Config 主配置結構
type Config struct {
	BaseEntity
	LastUpdate   string         `xml:"lastUpdate,attr"`
	Database     DatabaseConfig `xml:"database"`
	Servers      []Server       `xml:"server"`
	Description  string         `xml:"description"`
	Settings     Settings       `xml:"settings"`
	Flags        []string       `xml:"flags>flag"`
	Properties   Properties     `xml:"properties"`
	Metrics      Metrics        `xml:"metrics"`
	CustomField  string         `xml:"custom-field-name"`
	OptionalData *OptionalData  `xml:"optional,omitempty"`
}

// NamespacedEntity 命名空間基本實體
type NamespacedEntity struct {
	XMLName xml.Name `xml:"ns:root"`
	Xmlns   string   `xml:"xmlns:ns,attr"`
}

// NamespacedSettings 帶命名空間的設定
type NamespacedSettings struct {
	XMLName xml.Name `xml:"ns:settings"`
	Theme   string   `xml:"ns:theme"`
	Mode    string   `xml:"ns:mode"`
}

// NamespacedConfig 帶命名空間的配置
type NamespacedConfig struct {
	NamespacedEntity
	Settings NamespacedSettings `xml:"ns:settings"`
}

// DeepConfig 深層配置結構
type DeepConfig struct {
	Deeper struct {
		Deepest string `xml:"deepest"`
	} `xml:"deeper"`
}

// ExtraConfig 額外配置結構
type ExtraConfig struct {
	XMLName xml.Name   `xml:"root"`
	Value   string     `xml:",chardata"`
	Deep    DeepConfig `xml:"deep"`
	Special string     `xml:",innerxml"`
}
