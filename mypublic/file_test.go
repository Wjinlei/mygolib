package mypublic

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	if err := DownloadFile("http://d.hws.com/linux/master/script/install.sh", "./install.sh"); err != nil {
		t.Error(err)
	}
}

func TestWriteFile(t *testing.T) {
	if err := WriteFile("./testfile1.txt", "TestWriteFile"); err != nil {
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	_, err := ReadFile("./testfile1.txt")
	if err != nil {
		t.Error(err)
	}
}

func TestReadLines(t *testing.T) {
	lines, err := ReadLines("./.sayHello.sh")
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(lines[1], "echo \"Hello\"") {
		t.Error("could not read correctly")
	}
}

func TestReadLinesOffsetN(t *testing.T) {
	lines, err := ReadLinesOffsetN("./.sayHello.sh", 0, 2)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(lines[1], "echo \"Hello\"") {
		t.Error("could not read correctly")
	}
}

func TestMoveFile(t *testing.T) {
	if err := Move("./testfile1.txt", "./testfile1.txt.bak"); err != nil {
		t.Error(err)
	}
}

func TestCopyFile(t *testing.T) {
	if err := Copy("file.go", "file.go.bak"); err != nil {
		t.Error(err)
	}
}

func TestParseMode(t *testing.T) {
	stat, err := os.Stat("file.go")
	if err != nil {
		t.Error(err)
	}
	u, g, o := ParseMode(stat.Mode())
	fmt.Printf("%d%d%d\n", u, g, o)
}

func TestChmod(t *testing.T) {
	if err := Chmod("/home/wangjl/tmps/dockerfiles", 0777, true); err != nil {
		t.Error(err)
	}
}
