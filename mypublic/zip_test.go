package mypublic

import "testing"

/*
func TestTGZ(t *testing.T) {
	if err := TGZ("/home/wangjl/tmps/test", "test.tar.gz"); err != nil {
		t.Error(err)
	}
}

func TestZIP(t *testing.T) {
	if err := ZIP("/home/wangjl/tmps/test", "test.zip", "gbk"); err != nil {
		t.Error(err)
	}
}
*/

func TestZIPDecrypt(t *testing.T) {
	if err := ZIPDecrypt("/home/wangjl/Downloads/wwwroot_20210310155454.zip", "/home/wangjl/tmps/wwwroot", "", "gbk"); err != nil {
		t.Error(err)
	}
}

/*
func TestTGZDecrypt(t *testing.T) {
	if err := TGZDecrypt("test.tar.gz", "/tmp/test", "utf8"); err != nil {
		t.Error(err)
	}
}
*/
