package webpage

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WebPage []byte

func (wp *WebPage) Html() string {
	return string(*wp)
}

func (wp *WebPage) Text() string {
	doc := wp.doc()
	return doc.Text()
}

func (wp *WebPage) Title() string {
	doc := wp.doc()
	title := doc.Find("title")
	return title.Text()
}

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

func (wp *WebPage) BodyHtml() string {
	doc := wp.doc()
	body := doc.Find("body")
	html, _ := body.Html()
	return html
}

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
