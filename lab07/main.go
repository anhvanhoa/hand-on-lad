package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
	Viết ứng dụng command file tên search , nhập vào đường dẫn thư mục và một từ khóa. Thư
	mục này chứa một số file *.txt, hãy tìm từ khóa trong tất cả file và in ra màn hình kết quả gồm
	tên file, dòng thứ mấy, nội dung dòng chứa từ khóa
*/

func searchInFile(filePath string, keyword string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Printf("%s\n%d: %s\n", filepath.Base(filePath), lineNumber, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	if input[0] != "search" {
		fmt.Println("Invalid command")
		return
	}
	folderName := input[1]
	searchString := input[2:]
	files, err := os.ReadDir(folderName)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			filePath := filepath.Join(folderName, file.Name())
			searchInFile(filePath, strings.Join(searchString, ""))
		}
	}
}
