package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url"
	"errors"
	"io"
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
		// aParsedURL , err := url.Parse(strings.TrimSpace(aUrl))
		// if err != nil {
		// 	fmt.Printf("[crawlPage] could not parse link URL: %s\n", err)
		// 	continue
		// }
		// _,ok := pages[normaliseURL(aParsedURL.String())]
		// if !ok {
		// 	crawlPage(trimBase, aUrl, pages, bodies)	
		// }
	}
	
}


// direct html access
func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return "",err
	}
	req.Header.Set("User-Agent", "GoCraw/0.0.1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return "",err
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)
	if res.StatusCode > 400 {
		fmt.Printf("client: bad status code: %s [%v]\n",http.StatusText(res.StatusCode), res.StatusCode, )
		return "",errors.New(http.StatusText(res.StatusCode))
	}
	ct := res.Header.Get("content-type")
	if !strings.Contains(ct, "text/html") {
		fmt.Printf("client: bad content-type: %s \n",ct)
		return "",errors.New("bad content-type "+ct)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return "",err
	}
	// fmt.Printf("client: response body: %s\n", resBody)
	return string(resBody),nil
}


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
 	pages := map[string]int{}
 	bodies := map[string]string{}
	crawlPage(BASE_URL, BASE_URL, pages,bodies)
	
	fmt.Printf("[crawler] response bodies: %d\n", len(bodies))
}
func main() {
	mainCrawlPage()
}