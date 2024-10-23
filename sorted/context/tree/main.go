package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type contextKey string

const (
	// 為不同層級定義不同的 key
	ctx1Key = contextKey("ctx1_value")
	ctx2Key = contextKey("ctx2_value")
	ctx3Key = contextKey("ctx3_value")
	ctx4Key = contextKey("ctx4_value")
)

func main() {
	var wg sync.WaitGroup

	// Context 層級
	rootCtx := context.Background()

	// Ctx 1
	ctx1, cancel1 := context.WithCancel(rootCtx)
	ctx1 = context.WithValue(ctx1, ctx1Key, "ctx1-value")

	// Ctx 2 (G2, G3 共享)
	ctx2 := context.WithValue(ctx1, ctx2Key, "ctx2-value")

	// Ctx 3 (G4, G5 共享)
	ctx3 := context.WithValue(ctx2, ctx3Key, "ctx3-value")

	// Ctx 4 (G6 使用)
	ctx4 := context.WithValue(ctx2, ctx4Key, "ctx4-value")

	wg.Add(6)

	// G1 使用 ctx1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("G1 cancelled")
				return
			default:
				fmt.Printf("G1 lookup: ctx1=%v\n",
					ctx1.Value(ctx1Key))
				time.Sleep(time.Second)
			}
		}
	}()

	// G2 使用 ctx2
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("G2 cancelled")
				return
			default:
				fmt.Printf("G2 lookup: ctx1=%v, ctx2=%v\n",
					ctx2.Value(ctx1Key),
					ctx2.Value(ctx2Key))
				time.Sleep(time.Second)
			}
		}
	}()

	// G3 使用 ctx2
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("G3 cancelled")
				return
			default:
				fmt.Printf("G3 lookup: ctx1=%v, ctx2=%v\n",
					ctx2.Value(ctx1Key),
					ctx2.Value(ctx2Key))
				time.Sleep(time.Second)
			}
		}
	}()

	// G4 使用 ctx3
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx3.Done():
				fmt.Println("G4 cancelled")
				return
			default:
				fmt.Printf("G4 lookup: ctx1=%v, ctx2=%v, ctx3=%v\n",
					ctx3.Value(ctx1Key),
					ctx3.Value(ctx2Key),
					ctx3.Value(ctx3Key))
				time.Sleep(time.Second)
			}
		}
	}()

	// G5 使用 ctx3
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx3.Done():
				fmt.Println("G5 cancelled")
				return
			default:
				fmt.Printf("G5 lookup: ctx1=%v, ctx2=%v, ctx3=%v\n",
					ctx3.Value(ctx1Key),
					ctx3.Value(ctx2Key),
					ctx3.Value(ctx3Key))
				time.Sleep(time.Second)
			}
		}
	}()

	// G6 使用 ctx4
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx4.Done():
				fmt.Println("G6 cancelled")
				return
			default:
				fmt.Printf("G6 lookup: ctx1=%v, ctx2=%v, ctx4=%v\n",
					ctx4.Value(ctx1Key),
					ctx4.Value(ctx2Key),
					ctx4.Value(ctx4Key))
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("\nCancelling from ctx1...")
	cancel1()

	wg.Wait()
	fmt.Println("All goroutines finished")
}
