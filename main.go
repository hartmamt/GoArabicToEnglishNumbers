package main

import (
	"fmt"

	arabic "github.com/abdullahdiaa/garabic"
)

func main() {

    // Make slice of Arabic dates
    dates := []string{
        "٢٠٢٣-٠٢-٠٧T١٢:٤٣:٢٦+00:00",
        "٢٠٢٣-٠٢-٠٧T١٣:١٣:١٩+00:00",
        "٢٠٢٣-٠٢-٠٧T١٣:١٨:٣٩+00:00",
        "٢٠٢٣-٠٢-٠٧T١٣:١٨:٣٩+00:00",
        "٢٠٢٣-٠٢-٠٧T١٣:٢٥:٣٧+00:00",
        "٢٠٢٣-٠٢-٠٦T١٩:٢٧:٤٧+00:00",
        "٢٠٢٣-٠٢-٠٦T٠٦:٥١:٢٠+00:00",
        "٢٠٢٣-٠٢-٠٦T١٢:٥٨:٤٣+00:00",
        "٢٠٢٣-٠٢-٠٦T١٨:٥٦:٠٨+00:00",
        "٢٠٢٣-٠٢-٠٧T١٢:٢٨:٢٧+00:00",
    }
   
    // conver dates to English
    for _, date := range dates {
        normalized := arabic.ToEnglishDigits(date)
        fmt.Println(normalized)
    }
  
    // prints
    /*
    2023-02-07T12:43:26+00:00
    2023-02-07T13:13:19+00:00
    2023-02-07T13:18:39+00:00
    2023-02-07T13:18:39+00:00
    2023-02-07T13:25:37+00:00
    2023-02-06T19:27:47+00:00
    2023-02-06T06:51:20+00:00
    2023-02-06T12:58:43+00:00
    2023-02-06T18:56:08+00:00
    2023-02-07T12:28:27+00:00
    */
}
