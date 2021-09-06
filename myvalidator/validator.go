package myvalidator

import (
	"regexp"
	"unicode"

	"github.com/caixw/lib.go/assert"
)

func isMatch(exp *regexp.Regexp, val interface{}) bool {
	switch v := val.(type) {
	case []rune:
		return exp.MatchString(string(v))
	case []byte:
		return exp.Match(v)
	case string:
		return exp.MatchString(v)
	default:
		return false
	}
}

// 验证一个值是否标准的URL格式。支持IP和域名等格式
func IsURL(val interface{}) bool {
	return isMatch(url, val)
}

// 验证一个值是否为IP，可验证IP4和IP6
func IsIP(val interface{}) bool {
	return isMatch(ip, val)
}

// 验证一个值是否为IPV6
func IsIP6(val interface{}) bool {
	return isMatch(ip6, val)
}

// 验证一个值是否为IPV4
func IsIP4(val interface{}) bool {
	return isMatch(ip4, val)
}

// 验证一个值是否为域名
func IsDomain(val interface{}) bool {
	return isMatch(domain, val)
}

// 验证一个值是否匹配一个邮箱。
func IsEmail(val interface{}) bool {
	return isMatch(email, val)
}

func IsDigit(str string) bool {
	for _, r := range str {
		if ok := unicode.IsDigit(r); !ok {
			return false
		}
	}
	return true
}

func IsLetter(str string) bool {
	for _, r := range str {
		if ok := unicode.IsLetter(r); !ok {
			return false
		}
	}
	return true
}

func IsLetterOrDigit(str string) bool {
	for _, r := range str {
		if ok := unicode.IsDigit(r); !ok {
			if ok := unicode.IsLetter(r); !ok {
				return false
			}
		}
	}
	return true
}

func IsLetterOrDigitOrUnderscoreOrUnderlineOrDecimal(str string) bool {
	for _, r := range str {
		if ok := unicode.IsDigit(r); !ok {
			if ok := unicode.IsLetter(r); !ok {
				if r != '_' && r != '-' && r != '.' {
					return false
				}
			}
		}
	}
	return true
}

func IsContains(container, item interface{}) bool {
	return assert.IsContains(container, item)
}
