package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	fmt.Println("xxx")
	url := "https://ww2.jujustukaisen.com/"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the website:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// fmt.Println(string(body))

	text := string(body)
	re := regexp.MustCompile(`ww2.jujustukaisen.com/manga/jujutsu-kaisen-chapter-\s*(\d+)`)

	matches := re.FindAllString(text, -1)

	// if match != "" {
	// 	fmt.Println("Extracted string:", match)
	// } else {
	// 	fmt.Println("No match found.")
	// }

	// Extract and print the chapter numbers
	for _, match := range matches {
		chapterNumber := match
		fmt.Println("Chapter number:", chapterNumber)
	}

}
