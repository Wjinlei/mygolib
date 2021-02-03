package myvalidator

import "regexp"

func regexpCompile(str string) *regexp.Regexp {
	return regexp.MustCompile("^" + str + "$")
}
