package main
import (
	"net/url"
	"fmt"

	"strings"
	"github.com/PuerkitoBio/goquery"
)
func normalizeURL(rawURL string) (string, error) {
	// url.Parse(rawURL string) (*URL, error)
	parsedURL , err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("ERROR: Can't parse URL [%s]\n",err)
		return "", err
	}

	for ;; {
		if parsedURL.Path[len(parsedURL.Path)-1] != '/' {
			break
		} 
		parsedURL.Path = parsedURL.Path[:len(parsedURL.Path)-1]

	}

	return parsedURL.Host+ parsedURL.Path, nil
}

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	header := doc.Find("h1").Text()
	fmt.Println(header)
	return header
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	paragraph := doc.Find("p").First().Text()
	return paragraph
}
func getFirstMainParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	main := doc.Find("main").First()
	paragraph := main.Find("p").First().Text()
	fmt.Println(paragraph)
	return paragraph
}
