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
		enc.generateAttrsAndDataJson(elem.Attrs, elem.Data)
	} else if elem.HasAttrs() {
		enc.generateAttrsJson(elem.Attrs)
	} else if elem.HasData() {
		enc.generateDataJson(elem.Data)
	}

	if elem.HasChild() {
		enc.generateChildrenJson(elem)
	}
}

func (enc *Encoder) generateChildrenJson(elem *Element) {
	for index, elem := range elem.Children {
		enc.generateLabelJson(elem.Label)
		enc.generateJsonFromElement(elem)
		if index != len(elem.Children)-1 {
			enc.appendComma()
		}
	}
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
	enc.generateLabelJson(elem.Label)
	enc.write("{\n")
	enc.generateAttrsJson(elem.Attrs)
	enc.appendComma()
	enc.generateChildrenJson(elem)
	enc.write("\n}")
}

func (enc *Encoder) generateAttrsAndDataJson(attrs Attributes, data string) {
	enc.generateLabelJson("#text")
	enc.generateDataJson(data)
	enc.appendComma()
	enc.generateAttrsJson(attrs)
}
