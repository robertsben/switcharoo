## Switcheroo

# Concept

A package to convert data notation types.

# Usage

Initialise from the command line:
```
go install github.com/robertsben/switcheroo

./switcheroo -source {relative/path/to/source-file.xml} -destination {relative/path/to/destination-file.json} -debug
```

Debug is optional. Source and destination are also optional, if you're happy to use the... somewhat shoddy... web app (on localhost:8080)[http://localhost:8080].

# @TODO

* change the structs in `tree.go`?
* improve the performance somehow perhaps - make stuff concurrent?
* XML to YML
* JSON to XML&YML
* YML to XML&JSON
* make this whole package include-able