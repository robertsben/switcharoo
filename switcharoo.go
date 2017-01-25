package main

import (
	"flag"
	"io/ioutil"
	"github.com/davecgh/go-spew/spew"
)

type Page struct {
	Title string
	Body []byte
}

type Data struct {
	Conversion []byte
}

var Debug bool
var DumpFail bool

func conversionHandler(inputfilename string, outputfilename string) {
	input, _ := ioutil.ReadFile(inputfilename)
	output, err := Convert(string(input))

	if err != nil {
		if Debug {
			spew.Dump(err)
		}
		if DumpFail {
			ioutil.WriteFile("./failure.json", output.Bytes(), 0644)	
		}
		data_out := []byte(err.Error())
	} else {
		data_out := output.Bytes()
	}
	ioutil.WriteFile(outputfilename, data_out, 0644)
}

func main() {
	var source string
	var destination string
	flag.StringVar(&source, "source", "", "path to the source xml file")
	flag.StringVar(&destination, "destination", "./example.json", "path to the output json file")
	flag.BoolVar(&Debug, "debug", false, "whether to add debugging logs")
	flag.BoolVar(&DumpFail, "dumpfail", true, "whether to write failed conversion to failure.json")
	flag.Parse()

	conversionHandler(source, destination)
}