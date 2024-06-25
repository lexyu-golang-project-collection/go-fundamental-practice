package main

import "fmt"

// 定義 Room 結構體
type Room struct {
	name string
}

// 初始化 Living Room 的方法
func NewLivingRoom() *Room {
	fmt.Println("Living Room Initialization Start")
	room := &Room{name: "Living Room"}
	fmt.Println("Living Room Initialization End")
	return room
}

// 初始化 Bedroom 的方法
func NewBedroom() *Room {
	fmt.Println("Bedroom Initialization Start")
	room := &Room{name: "Bedroom"}
	fmt.Println("Bedroom Initialization End")
	return room
}

// 定義 House 結構體
type House struct {
	rooms []*Room
}

// 定義 House 結構體的初始化方法
func NewHouse() *House {
	fmt.Println("House Initialization Start")
	house := &House{
		rooms: []*Room{
			NewLivingRoom(), // 初始化客廳
			NewBedroom(),    // 初始化臥室
		},
	}
	fmt.Println("House Initialization End")
	return house
}

// 定義 House 結構體的方法展示房間
func (h *House) ShowRooms() {
	fmt.Println("House has the following rooms:")
	for _, room := range h.rooms {
		fmt.Println(" - " + room.name)
	}
}

func main() {
	// 創建一個 House 實例
	house := NewHouse()
	// 調用 House 實例的 ShowRooms 方法
	house.ShowRooms()
}
