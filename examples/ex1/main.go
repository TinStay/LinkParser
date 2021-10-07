package main

import (
	"fmt"
	"strings"

	link "github.com/Basics/src/github.com/TinStay/LinkParser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
    A link to another page
    <span> some span  </span>
  </a>
  <a href="/page-two">A link to a second page <span>another span</span></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)

	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
