package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Poems lists the poems for a poet.
func Poems(poet string) (poemURLs []string, err error) {
	listingPage := "http://www.poemhunter.com/" + poet + "/poems/"
	for {
		page, err := fetchPage(listingPage)
		if err != nil {
			return nil, err
		}
		tds := scrape.FindAll(page, func(n *html.Node) bool {
			return n.DataAtom == atom.Td && scrape.ByClass("title")(n)
		})
		for _, td := range tds {
			link, ok := scrape.Find(td, scrape.ByTag(atom.A))
			if !ok {
				continue
			}
			poemURL := "http://www.poemhunter.com/" + scrape.Attr(link, "href")
			poemURLs = append(poemURLs, poemURL)
		}
		next, ok := scrape.Find(page, func(n *html.Node) bool {
			return n.DataAtom == atom.A && strings.HasPrefix(scrape.Text(n), "next page")
		})
		if ok {
			if nextPage := scrape.Attr(next, "href"); nextPage != "" {
				listingPage = "http://www.poemhunter.com/" + nextPage
				continue
			}
		}
		break
	}
	return
}

// TopPoetNames lists the top poets.
func TopPoetNames() (names []string, err error) {
	page, err := fetchPage("http://www.poemhunter.com/p/t/l.asp?p=1&l=Top500")
	if err != nil {
		return nil, err
	}
	links := scrape.FindAll(page, scrape.ByClass("photo"))
	for _, link := range links {
		link := scrape.Attr(link, "href")
		// The href is of the form "/william-shakespeare/".
		names = append(names, link[1:len(link)-1])
	}
	return
}

// PoemText reads a specific poem.
func PoemText(poemURL string) (string, error) {
	resp, err := http.Get(poemURL)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		return "", errors.New("parse page: " + err.Error())
	}
	paragraphs := scrape.FindAll(root, scrape.ByTag(atom.P))
	if len(paragraphs) < 2 {
		return "", errors.New("cannot find paragraph")
	}
	return breakableText(paragraphs[1]), nil
}

func breakableText(el *html.Node) string {
	if el.Type == html.TextNode {
		return strings.TrimSpace(el.Data)
	} else if el.Type == html.ElementNode {
		if el.DataAtom == atom.Br {
			return "\n"
		} else {
			var res string
			child := el.FirstChild
			for child != nil {
				res += breakableText(child)
				child = child.NextSibling
			}
			return res
		}
	}
	return ""
}

func fetchPage(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return html.Parse(resp.Body)
}
