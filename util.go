package webpage

import (
	"regexp"
)

func removeExtraSpaces(text string) string {
	spaces := regexp.MustCompile("(^[\n\t\\s]+|[\n\t\\s]+$)")
	return spaces.ReplaceAllString(text, "")
}

func removeDuplicateLinebreak(text string) string {
	linebreak := regexp.MustCompile("\n{4,}")
	return linebreak.ReplaceAllString(text, "\n\n\n")
}
