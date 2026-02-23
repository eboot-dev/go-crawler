package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

//Recursive page crawling
func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int, bodies map[string]string) {
	baseURL , err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("[crawlPage] could not parse base URL: %s\n", err)
		os.Exit(1)
	}
	currentURL , err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("[crawlPage] could not parse current URL: %s\n", err)
		// os.Exit(1)
		return
	}
	
	if !strings.Contains(currentURL.Host, baseURL.Host) {
		// fmt.Printf("[crawlPage] IGNORE URL: Different Hosts [curret: %s] [base: %s].\n",currentURL.Host,baseURL.Host)
		return
	} 

	pURL := currentURL.String()
	currentURLNorm, err := normalizeURL(pURL)
	if err != nil {
		fmt.Printf("[crawlPage] could not normalize URL: %s\n", err)
		return
	}
	cnt,ok := pages[currentURLNorm]
	if ok {
		// Already explored
		// fmt.Println("[crawlPage] Explored already!")
		pages[currentURLNorm] = cnt + 1
		return
	}
	// To be explored
	fmt.Printf("[crawlPage] Exploring %s\n",pURL)
	body, err := getHTML(pURL)
	if err != nil {
		fmt.Printf("[crawlPage] ERROR: Unreachable URL [%s]\n",err)
		pages[currentURLNorm] = 1 // mark nod visited to avoi recursion
		bodies[currentURLNorm] = ""
		return
	}
	pages[currentURLNorm] = 1
	bodies[currentURLNorm] = body

	pageOutgoingLinks,err := getURLsFromHTML(body,currentURL)	
	if err != nil {
		fmt.Printf("[crawlPage] ERROR: Can't get links [%s]\n",err)
		return
	} 
	fmt.Printf("[crawlPage] Got %d links\n",len(pageOutgoingLinks))
	for _,aUrl := range pageOutgoingLinks {
		crawlPage(rawBaseURL, aUrl, pages, bodies)
	}
	
}


func crawlSomePage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlSomePage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	parsedURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlSomePage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != parsedURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	// increment if visited
	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL]++
		return
	}

	// mark as visited
	pages[normalizedURL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, parsedURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		crawlSomePage(rawBaseURL, nextURL, pages)
	}
}
