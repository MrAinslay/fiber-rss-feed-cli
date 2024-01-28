package utils

import (
	"fmt"

	"golang.org/x/net/html"
)

func ProcessArticleBody(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "p" {
		processNode(n.FirstChild)
	}

	if n.Type == html.ElementNode && n.Data == "figure" {
		processNode(n.FirstChild)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessArticleBody(c)
	}
}

func processNode(n *html.Node) {
	switch n.Data {
	case "img":
		for _, a := range n.Attr {
			if a.Key == "src" {
				fmt.Println("Image URL:", a.Val)
			}
		}
	default:
		fmt.Println(n.Data)
	}
}
