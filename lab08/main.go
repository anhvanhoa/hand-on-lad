package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(dir string, prefix string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for i, file := range files {
		// Kiêm tra xem file cuối cùng không
		isLastItem := i == len(files)-1
		if isLastItem {
			fmt.Print(prefix + "└── ")
		} else {
			fmt.Print(prefix + "├── ")
		}
		if file.IsDir() {
			fmt.Println(file.Name() + "/")
			// In cấu trúc cây của thư mục con
			newPrefix := "│   "
			printTree(filepath.Join(dir, file.Name()), newPrefix)
		} else {
			fmt.Println(file.Name())
		}
	}
}

/*
	Trong linux có một ứng dụng tree hiển thị cấu trúc cây thư mục trực quan trong màn hình
	console, hãy viết ứng dụng command line có chức năng giống với ứng dụng tree với một tham
	số là đường dẫn thư mục cần quét và hiển thị cấu trúc cây thư mục.
*/

func main() {
	// Nhập đường dẫn thư mục từ người dùng
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	if input[0] != "tree" {
		fmt.Println("Invalid command")
		return
	}
	folderName := input[1]
	// Kiểm tra sự tồn tại của thư mục
	info, err := os.Stat(folderName)
	if err != nil || !info.IsDir() {
		fmt.Println("The specified path is not a valid directory.")
		return
	}

	// In cấu trúc cây của thư mục
	printTree(folderName, "")
}
