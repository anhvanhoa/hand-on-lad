package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Đọc file .txt
	var nameFile string = "paragraph.txt"
	file, error := os.ReadFile(nameFile)
	// Convert file sang string
	content := string(file)
	// Xử lý lỗi
	if error != nil {
		panic(error)
	}

	// Các từ nhạy cảm cần thay thế nguyên âm
	sensitiveWords := []string{"sex", "fuck", "drug", "kill"}
	// Hàm thay thế nguyên âm trong từ nhạy cảm
	replaceVowels := func(word string) string {
		vowels := "aeiouAEIOU"
		for _, vowel := range vowels {
			word = strings.ReplaceAll(word, string(vowel), "*")
		}
		return word
	}

	// Kiểm tra xem nội dung có từ nhạy cảm hay không
	for _, word := range sensitiveWords {
		if strings.Contains(content, word) {
			// Tạo phiên bản đã thay thế nguyên âm
			censoredWord := replaceVowels(word)
			content = strings.ReplaceAll(content, word, censoredWord)
		}
	}
	// In
	fmt.Println(string(file))
}
