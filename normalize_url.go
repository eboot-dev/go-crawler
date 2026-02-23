package main
import (
	"net/url"
	"fmt"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}


func oldNormalizeURL(rawURL string) (string, error) {
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


