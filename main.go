package main

import (
	"fmt"

	arabic "github.com/abdullahdiaa/garabic"
)

func main() {

    normalized := arabic.ToEnglishDigits("٢٠٢٣-٠٢-٠٦T١٨:٥٦:٠٨+00:00")
	fmt.Println(normalized)
}
