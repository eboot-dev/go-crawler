package main

import (
	"fmt"
	"net/http"
	_"net/url"
	"errors"
	"io"
	"strings"
)


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
	defer res.Body.Close()

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)
	if res.StatusCode > 399 {
		fmt.Printf("client: bad status code: %s [%v]\n",http.StatusText(res.StatusCode), res.StatusCode, )
		return "", fmt.Errorf("%s",http.StatusText(res.StatusCode))
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