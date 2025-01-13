package main

import (
	"fmt"
	"time"
)

// --- Bad Example: Returning Interface ---
type DataStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type memoryStore struct {
	data map[string]string
}

// Bad: Returns interface instead of concrete type
func NewDataStore() DataStore {
	return &memoryStore{
		data: make(map[string]string),
	}
}

func (m *memoryStore) Get(key string) (string, error) {
	if v, ok := m.data[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("key not found")
}

func (m *memoryStore) Set(key, value string) error {
	m.data[key] = value
	return nil
}

// --- Good Example: Return Concrete Type, Accept Interface ---
type Reader interface {
	Get(key string) (string, error)
}

type Writer interface {
	Set(key, value string) error
}

type Store struct {
	data map[string]string
}

// Good: Returns concrete type
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, error) {
	if v, ok := s.data[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("key not found")
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return fmt.Errorf("empty key not allowed")
	}
	s.data[key] = value
	return nil
}

// Good: Accepts interfaces as parameters
func CopyData(r Reader, w Writer, key string) error {
	value, err := r.Get(key)
	if err != nil {
		return fmt.Errorf("read error: %v", err)
	}
	return w.Set(key, value)
}

func ProcessDataAsync(r Reader, writers []Writer, key string) <-chan error {
	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		value, err := r.Get(key)
		if err != nil {
			errCh <- fmt.Errorf("read error: %v", err)
			return
		}

		for _, w := range writers {
			if err := w.Set(key, value); err != nil {
				errCh <- fmt.Errorf("write error: %v", err)
				return
			}
		}
	}()

	return errCh
}

func main() {
	// Bad usage example
	badStore := NewDataStore()
	badStore.Set("test", "data")

	// Good usage example
	store1 := NewStore()
	store2 := NewStore()

	// 準備資料
	store1.Set("key1", "value1")

	// 同步複製
	fmt.Println("Synchronous copy:")
	err := CopyData(store1, store2, "key1")
	if err != nil {
		fmt.Printf("Sync copy failed: %v\n", err)
	} else {
		v, _ := store2.Get("key1")
		fmt.Printf("Copied value: %s\n", v)
	}

	// 非同步複製
	fmt.Println("\nAsynchronous copy:")
	writers := []Writer{store2}
	errCh := ProcessDataAsync(store1, writers, "key1")

	// 等待非同步操作完成
	select {
	case err := <-errCh:
		if err != nil {
			fmt.Printf("Async copy failed: %v\n", err)
		} else {
			v, _ := store2.Get("key1")
			fmt.Printf("Async copied value: %s\n", v)
		}
	case <-time.After(time.Second):
		fmt.Println("Timeout waiting for async copy")
	}
}
