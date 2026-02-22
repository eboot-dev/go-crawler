package main
import (
	"net/url"
	"fmt"
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