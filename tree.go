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
	Children 	map[string]Elements
	Label 		string
	Attrs		Attributes
	Data 		string
}

type Elements []*Element

func (self *Element) AddAttribute(attr xml.Attr) {
	self.Attrs = append(self.Attrs, &Attribute{Label: attr.Name.Local, Value: SanitiseData(string(attr.Value))})
}

func (self *Element) AddChild(child *Element) {
	if nil == self.Children {
		self.Children = make(map[string]Elements)
	}
	self.Children[child.Label] = append(self.Children[child.Label], child)
}

func (self *Element) AddSelfToParentsChildren() {
	self.Parent.AddChild(self)
}

func (self *Element) IsRoot() bool {
	return self.Parent == nil
}

func (self *Element) IsEmptyField() bool {
	return !self.IsRoot() && !self.HasChild() && !self.HasData() && !self.HasAttrs()
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

func (self *Element) HasLikeSiblings() bool {
	return len(self.Parent.Children[self.Label]) > 1
}
