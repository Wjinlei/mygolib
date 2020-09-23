package myxml

import (
	"fmt"
	"testing"
)

func TestNewDocument(t *testing.T) {
	doc := NewDocument()
	if err := doc.Document.ReadFromString("<root><a>aaaaaaaaaa</a><b>bbbbbbbbbbbb</b></root>"); err != nil {
		t.Error(err)
	}
}

func TestSelectElement(t *testing.T) {
	doc := NewDocument()
	if err := doc.Document.ReadFromString("<root><a>aaaaaaaaaa</a><b>bbbbbbbbbbbb</b></root>"); err != nil {
		t.Error(err)
	}
	// SelectElement 必须一级一级的找下去
	root := doc.SelectElement("root") // 先找根节点
	a := root.SelectElement("a")      // 再找下级节点
	fmt.Println(a.Text())
}

func TestFindElement(t *testing.T) {
	doc := NewDocument()
	if err := doc.Document.ReadFromString("<root><a>aaaaaaaaaa</a><b>bbbbbbbbbbbb</b></root>"); err != nil {
		t.Error(err)
	}
	// 找root下的b节点
	b := doc.FindElement("/root/b")
	fmt.Println(b.Text())
}
