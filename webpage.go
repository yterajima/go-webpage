package webpage

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// WebPage is the type that represents the pages of the website
type WebPage struct {
	Title       string
	Keywords    string
	Description string
	BodyText    string
}

// New generate WebPage structure
func New(page []byte) WebPage {
	doc := getDoc(page)

	wp := WebPage{
		Title:       getTitle(doc),
		Keywords:    getKeywords(doc),
		Description: getDescription(doc),
		BodyText:    getBodyText(doc),
	}

	return wp
}

// Title return text in <title> tag
func (wp *WebPage) GetTitle() string {
	return wp.Title
}

// Keywords return content of <meta name="keywords" ~>
func (wp *WebPage) GetKeywords() string {
	return wp.Keywords
}

// Description return content of <meta name="description" ~>
func (wp *WebPage) GetDescription() string {
	return wp.Description
}

// BodyText return text in <body> tag
func (wp *WebPage) GetBodyText() string {
	return wp.BodyText
}

func getDoc(page []byte) *goquery.Document {
	reader := bytes.NewReader(page)
	doc, _ := goquery.NewDocumentFromReader(reader)
	return doc
}

func getTitle(doc *goquery.Document) string {
	title := doc.Find("title")
	return title.Text()
}

func getKeywords(doc *goquery.Document) string {
	keyword := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); strings.EqualFold(name, "keywords") {
			keyword = s.AttrOr("content", "")
		}
	})
	return keyword
}

func getDescription(doc *goquery.Document) string {
	description := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); strings.EqualFold(name, "description") {
			description = s.AttrOr("content", "")
		}
	})
	return description
}

func getBodyText(doc *goquery.Document) string {
	body := doc.Find("body")
	return body.Text()
}
