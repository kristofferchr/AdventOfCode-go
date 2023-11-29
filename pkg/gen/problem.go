package gen

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

func GetProblems(year int, day int) ([]string, error) {
	client := NewClient("https://adventofcode.com")
	body, err := client.Get(fmt.Sprintf("%s/day/%s", strconv.Itoa(year), strconv.Itoa(day)))
	if err != nil {
		log.Fatalf("Error with HTTP request: %s", err.Error())
	}

	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal(err)
	}

	markdownContent := doc.Find(".day-desc").
		Map(func(i int, s *goquery.Selection) string {
			content := s.Clone()
			title := content.Find("h2").Remove()

			converter := md.NewConverter("", true, nil)
			result := converter.Convert(content)

			titleFormatted := strings.ReplaceAll(title.Text(), "-", "")
			return fmt.Sprintf("# %s \n\n%s ", titleFormatted, result)
		})

	return markdownContent, nil
}
func GetProblem(year int, day int) (string, error) {
	problems, err := GetProblems(year, day)
	if err != nil {
		return "", err
	}

	return strings.Join(problems, "\n"), nil
}
