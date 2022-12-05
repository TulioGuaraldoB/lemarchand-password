package formatter

import "regexp"

func OnlySpecialCharacters(str string) string {
	newString := str
	strRegexp := regexp.MustCompile(`[^!#@$%^&*()-+\\/{}]`)
	newString = strRegexp.ReplaceAllString(str, "")

	return newString
}
