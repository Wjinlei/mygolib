package encoding

import "github.com/miolini/gopunycode"

func GetPunyCode(old string) string {
	punycode, err := gopunycode.ToASCII(old)
	if err != nil {
		return old
	}
	return punycode
}
