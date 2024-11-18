package main

import "fmt"

// Client 只依賴需要的介面
type DataReader struct {
	querier Querier
}

func NewDataReader(q Querier) *DataReader {
	return &DataReader{querier: q}
}

func (d *DataReader) ReadData() {
	data, _ := d.querier.Query("SELECT * FROM table")
	fmt.Println("Query result:", data)
}

func main() {
	// 使用實際實作
	reader := NewDataReader(BetterReader{})
	reader.ReadData()

	// 使用 mock 測試
	testReader := NewDataReader(MockQuerier{})
	testReader.ReadData()
}
