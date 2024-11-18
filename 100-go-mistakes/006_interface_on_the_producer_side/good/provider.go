package main

// Provider 只提供小而專注的介面
type Querier interface {
	Query(sql string) ([]string, error)
}

// Provider 的實作
type BetterReader struct{}

func (b BetterReader) Query(sql string) ([]string, error) {
	return []string{"data1", "data2"}, nil
}

// 用於測試的 Mock 實作
type MockQuerier struct{}

func (m MockQuerier) Query(sql string) ([]string, error) {
	return []string{"mock1", "mock2"}, nil
}
