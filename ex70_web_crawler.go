package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
type Crawler struct {
	fetcher Fetcher
	fetched map[string]bool
}

func (crawler Crawler) Crawl(url string, depth int) {
	if depth <= 0 {
		return
	}

	body, urls, err := crawler.fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	crawler.fetched[url] = true

	fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup
	for _, url := range urls {
		// ignore url if already fetched
		if !crawler.fetched[url] {
			wg.Add(1)
			go func(url string) {
				crawler.Crawl(url, depth-1)
				wg.Done()
			}(url)
		}
	}
	wg.Wait()
}

func main() {
	crawler := Crawler{fetcher, make(map[string]bool)}
	crawler.Crawl("http://golang.org/", 4)
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
