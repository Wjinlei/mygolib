package myxml

import (
	"github.com/beevik/etree"
)

type Document struct {
	Document *etree.Document
}

// NewDocument 获取新的文档流
func NewDocument() *Document {
	return &Document{Document: etree.NewDocument()}
}

// SelectElements 返回选择的元素
func (d *Document) SelectElement(name string) *etree.Element {
	return d.Document.SelectElement(name)
}

// SelectElements 返回选择的元素集合
func (d *Document) SelectElements(name string) []*etree.Element {
	return d.Document.SelectElements(name)
}

// FindElement 根据路径返回元素
func (d *Document) FindElement(path string) *etree.Element {
	return d.Document.FindElement(path)
}

// FindElements 根据路径返回元素集合
func (d *Document) FindElements(path string) []*etree.Element {
	return d.Document.FindElements(path)
}
