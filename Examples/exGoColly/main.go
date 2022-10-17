package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type hashnodeScraper struct {
	Articles []Article
}

type Article struct {
	Title       string
	Description string
	Link        string
	Author      string
	Date        string
}

//create Hashnode scrapper struct
func NewHashnodeScraper() *hashnodeScraper {
	return &hashnodeScraper{
		Articles: []Article{},
	}
}

func (h *hashnodeScraper) ScrapUrl(url string, domains ...string) {

	c := colly.NewCollector(colly.AllowedDomains(domains...))
	titles := []string{}
	links := []string{}
	descriptions := []string{}
	authors := []string{}
	dates := []string{}

	//links & titles
	c.OnHTML("div.css-4gdbui div.css-1wg9be8 div.css-16fbhyp h1.css-1j1qyv3 a.css-4zleql", func(h *colly.HTMLElement) {
		titles = append(titles, h.Text)
		links = append(links, h.Attr("href"))
	})

	//description
	c.OnHTML("div.css-4gdbui div.css-1wg9be8 div.css-16fbhyp p.css-1072ocs a.css-4zleql", func(h *colly.HTMLElement) {
		descriptions = append(descriptions, h.Text)
	})

	//author
	c.OnHTML("div.css-4gdbui div.css-dxz0om div.css-tel74u div.css-2wkyxu div.css-1ajtyzd a.css-c3r4j7", func(h *colly.HTMLElement) {
		authors = append(authors, h.Text)
	})

	//date
	c.OnHTML("div.css-4gdbui div.css-dxz0om div.css-tel74u div.css-2wkyxu div.css-1n08q4e a.css-1u6dh35", func(h *colly.HTMLElement) {
		dates = append(dates, h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error: %s\n", err.Error())
	})

	c.Visit(url)

	fmt.Println(len(titles))
	fmt.Println(len(descriptions))
	fmt.Println(len(links))
	fmt.Println(len(authors))
	fmt.Println(len(dates))

	h.CombineData(titles, descriptions, links, authors, dates)

}

//Combine articles data together in structs
func (h *hashnodeScraper) CombineData(title, description, link, author, date []string) {

	for i := 0; i < len(title); i++ {

		a := Article{
			Title:       title[i],
			Description: description[i],
			Link:        link[i],
			Author:      author[i],
			Date:        date[i],
		}

		h.Articles = append(h.Articles, a)

	}

}

func main() {

	url := "https://hashnode.com/n/go"
	domains := []string{"www.hashnode.com", "hashnode.com"}

	hs := NewHashnodeScraper()
	hs.ScrapUrl(url, domains...)

	fmt.Printf("%#v", hs)
}
