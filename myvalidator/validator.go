package myvalidator

import (
	"regexp"
)

const (
	// 匹配IP4
	ip4Reg = `((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`

	// 匹配IP6
	ip6Reg = `(([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:)` +
		`{6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)` +
		`(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:)` +
		`{5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)` +
		`(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:)` +
		`{4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:` +
		`((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d))` +
		`{3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|` +
		`((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)` +
		`(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:)` +
		`{2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:` +
		`((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d))` +
		`{3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|` +
		`((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)` +
		`(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|` +
		`((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)` +
		`(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))`

	// 同时匹配IPV4和IPV6
	ipReg = "(" + ip4Reg + ")|(" + ip6Reg + ")"

	// 匹配域名
	domainReg = `(\*\.){0,1}[a-zA-Z0-9_-][a-zA-Z0-9_-]{0,62}` +
		`(\.[a-zA-Z0-9][a-zA-Z0-9_-]{0,62})*(\.[a-zA-Z][a-zA-Z0-9-]{0,20}){1}`

	// 匹配URL
	urlReg = `((https|http|ftp|rtsp|mms)?://)?` + // 协议
		`(([0-9a-zA-Z]+:)?[0-9a-zA-Z_-]+@)?` + // pwd:user@
		"(" + ip4Reg + "|(" + domainReg + "))" + // IP或域名
		`(:\d{1,4})?` + // 端口
		`(/+[a-zA-Z0-9][a-zA-Z0-9_.-]*/*)*` + // path
		`(\?([a-zA-Z0-9_-]+(=[a-zA-Z0-9_-]*)*)*)*` // query

	// 匹配邮箱
	emailReg = `[\w.-]+@[\w_-]+\w{1,}[\.\w-]+`
)

func regexpCompile(str string) *regexp.Regexp {
	return regexp.MustCompile("^" + str + "$")
}

var (
	ip4    = regexpCompile(ip4Reg)
	ip6    = regexpCompile(ip6Reg)
	ip     = regexpCompile(ipReg)
	url    = regexpCompile(urlReg)
	email  = regexpCompile(emailReg)
	domain = regexpCompile(domainReg)
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
