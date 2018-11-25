package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	// SafeCounter is safe to use concurrently.
	type SafeCounter struct {
		v   map[string]int
		mux sync.Mutex
	}

	var worker func(string, int, Fetcher, chan int, SafeCounter)
	worker = func(url string, depth int, fetcher Fetcher, quit chan int, c SafeCounter) {

		defer close(quit)

		if depth <= 0 {
			return
		}

		c.mux.Lock()
		c.v[url]++
		c.mux.Unlock()

		if c.v[url] > 1 {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		num := len(urls)
		if num > 0 {
			chs := make([]chan int, num)
			i := 0
			for _, u := range urls {
				chs[i] = make(chan int)
				go worker(u, depth-1, fetcher, chs[i], c)
				i++
			}
			for i = 0; i < num; i++ {
				<-chs[i]
			}
		}
	}

	c := SafeCounter{v: make(map[string]int)}
	ch := make(chan int)
	worker(url, depth, fetcher, ch, c)
	<-ch

}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
