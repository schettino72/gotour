package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
type Page struct {
	url string
	depth int
	err error
	body string
	links []string
}

func NewPage (url string, depth int) (*Page){
	return &Page{url, depth, nil, "", nil}
}



func FetchPage(page *Page, fetcher Fetcher, result_q chan *Page) {
	page.body, page.links, page.err = fetcher.Fetch(page.url)
	result_q <- page
}

func Crawl(urls []string, max_depth int){
	result_q := make(chan *Page)
	fetched := make(map[string]bool)

	// kick start with initial urls
	fetching := len(urls)
	for _, url := range(urls){
		go FetchPage(NewPage(url, 0), fetcher, result_q)
	}

	// loop until nothing to fetch
	for {
		page := <- result_q
		fetched[page.url] = true

		if page.err != nil {
			fmt.Println(page.err)
		} else {
			fmt.Printf("found: %s %q\n", page.url, page.body)
		}

		// fetch linked pages
		next_depth := page.depth + 1
		if next_depth < max_depth {
			for _, url := range page.links {
				// ignore url if already fetched
				if !fetched[url] {
					fetching++
					go FetchPage(NewPage(url, next_depth), fetcher, result_q)
				}
			}
		}

		// done with this page, check nothing left
		fetching--
		if fetching == 0 {
			break
		}
	}
}

func main() {
	Crawl([]string{ "http://golang.org/" },  4)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		time.Sleep(1 * time.Second)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
