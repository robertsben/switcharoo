package main

import (
	"encoding/xml"
)

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

func (self *Element) HasChild() bool {
	return len(self.Children) > 0
}

func (self *Element) HasData() bool {
	return len(self.Data) > 0
}

func (self *Element) HasAttrs() bool {
	return len(self.Attrs) > 0
}

func (self *Element) HasAttrsAndData() bool {
	return self.HasAttrs() && self.HasData()
}

func (self *Element) HasAttrsAndChild() bool {
	return self.HasAttrs() && self.HasChild()
}