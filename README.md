# Go Embed Generator

## Summary

`goembed` generates a file that you may include in your project to include a file's content as code, embedded within the
binary.

## Installation

You can use `go` tool to install it to you Go home (`$HOME/go/bin`) as `goembed` binary.

```shell
go install github.com/rhaseven7h/goembed@latest
```

## Usage

You must specify the package name for the generated file, the output path and/or file name, the variable name for the
generated data, and finally, the input file.

```shell
goembed \ 
   -package mypackagename \
   -output embeddedfiles/dataone.go \
   -variable MyData \
   my_input_file.txt
```

You may specify `-` for both input and out files, which will use standard input and standard output respectively.

## Output

An example output for the above command would be the following:

```go
package data

import (
	"encoding/base64"
)

var InitData []byte

func init() {
	InitData, _ = base64.StdEncoding.DecodeString(
		"" +
			"ZnVuY3Rpb24gZ3ZtKCkgewogIE9VVFBVVD0kKHNndm0gIiRAIikKICBFWElUX0NP" +
			"REU9JD8KICBpZiBbICRFWElUX0NPREUgLW5lIDEgXTsgdGhlbgogICAgZWNobyAk" +
			"T1VUUFVUCiAgICByZXR1cm4KICBmaQogIGV2YWwgIiR7T1VUUFVUfSIKfQo=" +
			"",
	)
}
```

With this, you may import the `data` package, and use the `data.InitData`
variable which will contain the original input contents in a `[]byte` slice.

## Go Generate

You may use this tool with Go Generate functionality by adding a line as follows to your source code:

```go
//go:generate goembed -package pkgname -output data/input.go -variable MyData my_input_file.txt
```

## Copyright

This work is licensed under
a [Creative Commons Attribution-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-sa/4.0/).

![ Creative Commons Attribution-ShareAlike 4.0 International License.](https://i.creativecommons.org/l/by-sa/4.0/88x31.png)
