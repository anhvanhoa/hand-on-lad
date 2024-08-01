package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Viết ứng dụng command line, nhập vào tên file, và số nguyên dương Y. Đọc file file ở dòng Y rồi
	in ra. Chú ý ngoại lệ khi Y > tổng số dòng của file và khi file có kích thước cực kỳ lớn.
*/

func main() {
	var fileName string
	var lineNumber int

	// Nhập tên file và số dòng cần đọc
	fmt.Print("Enter the file name: ")
	fmt.Scan(&fileName)

	fmt.Print("Enter the line number to read: ")
	fmt.Scan(&lineNumber)

	if lineNumber <= 0 {
		fmt.Println("Line number must be greater than 0.")
		return
	}

	// Mở file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Đọc dòng từ file
	scanner := bufio.NewScanner(file)
	currentLine := 0

	for scanner.Scan() {
		currentLine++
		if currentLine == lineNumber {
			fmt.Println("Line", lineNumber, ":", scanner.Text())
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if currentLine < lineNumber {
		fmt.Printf("File has only %d lines. Line %d does not exist.\n", currentLine, lineNumber)
	}
}
