package main

import (
	"encoding/json"
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

func IsValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func (enc *Encoder) Encode(root *Element) error {
	if enc.err != nil {
		return enc.err
	}

	if root == nil {
		return nil
	}

	if Debug {
		spew.Dump(root)
	}

	enc.generateJsonFromElement(root)

	if Debug {
		spew.Dump(enc)
	}

	return nil

}

func (enc *Encoder) generateJsonFromElement(elem *Element) {
	switch {

	case elem.HasAttrsAndChild():
		enc.generateAttrsAndChildrenJson(elem)

	case elem.HasAttrsAndData():
		enc.generateAttrsAndDataJson(elem)

	case elem.HasAttrs():
		enc.generateWrappedAttrsJson(elem)

	case elem.HasData():
		enc.generateLabelJsonIfWithoutLikeSiblings(elem)
		enc.generateDataJson(elem.Data)

	case elem.IsEmptyField():
		enc.generateLabelJsonIfWithoutLikeSiblings(elem)
		enc.write("null")

	case elem.HasChild():
		enc.generateWrappedChildrenJson(elem)

	}

}

func (enc *Encoder) generateAttrsAndChildrenJson(elem *Element) {
	enc.startJsonObjectWithLabelIfNeeded(elem)
	enc.generateAttrsJson(elem.Attrs)
	enc.appendComma()
	enc.generateChildrenJson(elem)
	enc.endJsonObject()
}

func (enc *Encoder) generateAttrsAndDataJson(elem *Element) {
	enc.startJsonObjectWithLabelIfNeeded(elem)
	enc.generateLabelJson("#text")
	enc.generateDataJson(elem.Data)
	enc.appendComma()
	enc.generateAttrsJson(elem.Attrs)
	enc.endJsonObject()
}

func (enc *Encoder) generateWrappedAttrsJson(elem *Element) {
	enc.startJsonObjectWithLabelIfNeeded(elem)
	enc.generateAttrsJson(elem.Attrs)
	enc.endJsonObject()
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

func (enc *Encoder) generateWrappedChildrenJson(elem *Element) {
	if !elem.IsRoot() {
		enc.generateLabelJsonIfWithoutLikeSiblings(elem)
	}
	enc.startJsonObject()
	enc.generateChildrenJson(elem)
	enc.endJsonObject()
}

func (enc *Encoder) generateChildrenJson(elem *Element) {
	count := 0
	for index, siblings := range elem.Children {
		if len(siblings) > 1 {
			enc.generateLabelJson(index)
			enc.generateLikeSiblingsJson(siblings)	
		} else {
			enc.generateJsonFromElement(siblings[0])
		}
		if count != (len(elem.Children)-1) {
			enc.appendComma()
		}
		count++
	}
}

func (enc *Encoder) generateLikeSiblingsJson(siblings Elements) {
	enc.write("[\n")
	for index, sibling := range siblings {
		enc.generateJsonFromElement(sibling)
		if index != (len(siblings)-1) {
			enc.appendComma()
		}
	}
	enc.write("]\n")
}

func (enc *Encoder) startJsonObjectWithLabelIfNeeded(elem *Element) {
	enc.generateLabelJsonIfWithoutLikeSiblings(elem)
	enc.startJsonObject()
}

func (enc *Encoder) generateLabelJsonIfWithoutLikeSiblings(elem *Element) {
	if !elem.HasLikeSiblings() {
		enc.generateLabelJson(elem.Label)
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

func (enc *Encoder) startJsonObject() {
	enc.write("{\n")
}

func (enc *Encoder) endJsonObject() {
	enc.write("\n}")
}

func (enc *Encoder) appendComma() {
	enc.write(",\n")
}
