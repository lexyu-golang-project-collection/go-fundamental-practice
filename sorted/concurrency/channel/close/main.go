package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// ChannelCloseDemo1()
	// ChannelCloseDemo2()
	// ChannelCloseDemo3()
	// ChannelCloseWithWaitGroup()
	ChannelCloseWithTwoWayChannel()
	// ChannelCloseWithContext()
}

func ChannelCloseDemo1() {
	ch := make(chan struct{})

	// 情況 1: 直接讀取
	go func() {
		<-ch // 會等待直到 channel 關閉
		fmt.Println("Receiver 1: detected channel close")
	}()
	close(ch)
	time.Sleep(2 * time.Second)
}

func ChannelCloseDemo2() {
	ch := make(chan struct{})

	// 情況 2: 使用 for-range
	go func() {
		for range ch { // 會等待直到 channel 關閉
			// 不會執行到這裡，因為 channel 類型是 struct{}
		}
		fmt.Println("Receiver 2: detected channel close")
	}()

	close(ch)
	time.Sleep(2 * time.Second)
}

func ChannelCloseDemo3() {
	ch := make(chan struct{})

	// 情況 3: 使用 select
	go func() {
		select {
		case <-ch: // 會在 channel 關閉時被觸發
			fmt.Println("Receiver 3: detected channel close")
		}
	}()

	close(ch)
	time.Sleep(2 * time.Second)
}

func ChannelCloseWithWaitGroup() {
	fmt.Println("\n=== Using WaitGroup ===")
	var wg sync.WaitGroup
	ch := make(chan struct{})

	fmt.Println("1. Starting goroutine...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2. Goroutine waiting for channel...")
		<-ch
		fmt.Println("4. Goroutine received channel close signal")
		time.Sleep(3 * time.Second) // 模擬一些工作
		fmt.Println("5. Goroutine finished work")
	}()

	time.Sleep(time.Second) // 讓 goroutine 有時間開始等待
	fmt.Println("3. Main: closing channel...")
	close(ch)

	fmt.Println("6. Main: waiting for goroutine to complete...")
	wg.Wait()
	fmt.Println("7. Main: all done!")
}

func ChannelCloseWithTwoWayChannel() {
	fmt.Println("\n=== Using Two-Way Channel ===")
	ch := make(chan struct{})
	done := make(chan struct{})

	fmt.Println("1. Starting goroutine...")
	go func() {
		fmt.Println("2. Goroutine waiting for channel...")
		<-ch
		fmt.Println("4. Goroutine received channel close signal")
		time.Sleep(time.Second) // 模擬一些工作
		fmt.Println("5. Goroutine signaling completion...")
		close(done)
	}()

	time.Sleep(time.Second) // 讓 goroutine 有時間開始等待
	fmt.Println("3. Main: closing channel...")
	close(ch)

	fmt.Println("6. Main: waiting for completion signal...")
	<-done
	fmt.Println("7. Main: received completion signal!")
}

func ChannelCloseWithContext() {
	fmt.Println("\n=== Using Context ===")
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})

	fmt.Println("1. Starting goroutine...")
	go func() {
		fmt.Println("2. Goroutine waiting for channel...")
		<-ch
		fmt.Println("4. Goroutine received channel close signal")
		time.Sleep(time.Second) // 模擬一些工作
		fmt.Println("5. Goroutine canceling context...")
		cancel()
	}()

	time.Sleep(time.Second) // 讓 goroutine 有時間開始等待
	fmt.Println("3. Main: closing channel...")
	close(ch)

	fmt.Println("6. Main: waiting for context cancellation...")
	<-ctx.Done()
	fmt.Println("7. Main: context cancelled!")
}
