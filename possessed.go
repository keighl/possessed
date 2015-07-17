package possessed

import (
	"strings"
	"regexp"
)

var (
	APOSTROPHE_CHAR = `'`
	DEPOSSESS_EXP = regexp.MustCompile(`('s|’s|'|’)$`)
)

func Possess(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if strings.ToLower(s) == "it" {
		return s + "s"
	}

	lastChar := s[len(s)-1:]
	if lastChar == "s" {
		return s + APOSTROPHE_CHAR
	}

	return s + APOSTROPHE_CHAR + "s"
}

func Unpossess(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if strings.ToLower(s) == "its" {
		return s[0:2]
	}

	return DEPOSSESS_EXP.ReplaceAllString(s, "")
}
