package gen

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func SubmitDay(year int, day int, answer string, level string) {
	client := NewClient("https://adventofcode.com")
	body, err := client.SubmitAnswer(day, year, level, answer)
	if err != nil {
		log.Fatalf("Error with HTTP request: %s", err.Error())
	}

	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Find("article").Text())
}
