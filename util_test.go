package webpage

import (
	"testing"
)

func TestRemoveExtraSpaces(t *testing.T) {
	str := "  abc\n"

	if removeExtraSpaces(str) != "abc" {
		t.Error("removeExtraSpaces() return 'abc'")
	}
}

func TestRemoveDuplicateLinebreak(t *testing.T) {

	src := `
	aaa




	bbb
	`

	expected := `
	aaa


	bbb
	`

	if removeDuplicateLinebreak(src) != expected {
		t.Error("removeDuplicateLinebreak() return: " + expected)
	}
}
