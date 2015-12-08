package webpage

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	page := getPage()
	wp := New(page)

	if wp.Title != "sample-1.html" {
		t.Error("wp.Title) should return 'sample-1.html'")
	}

	if wp.Keywords != "keyword1, keyword2, keyword3" {
		t.Error("wp.Keywords should return 'keyword1, keyword2, keyword3'")
	}

	if wp.Description != "some description" {
		t.Error("wp.Description should return 'some description'")
	}

	if string(wp.Data) != string(page) {
		t.Error("wp.Data should return original data")
	}
}

func TestTitle(t *testing.T) {
	page := getPage()
	wp := New(page)

	if wp.GetTitle() != "sample-1.html" {
		t.Error("Title() should return 'sample-1.html'")
	}
}

func TestKeywords(t *testing.T) {
	page := getPage()
	wp := New(page)

	if wp.GetKeywords() != "keyword1, keyword2, keyword3" {
		t.Error("Keywords() should return 'keyword1, keyword2, keyword3'")
	}
}

func TestDescription(t *testing.T) {
	page := getPage()
	wp := New(page)

	if wp.GetDescription() != "some description" {
		t.Error("Description() should return 'some description'")
	}
}

func TestBodyText(t *testing.T) {
	page := getPage()
	wp := New(page)
	bodyText := wp.GetBodyText()

	if m1, _ := regexp.MatchString("hogehoge", bodyText); !m1 {
		t.Error("BodyText() should contain 'hogehoge'")
	}

	if m2, _ := regexp.MatchString("something text", bodyText); !m2 {
		t.Error("BodyText() should contain 'something text'")
	}
}

func getPage() []byte {
	page, _ := ioutil.ReadFile("./testdata/sample-1.html")
	return page
}
