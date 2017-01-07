package main

import (
	"encoding/xml"
	"io"
	// "fmt"
	"github.com/davecgh/go-spew/spew"
)

// <tree.go>

type Attribute struct {
	Label 	string
	Value 	string
}

type Attributes []*Attribute

func (self *Element) AddAttribute(attr xml.Attr) {
	self.Attrs = append(self.Attrs, &Attribute{Label: attr.Name.Local, Value: attr.Value})
}

type Node struct {
	Children 	Nodes
	Label		string
}

type Nodes []*Node

func (self *Node) AddChild(child *Node) {
	self.Children = append(self.Children, child)
}

type Element struct {
	Parent		*Node
	Self 		*Node
	Attrs		Attributes
	Data 		string
}

func (self *Element) AddSelfToParentsChildren() {
	self.Parent.AddChild(self.Self)
}

// </tree.go>



type Decoder struct {
	r 	io.Reader
	err error
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode(root *Node) error {
	decoder := xml.NewDecoder(dec.r)

	element := &Element{
		Parent: nil,
		Self: root,
	}

	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		spew.Dump(token)

		switch curr_tok := token.(type) {

		case xml.StartElement:
			element = &Element{
				Parent: element.Self,
				Self: &Node{Label: curr_tok.Name.Local},
			}
			for _, attr := range curr_tok.Attr {
				element.AddAttribute(attr)
			}
			element.AddSelfToParentsChildren()
		case xml.CharData:
			element.Data = string(curr_tok)
		case xml.EndElement:
			element = element.Parent
		}
			
	}
	spew.Dump(root)
	return nil
}
