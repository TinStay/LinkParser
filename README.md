# HTML Link Parser

## Exercise details
The goal of exercise is to create a package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link there should be a data structure returned that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc. Links are nested in different HTML elements.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

In situations like these we want to get output like:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```


### Notes

**1. Use the x/net/html package**

**2. Ignore nested links**

Ignore any links nested inside of another link. Eg with following HTML:

```html
<a href="#">
  Something here <a href="/dog">nested dog link</a>
</a>
```

