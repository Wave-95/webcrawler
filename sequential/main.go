package main

import (
	"fmt"
	"time"
)

//Create a program that sequentially crawls the web (via a mock data set)
//Implement some time delays for each web crawl to simulate crawl time

//The crawler should start with a root url. After fetching the url, the crawler should return
//a list of new urls to crawl against.

//Mock results -- map of url to urls

// LoadTime is a mock value to be used by time.Sleep to simulate wait times for pages to fully load and be crawled against
type Page struct {
	Url      string
	LoadTime int
}

// Load method for Page type makes use of a fake data structure to simulate web crawling. Some key value pairs will map back to root nodes
func (p Page) Load() []Page {
	if pages, ok := mockResults[p.Url]; ok {
		return pages
	}
	return []Page{}
}

var mockResults = map[string][]Page{
	"https://golang.org/": {
		Page{"https://golang.org/pkg/", 1000},
		Page{"https://golang.org/cmd/", 900},
	},
	"https://golang.org/pkg/": {
		Page{"https://golang.org/", 500},
		Page{"https://golang.org/cmd/", 900},
		Page{"https://golang.org/pkg/fmt/", 1500},
		Page{"https://golang.org/pkg/os/", 1400},
		Page{"https://golang.org/pkg/strconv/", 2000},
		Page{"https://golang.org/pkg/crypto/", 1200},
		Page{"https://golang.org/pkg/image/", 1900},
	},
	"https://golang.org/pkg/fmt/": {
		Page{"https://golang.org/", 500},
		Page{"https://golang.org/pkg/", 1000},
	},
	"https://golang.org/pkg/os/": {
		Page{"https://golang.org/", 500},
		Page{"https://golang.org/pkg/", 1000},
	},
	"https://golang.org/pkg/image/": {
		Page{"https://golang.org/pkg/image/Alpha", 1400},
		Page{"https://golang.org/pkg/image/Alpha16", 1600},
		Page{"https://golang.org/pkg/image/CMYK", 1300},
		Page{"https://golang.org/pkg/image/Config", 1500},
	},
}

var urlExists = make(map[string]bool)

// Recursive crawl function to iterate through map of mock results. Will stop at a certain depth.
func Crawl(page Page, depth int) {
	if depth <= 0 {
		return
	}
	if urlExists[page.Url] {
		return
	}
	urlExists[page.Url] = true
	time.Sleep(time.Duration(page.LoadTime * int(time.Millisecond)))
	pages := page.Load()
	if len(pages) > 0 {
		for _, childPage := range pages {
			fmt.Printf("[%s]---------------------->[%s]\n", page.Url, childPage.Url)
			Crawl(childPage, depth-1)
		}
	} else {
		fmt.Printf("[%s]----------------------> NONE FOUND\n", page.Url)
	}
	return
}

// Measure time for program to finish
func main() {
	start := time.Now()
	root := Page{"https://golang.org/", 500}
	Crawl(root, 4)
	duration := time.Since(start)
	fmt.Println(duration)
}
