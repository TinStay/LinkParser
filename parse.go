package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a <a href> tag
type Link struct {
	Href string
	Text string
}

// Take in an HTML document and return a slice of links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	// depth-first search
	// dfs(doc, "")

	// Find link nodes in doc
	nodes := linkNodes(doc)

	var resultLinks []Link

	// Build links
	for _, node := range nodes {
		resultLinks = append(resultLinks, buildLink(node))
	}

	// Return links
	return resultLinks, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	// Get link href
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	// Parse link text
	ret.Text = getText(n)

	return ret
}

func getText(n *html.Node) string {
	// Base cases
	if n.Type == html.TextNode {
		return n.Data
	} 
	
	// e.g Comments
	if n.Type != html.ElementNode {
		return ""
	} 
	
	var ret string

	// Recursive call
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c)
	}

	// Return trimmed string
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	// Base case
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node

	// Recurring loop
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

// func dfs(n *html.Node, padding string) {
// 	// Formatting
// 	msg := n.Data
// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}

// 	fmt.Println(padding, msg)

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}

// }
