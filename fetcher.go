package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)


var Fetch = fetch

func fetch() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
		fmt.Println(band)
	})
}
