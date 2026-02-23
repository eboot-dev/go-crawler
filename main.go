package main

import (
	"fmt"
	"os"
	_"net/url"
)


func mainGetHTML() {
	fmt.Println("Hello, World!")
	argc := len(os.Args[1:])
	if argc < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if argc > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	BASE_URL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n",BASE_URL)
	body, err := getHTML(BASE_URL)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("[crawler] response body: %s\n", body)
}

func mainCrawlPage() {
	fmt.Println("Hello, World! From mainCrawlPage")
	argc := len(os.Args[1:])
	if argc < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if argc > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	BASE_URL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n",BASE_URL)
 	pages := map[string]int{}
 	bodies := map[string]string{}
	crawlPage(BASE_URL, BASE_URL, pages,bodies)
	
	fmt.Printf("[crawler] response bodies: %d\n", len(bodies))
}
func main() {
	mainCrawlPage()
	BASE_URL := os.Args[1]
	pages := make(map[string]int)
	fmt.Println("Hello, World! From crawlSomePage")
	crawlSomePage(BASE_URL, BASE_URL, pages)
	fmt.Println("\nDONE crawlSomePage\n")
	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}