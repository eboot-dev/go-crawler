package main

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	header := doc.Find("h1").Text()
	// fmt.Println(header)
	return header
}


func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// main := doc.Find("main")
	// var p string
	// if main.Length() > 0 {
	// 	p = (main.First()).Find("p").First().Text()
	// } else {
		p := doc.Find("p").First().Text()
	// }
	paragraph := strings.TrimSpace(p)
	// fmt.Println(paragraph)
	return paragraph
}


func getFirstMainParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	main := doc.Find("main")
	if main.Length() <= 0 {
		fmt.Println("ERROR: No main tag found!")
		return ""
	}
	p := (main.First()).Find("p").First().Text()
	paragraph := strings.TrimSpace(p)
	// fmt.Println(paragraph)
	return paragraph
}



