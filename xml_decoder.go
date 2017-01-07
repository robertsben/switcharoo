package main

import (
	"encoding/xml"
	"io"
	// "fmt"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

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
