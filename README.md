# Go Embed Generator

## Summary

`goembed` generates a file that you may include in your project to include a file's content as code, embedded within the
binary.

## Usage

You must specify the package name for the generated file, the output path and/or file name, the variable name for the
generated data, and finally, the input file.

```shell
goembed \ 
   -package mypackagename \
   -output embeddedfiles/dataone.go \
   -variable myData \
   my_input_file.txt
```

You may specify `-` for both input and out files, which will use standard input and standard output respectively.

## Output

An example output for the above command would be the following:

```text
package testo

import (
	"encoding/base64"
)

testoVar, _ := base64.StdEncoding.DecodeString("" +
	"cGFja2FnZSBtYWluCgppbXBvcnQgKAoJImVuY29kaW5nL2Jhc2U2NCIKCSJmbGFn" +
	"IgoJImZtdCIKCSJpby9pb3V0aWwiCgkib3MiCikKCmNvbnN0ICgKCW1pc3NpbmdQ" +
	.
	.
	.
	
	"djogJXNcbiIsIG91dHB1dExpbmUsIGVyci5FcnJvcigpKQoJCQlvcy5FeGl0KDEp" +
	"CgkJfQoJfQp9Cg==" +
	""
)
```

With this, you may directly use the `testoVar` variable which will contain the original input contents in a `[]byte`
slice.

## Go Generate

You may use this tool with Go Generate functionality by adding a line as follows to your source code:

```go
//go:generate goembed -package pkgname -output data/input.go -variable myData my_input_file.txt
```

## Copyright

This work is licensed under
a [Creative Commons Attribution-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-sa/4.0/).

![ Creative Commons Attribution-ShareAlike 4.0 International License.](https://i.creativecommons.org/l/by-sa/4.0/88x31.png)