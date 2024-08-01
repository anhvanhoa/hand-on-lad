package main

import (
	"fmt"
	"os"
)

func main() {
	var paragraph string = "Xin chao techmaster"
	var nameFile string = "paragraph.txt"
	// Tạo file .txt
	os.Create(nameFile)
	// Ghi vào file .txt
	os.WriteFile(nameFile, []byte(paragraph), 0644)
	// Đọc file .txt
	file, error := os.ReadFile(nameFile)
	// Xử lý lỗi
	if error != nil {
		panic(error)
	}
	// In
	fmt.Println(string(file))

}
