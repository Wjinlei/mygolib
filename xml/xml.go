package xml

import (
	"strconv"

	"github.com/beevik/etree"
)

type Document struct {
	Document *etree.Document
}

type Element struct {
	Element *etree.Element
}

// NewDocument 获取新的文档流
func NewDocument() *Document {
	return &Document{Document: etree.NewDocument()}
}

// SelectElements 返回选择的元素
func (d *Document) SelectElement(name string) *Element {
	return &Element{Element: d.Document.SelectElement(name)}
}

// SelectElements 返回选择的元素集合
func (d *Document) SelectElements(name string) []*Element {
	var elements []*Element
	etreeElements := d.Document.SelectElements(name)
	for _, element := range etreeElements {
		elements = append(elements, &Element{Element: element})
	}
	return elements
}

// FindElement 根据路径返回元素
func (d *Document) FindElement(path string) *Element {
	return &Element{Element: d.Document.FindElement(path)}
}

// FindElements 根据路径返回元素集合
func (d *Document) FindElements(path string) []*Element {
	var elements []*Element
	etreeElements := d.Document.FindElements(path)
	for _, element := range etreeElements {
		elements = append(elements, &Element{Element: element})
	}
	return elements
}

// SelectElements 返回选择的元素
func (e *Element) SelectElement(name string) *Element {
	return &Element{Element: e.Element.SelectElement(name)}
}

// SelectElements 返回选择的元素集合
func (e *Element) SelectElements(name string) []*Element {
	var elements []*Element
	etreeElements := e.Element.SelectElements(name)
	for _, element := range etreeElements {
		elements = append(elements, &Element{Element: element})
	}
	return elements
}

// FindElement 根据路径返回元素
func (e *Element) FindElement(path string) *Element {
	return &Element{Element: e.Element.FindElement(path)}
}

// FindElements 根据路径返回元素集合
func (e *Element) FindElements(path string) []*Element {
	var elements []*Element
	etreeElements := e.Element.FindElements(path)
	for _, element := range etreeElements {
		elements = append(elements, &Element{Element: element})
	}
	return elements
}

func (e *Element) Text() string {
	return e.Element.Text()
}

func (e *Element) Uint() uint {
	return uint(e.Int())
}

func (e *Element) Int() int {
	result, err := strconv.Atoi(e.Element.Text())
	if err != nil {
		return 0
	}
	return result
}
