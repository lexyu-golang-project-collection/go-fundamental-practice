package main

import "fmt"

func main() {
	var solutions [][]int
	pos := []int{1, 2, 3, 4}

	// 直接添加 pos
	solutions = append(solutions, pos)

	// 修改 pos
	pos[0] = 9

	fmt.Println("After modifying pos:")
	fmt.Println("solutions:", solutions)
	fmt.Println("pos:", pos)

	// 重置並使用複製方法
	solutions = [][]int{}
	pos = []int{1, 2, 3, 4}

	// 創建副本並添加
	solution := make([]int, len(pos))
	copy(solution, pos)
	solutions = append(solutions, solution)

	// 再次修改 pos
	pos[0] = 9

	fmt.Println("\nAfter modifying pos (with copy):")
	fmt.Println("solutions:", solutions)
	fmt.Println("pos:", pos)
}

// 輸出：
// After modifying pos:
// solutions: [[9 2 3 4]]
// pos: [9 2 3 4]
//
// After modifying pos (with copy):
// solutions: [[1 2 3 4]]
// pos: [9 2 3 4]
