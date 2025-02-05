package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Đặt một số câu hỏi: Bạn tên là gì? Bạn sinh ngày nào? Bạn làm nghề gì ? Nhận câu trả lời sau
	đó ghi cả câu hỏi và trả lời vào file person.txt tách thành từng dòng cho dễ nhìn.
*/

func main() {
	var questions = [4]string{"Bạn tên là gì?", "Bạn bao nhiêu tuổi?", "Bạn đến từ đâu?", "Hôm nay bạn như thế nào?"}
	var nameFile string = "person.txt"
	// Đọc file .txt
	file, err := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE, 0644)
	// Xử lý lỗi
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	for _, question := range questions {
		fmt.Println(question)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		paragraph := scanner.Text()
		// Ghi vào file .txt
		file.WriteString(question + "\n")
		file.WriteString(paragraph + "\n")
	}

	// Đọc file .txt
	content, error := os.ReadFile(nameFile)
	// Xử lý lỗi
	if error != nil {
		panic(error)
	}
	// In
	fmt.Println("Nội dung file:")
	fmt.Println(string(content))
}
