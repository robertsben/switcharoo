package main

import (
	"encoding/xml"
	"io"
	// "fmt"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

// <tree.go>

type Attribute struct {
	Label 	string
	Value 	string
}

type Attributes []*Attribute

type Element struct {
	Parent		*Element
	Children 	Elements
	Label 		string
	Attrs		Attributes
	Data 		string
}

type Elements []*Element

func (self *Element) AddAttribute(attr xml.Attr) {
	self.Attrs = append(self.Attrs, &Attribute{Label: attr.Name.Local, Value: attr.Value})
}

func (self *Element) AddChild(child *Element) {
	self.Children = append(self.Children, child)
}

func (self *Element) AddSelfToParentsChildren() {
	self.Parent.AddChild(self)
}

// </tree.go>



type Decoder struct {
	r 	io.Reader
	err error
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode(root *Element) error {
	decoder := xml.NewDecoder(dec.r)
	element := root

	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		spew.Dump(token)

		switch curr_tok := token.(type) {

		case xml.StartElement:
			element = &Element{
				Parent: element,
				Label: curr_tok.Name.Local,
			}
			for _, attr := range curr_tok.Attr {
				element.AddAttribute(attr)
			}
			element.AddSelfToParentsChildren()

		case xml.CharData:
			spew.Dump(parseCharToken(string(xml.CharData(curr_tok))))
			element.Data = parseCharToken(string(xml.CharData(curr_tok)))
		case xml.EndElement:
			element = element.Parent
		}
			
	}
	spew.Dump(root)
	return nil
}

func parseCharToken(charTok string) string {
	return strings.TrimSpace(charTok)
}
