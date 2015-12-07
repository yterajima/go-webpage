package webpage

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// WebPage is the type that represents the pages of the website
type WebPage []byte

// Html return html string of page
func (wp *WebPage) Html() string {
	return string(*wp)
}

// Text return full-text of page
func (wp *WebPage) Text() string {
	doc := wp.doc()
	return doc.Text()
}

// Title return text in <title> tag
func (wp *WebPage) Title() string {
	doc := wp.doc()
	title := doc.Find("title")
	return title.Text()
}

// Keywords return content of <meta name="keywords" ~>
func (wp *WebPage) Keywords() string {
	doc := wp.doc()

	keyword := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); strings.EqualFold(name, "keywords") {
			keyword = s.AttrOr("content", "")
		}
	})
	return keyword
}

// Description return content of <meta name="description" ~>
func (wp *WebPage) Description() string {
	doc := wp.doc()

	description := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); strings.EqualFold(name, "description") {
			description = s.AttrOr("content", "")
		}
	})
	return description
}

// BodyHtml return full html text in <body> tag
func (wp *WebPage) BodyHtml() string {
	doc := wp.doc()
	body := doc.Find("body")
	html, _ := body.Html()
	return html
}

// BodyText return text in <body> tag
func (wp *WebPage) BodyText() string {
	doc := wp.doc()
	body := doc.Find("body")
	return body.Text()
}

func (wp *WebPage) doc() *goquery.Document {
	reader := bytes.NewReader(*wp)
	doc, _ := goquery.NewDocumentFromReader(reader)
	return doc
}
