package main

import (
	"fmt"
	"os"
	"strconv"
	_"net/url"
	_"sync"
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
	return
}

const (
	maxConcurrency = 8	
)

func main() {
	fmt.Println("Hello, World! From cfg.crawlPage()")
	argc := len(os.Args[1:])
	if argc < 3 {
		fmt.Println("too few arguments provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}
	if argc > 3 {
		fmt.Println("too many arguments provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}
	BASE_URL := os.Args[1]

	// Convert string to int
	num, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Conversion Error:", err)
		os.Exit(1)
	}
	maxConcurrency := num

	cfg,err := configure(BASE_URL,maxConcurrency)
	if err != nil {
		fmt.Errorf("Couldn't create config: %v", err)
		return
	}
	// Convert string to int
	num2, err2 := strconv.Atoi(os.Args[3])
	if err2 != nil {
		fmt.Println("Conversion Error:", err2)
		os.Exit(1)
	}
	
	cfg.maxPages = num2
	cfg.wg.Add(1)
	go cfg.crawlPage(BASE_URL)
	cfg.wg.Wait()
	
	fmt.Println("Done Waiting!")
	
	cfg.mu.Lock()
	err = writeCSVReport(cfg.pages, "report.csv")
	cfg.mu.Unlock()
	if err != nil {
		fmt.Errorf("Couldn't write CSV report: %v", err)
		return
	}
}


func mainCrawl() {
	mainCrawlPage()
	BASE_URL := os.Args[1]
	// pages := make(map[string]int)
	// fmt.Println("Hello, World! From crawlSomePage")
	// crawlSomePage(BASE_URL, BASE_URL, pages)
	// fmt.Println("\nDONE crawlSomePage\n")
	// for normalizedURL, count := range pages {
	// 	fmt.Printf("%d - %s\n", count, normalizedURL)
	// }
	fmt.Println("Hello, World! From cfg.crawlPage()")
	cfg,err := configure(BASE_URL,maxConcurrency)
	if err != nil {
		fmt.Errorf("Couldn't create config: %v", err)
		return
	}
	cfg.wg.Add(1)
	go cfg.crawlPage(BASE_URL)
	cfg.wg.Wait()

	// for normalizedURL, count := range cfg.pages {
	// 	fmt.Printf("%d - %s\n", count, normalizedURL)
	// }
}