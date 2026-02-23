package main

import (
	"fmt"

	"strings"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"errors"
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

func toAbsoluteURLString(rawURL string, baseURL *url.URL) (string){
	parsedURL , err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil {
		fmt.Printf("ERROR: Can't parse URL [%s]\n",err)
		return ""
	}
	return baseURL.ResolveReference(parsedURL).String()

}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println(err)
		return nil,errors.New("Bad html body")
	}
	aHrefs := doc.Find("a[href]")
	l := aHrefs.Length()
	if l <= 0 {
		fmt.Println("ERROR: no links found!")
		return nil,nil
	}
	res := []string{}
	aHrefs.Each(func(_ int, s *goquery.Selection) {
        // For each '<a href>' it finds, it will run this function.
        href, ok := s.Attr("href")
        if ok {
        	c := toAbsoluteURLString(href,baseURL)
        	if c != "" {
        		res = append(res,c)	
        	}
        }
    })
	if len(res) == 0 {
		return nil,nil	
	}
	return res,nil
	
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println(err)
		return []string{},errors.New("Bad html body")
	}
	imgs := doc.Find("img[src]")
	l := imgs.Length()
	if l <= 0 {
		fmt.Println("ERROR: no images found!")
		return []string{},errors.New("no images found")
	}
	res := []string{}
	imgs.Each(func(_ int, s *goquery.Selection) {
        // For each '<a href>' it finds, it will run this function.
        img, ok := s.Attr("src")
        if ok {
        	c := toAbsoluteURLString(img,baseURL)
        	if c != "" {
        		res = append(res,c)	
        	}
        }
    })
	if len(res) == 0 {
		return nil,nil	
	}
	return res,nil
}
