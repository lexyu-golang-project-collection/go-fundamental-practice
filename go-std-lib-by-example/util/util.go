package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetProjectRoot() string {
	dir, err := executableDir()
	if err != nil {
		return ""
	}
	fmt.Println("dir = ", dir)
	return path.Dir(dir)
}

func executableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	fmt.Println("pathAbs = ", pathAbs)
	return filepath.Dir(pathAbs), nil
}

func UTF8Index(str, sub string) int {
	asciiPos := strings.Index(str, sub)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}

	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
