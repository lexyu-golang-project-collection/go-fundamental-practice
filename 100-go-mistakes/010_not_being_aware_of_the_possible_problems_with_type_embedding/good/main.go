package main

import (
	"io"
	"os"
	"sync"
)

// 1. 替代 Foo/Bar 的範例 - 使用明確的組合
type Customer struct {
	profile Profile
}

type Profile struct {
	ID   int
	Name string
}

func (c *Customer) GetProfileName() string {
	return c.profile.Name
}

// 2. 替代 InMem 的範例 - 隱藏同步細節
type Storage struct {
	mu sync.Mutex
	m  map[string]int
}

func NewStorage() *Storage {
	return &Storage{
		m: make(map[string]int),
	}
}

func (s *Storage) Get(key string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, exists := s.m[key]
	return v, exists
}

func (s *Storage) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

// 3. 替代 Logger 的範例 - 明確的組合和介面實作
type Logger struct {
	writer io.WriteCloser
}

func NewLogger(w io.WriteCloser) *Logger {
	return &Logger{
		writer: w,
	}
}

func (l *Logger) Write(p []byte) (int, error) {
	return l.writer.Write(p)
}

func (l *Logger) Close() error {
	return l.writer.Close()
}

// 使用範例
func main() {
	// Customer 使用範例
	customer := &Customer{
		profile: Profile{
			ID:   1,
			Name: "John",
		},
	}
	_ = customer.GetProfileName()

	// Storage 使用範例
	storage := NewStorage()
	storage.Set("key1", 100)
	value, exists := storage.Get("key1")
	_ = value
	_ = exists

	// Logger 使用範例
	logger := NewLogger(os.Stdout)
	_, _ = logger.Write([]byte("log message\n"))
	_ = logger.Close()
}
