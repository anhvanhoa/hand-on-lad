package main

import (
	"fmt"
	"os"
)

/*
	Hãy tạo ứng dụng command line, nhập vào tên file và một số nguyên dương X > 10 , hãy tạo ra
	một file gồm những dòng dài 256 ký tự ngẫu nhiên A-Za-z0-9 làm sao để kích thước của file là X
	megabytes (mega tính bằng 1024 * 1024 = 1.048.576 bytes). Điều gì xảy ra khi X cực kỳ lớn, bạn
	sẽ tạo file kích thước khủng. Hãy quan sát đến tới hạn nào thì ứng dụng báo lỗi, hãy giải thích
	nguyên nhân và cách khắc phục.
*/

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
