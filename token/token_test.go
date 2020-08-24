package token

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	_, err := GenerateToken("wjl", "admin")
	if err != nil {
		t.Error(err)
	}
}

func TestParseToken(t *testing.T) {
	token, err := GenerateToken("wjl", "admin")
	if err != nil {
		t.Error(err)
	}
	if _, err := ParseToken(token); err != nil {
		t.Error(err)
	}
}
