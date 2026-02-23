package main

import (
	"net/url"
	"fmt"
	"strings"
)

type PageData struct {
	URL 			string
	H1				string
	FirstParagraph 	string	
	OutgoingLinks	[]string
	ImageURLs		[]string
}

func extractPageData(html, pageURL string) PageData {
	pageH1 := getH1FromHTML(html)
	pageFirstParagraph := getFirstParagraphFromHTML(html)

	parsedURL , err := url.Parse(strings.TrimSpace(pageURL))
	if err != nil {
		fmt.Printf("ERROR: Can't parse pageURL [%s]\n",err)
		return PageData{
			URL:            pageURL,
			H1:             pageH1,
			FirstParagraph: pageFirstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}
	
	pageOutgoingLinks,err := getURLsFromHTML(html,parsedURL)	
	if err != nil {
		fmt.Printf("ERROR: Can't get links [%s]\n",err)
		pageOutgoingLinks = nil
	}
	pageImageURLs,err := getImagesFromHTML(html,parsedURL)
	if err != nil {
		fmt.Printf("ERROR: Can't get images [%s]\n",err)
		pageImageURLs = nil
	}
	return PageData{
		URL: parsedURL.String(),
		H1: pageH1,
		FirstParagraph: pageFirstParagraph,
		OutgoingLinks: pageOutgoingLinks,
		ImageURLs: pageImageURLs,
	}
}