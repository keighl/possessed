// Package possessed - Golang functions for converting an English string to/from its possessive form.
//
//	import (
//	    ps "github.com/keighl/possessed"
//	    "fmt"
//	)
//
//	fmt.Println(ps.Possess("Dave"))
//	// => Dave's
//
//	fmt.Println(ps.Possess("Gladys"))
//	// => Gladys'
//
//	fmt.Println(ps.Possess("it"))
//	// => its
//
//	fmt.Println(ps.Unpossess("Dave's"))
//	// => Dave
//
//	fmt.Println(ps.Unpossess("Gladys'"))
//	// => Gladys
//
//	fmt.Println(ps.Unpossess("its"))
//	// => it
//
//	// Change the apostrophe character
//	ps.ApostropheChar = `’`
//	fmt.Println(ps.Possess("Dave"))
//	// => Dave’s
//
package possessed

import (
	"regexp"
	"strings"
)

var (
	// ApostropheChar is the character to make strings possessive
	// e.g. Gladys', or Gladys’
	ApostropheChar = `'`
	// UnpossessRegex tells Unpossess() how to find possessive forms
	UnpossessRegex   = regexp.MustCompile(`('s|’s|'|’)$`)
)

// Possess returns the string in the possessive form
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
		return s + ApostropheChar
	}

	return s + ApostropheChar + "s"
}

// Unpossess returns the "unpossessed" form of the string (if it's possessive to begin with)
func Unpossess(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if strings.ToLower(s) == "its" {
		return s[0:2]
	}

	return UnpossessRegex.ReplaceAllString(s, "")
}
