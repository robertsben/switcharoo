package main

import (
	"io"
	"github.com/davecgh/go-spew/spew"
)

type Encoder struct {
	w 	io.Writer
	err	error
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (enc *Encoder) write(s string) {
	enc.w.Write([]byte(s))
}

func (enc *Encoder) Encode(root *Element) error {
	if enc.err != nil {
		return enc.err
	}

	if root == nil {
		return nil
	}

	spew.Dump(root)

	enc.generateJsonFromElement(root)

	spew.Dump(enc)

	return nil

}

func (enc *Encoder) generateJsonFromElement(elem *Element) {

	if elem.HasAttrsAndChild() {
		enc.generateAttrsAndChildrenJson(elem)
	} else if elem.HasAttrsAndData() {
		enc.generateAttrsAndDataJson(elem)
	} else if elem.HasAttrs() {
		if !elem.HasLikeSiblings() {
			enc.generateLabelJson(elem.Label)
		}
		enc.write("{\n")
		enc.generateAttrsJson(elem.Attrs)
		enc.write("\n}")
	} else if elem.HasData() {
		if !elem.HasLikeSiblings() {
			enc.generateLabelJson(elem.Label)
		}
		enc.generateDataJson(elem.Data)
	} else if elem.HasChild() {
		if !elem.IsRoot() {
			enc.generateLabelJson(elem.Label)
		}
		enc.write("{\n")
		enc.generateChildrenJson(elem)
		enc.write("\n}")
	}
}

func (enc *Encoder) generateChildrenJson(elem *Element) {
	for index, siblings := range elem.Children {
		if len(siblings) > 1 {
			enc.generateLabelJson(index)
			enc.generateLikeSiblingsJson(siblings)	
		} else {
			enc.generateJsonFromElement(siblings[0])
		}
		
	}
}

func (enc *Encoder) generateLikeSiblingsJson(siblings Elements) {
	enc.write("[\n")
	for index, sibling := range siblings {
		enc.generateJsonFromElement(sibling)
		if index != (len(siblings)-1) {
			spew.Dump(index)
			spew.Dump(len(siblings))
			enc.appendComma()
		}
	}
	enc.write("]\n")
}

func (enc *Encoder) generateLabelJson(label string) {
	enc.write("\"")
	enc.write(label)
	enc.write("\": ")
}

func (enc *Encoder) generateDataJson(data string) {
	enc.write("\"")
	enc.write(data)
	enc.write("\"")
}

func (enc *Encoder) appendComma() {
	enc.write(",\n")
}

func (enc *Encoder) generateAttrsJson(attrs Attributes) {
	for index, attr := range attrs {
		enc.generateLabelJson("-" + attr.Label)
		enc.generateDataJson(attr.Value)
		if index != len(attrs)-1 {
			enc.appendComma()
		}
	}
}

func (enc *Encoder) generateAttrsAndChildrenJson(elem *Element) {
	if !elem.HasLikeSiblings() {
		enc.generateLabelJson(elem.Label)
	}
	enc.write("{\n")
	enc.generateAttrsJson(elem.Attrs)
	enc.appendComma()
	enc.generateChildrenJson(elem)
	enc.write("\n}")
}

func (enc *Encoder) generateAttrsAndDataJson(elem *Element) {
	if !elem.HasLikeSiblings() {
		enc.generateLabelJson(elem.Label)
	}
	enc.write("{\n")
	enc.generateLabelJson("#text")
	enc.generateDataJson(elem.Data)
	enc.appendComma()
	enc.generateAttrsJson(elem.Attrs)
	enc.write("\n}")
}
