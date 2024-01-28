package utils

import (
	"fmt"

	"golang.org/x/net/html"
)

func ProcessArticleBody(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "p" {
		processNode(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessArticleBody(c)
	}
}

func processNode(n *html.Node) {
	fmt.Println("here", n.Data)
}
