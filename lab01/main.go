package main

import (
	"fmt"
	"os"
)

func main() {
	// Đọc file .txt
	nameFile := "paragraph.txt"
	file, error := os.ReadFile(nameFile)
	// Xử lý lỗi
	if error != nil {
		panic(error)
	}
	// In
	fmt.Println(string(file))
}
