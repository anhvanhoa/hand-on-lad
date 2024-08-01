package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName string
	var sizeMB int
	// Nhập tên file và kích thước file từ người dùng
	fmt.Print("Enter the file name: ")
	fmt.Scan(&fileName)

	fmt.Print("Enter the size in MB (greater than 10): ")
	fmt.Scan(&sizeMB)

	if sizeMB <= 10 {
		fmt.Println("Size must be greater than 10 MB.")
		return
	}

	// Kích thước cần tạo
	const LineLength = 256
	var totalSizeBytes int = sizeMB * 1024 * 1024
	fmt.Println("Total size in bytes:", totalSizeBytes)
	// Tạo hoặc mở file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Ghi nội dung vào file
	line := make([]byte, LineLength)
	for i := 0; i < LineLength; i++ {
		line[i] = 'A' + byte(i%26)
	}

	for totalSizeBytes > 0 {
		chunkSize := LineLength
		if chunkSize > totalSizeBytes {
			chunkSize = totalSizeBytes
		}
		_, err := file.Write(line[:chunkSize])
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		totalSizeBytes -= chunkSize
	}
	fmt.Println("File created successfully!")
}
