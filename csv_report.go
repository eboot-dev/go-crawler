package main

import (
	"encoding/csv"
	"os"
	"strings"
	"fmt"
)
/*
Implement the `writeCSVReport` function. It should:
- Create a CSV with these column headers: 
	page_url, h1, first_paragraph, outgoing_link_urls, image_urls
- Join slice values with semicolons (;) for the link and image URL columns

Here are some tips to get you started:

Open the file with `os.Create(filename)`
Use `csv.NewWriter(file)` to write the data
Write the header row with `writer.Write([]string{...})`
For each PageData in the map, write a row with its fields
Join slices with semicolons: strings.Join(page.OutgoingLinks, ";")

URL 			string
	H1				string
	FirstParagraph 	string	
	OutgoingLinks	[]string
	ImageURLs		[]string

*/
func writeCSVReport(pages map[string]PageData, filename string) error {
	file,err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error! Can't create file '%s' [%s]",filename,err)
		return err
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls",})

	for _,page := range pages {
		iurls := "''"
		if page.ImageURLs != nil {
			iurls = strings.Join(page.ImageURLs, ";")
		}
		links := "''"
		if page.OutgoingLinks != nil {
			links = strings.Join(page.OutgoingLinks, ";")
		}
		writer.Write([]string{
							page.URL,
							page.H1,
							page.FirstParagraph,
							links,
							iurls})
	}
	return nil

}