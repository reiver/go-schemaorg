# go-schemaorg

Package **schemaorg** provides an implementation of the [schema.org](http://schema.org/) vocabulary, as they are used on the Fediverse, for the Go programming language.

Use https://github.com/reiver/go-jsonld to marshal types in this package.

## Example

```golang
import (
	"github.com/reiver/go-activitystreams"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-schemaorg"
)

// ...

var person = activitystreams.Person{

	// ...

	Attachments: []any{
		schemaorg.PropertyValue{
			Name:  "Location",
			Value: "Metro Vancouver",
		},
		schemaorg.PropertyValue{
			Name:  "Home-Page",
			Value: "http://example.com/~joeblow",
		},
	},

	// ...
}.

// ...

bytes, err := jsonld.Marshal()
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-schemaorg

[![GoDoc](https://godoc.org/github.com/reiver/go-schemaorg?status.svg)](https://godoc.org/github.com/reiver/go-schemaorg)

## Import

To import package **schemaorg** use `import` code like the follownig:
```
import "github.com/reiver/go-schemaorg"
```

## Installation

To install package **schemaorg** do the following:
```
GOPROXY=direct go get github.com/reiver/go-schemaorg
```

## Author

Package **schemaorg** was written by [Charles Iliya Krempeaux](http://reiver.link)
