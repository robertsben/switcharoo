package main

import (
	"bytes"
)

func Convert(input string) (*bytes.Buffer, error) {
	root := &Element{}
	NewDecoder(bytes.NewBufferString(input)).Decode(root)

	output := new(bytes.Buffer)
	NewEncoder(output).Encode(root)

	if !IsValidJSON(string(output.Bytes())) {
		return output, DataError{Err: "Invalid JSON", Data: output.Bytes()}
	}

	return output, nil
}