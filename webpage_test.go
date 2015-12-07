package webpage

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestHtml(t *testing.T) {
	wp := getWebPage()
	page := getPage()

	if wp.Html() != string(page) {
		t.Error("Html() should return HTML data")
	}
}

func TestText(t *testing.T) {
	wp := getWebPage()
	text := wp.Text()

	if m1, _ := regexp.MatchString("sample-1.html", text); !m1 {
		t.Error("Text() should contain 'sample-1.html'")
	}

	if m2, _ := regexp.MatchString("hogehoge", text); !m2 {
		t.Error("Text() should contain 'hogehoge'")
	}

	if m3, _ := regexp.MatchString("something text", text); !m3 {
		t.Error("Text() should contain 'something text'")
	}
}

func TestTitle(t *testing.T) {
	wp := getWebPage()

	if wp.Title() != "sample-1.html" {
		t.Error("Title() should return 'sample-1.html'")
	}
}

func TestKeywords(t *testing.T) {
	wp := getWebPage()

	if wp.Keywords() != "keyword1, keyword2, keyword3" {
		t.Error("Keywords() should return 'keyword1, keyword2, keyword3'")
	}
}

func TestDescription(t *testing.T) {
	wp := getWebPage()

	if wp.Description() != "some description" {
		t.Error("Description() should return 'some description'")
	}
}

func TestBodyHtml(t *testing.T) {
	wp := getWebPage()
	bodyHtml := wp.BodyHtml()

	if m1, _ := regexp.MatchString("<div class=\"test\">hogehoge</div>", bodyHtml); !m1 {
		t.Error("BodyHtml() should contain '<div class=\"test\">hogehoge</div>'")
	}

	if m2, _ := regexp.MatchString("<p>something text</p>", bodyHtml); !m2 {
		t.Error("BodyHtml() should contain '<p>something text</p>'")
	}
}

func TestBodyText(t *testing.T) {
	wp := getWebPage()
	bodyText := wp.BodyText()

	if m1, _ := regexp.MatchString("hogehoge", bodyText); !m1 {
		t.Error("BodyText() should contain 'hogehoge'")
	}

	if m2, _ := regexp.MatchString("something text", bodyText); !m2 {
		t.Error("BodyText() should contain 'something text'")
	}
}

func getWebPage() WebPage {
	page := getPage()
	return WebPage(page)
}

func getPage() []byte {
	page, _ := ioutil.ReadFile("./testdata/sample-1.html")
	return page
}
