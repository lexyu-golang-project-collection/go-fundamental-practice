package main

import "fmt"

// Client 被迫接受整個大介面
type Client struct {
	db Database
}

func NewClient(db Database) *Client {
	return &Client{db: db}
}

func (c *Client) DoWork() {
	// 雖然只需要查詢功能
	data, _ := c.db.Query("SELECT * FROM table")
	fmt.Println("Query result:", data)

	// 但可以存取所有方法
	c.db.Insert("test")
	c.db.Update("test")
	c.db.Delete("1")
}

func main() {
	client := NewClient(SimpleReader{})
	client.DoWork()
}
