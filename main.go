package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	missingParameter     = "missing-parameter"
	outputLinesMaxLength = 64
)

func main() {
	var err error
	var packageName string
	var variableName string
	var outputFileName string
	flag.StringVar(&packageName,
		"package",
		missingParameter,
		"package name to use on generated files",
	)
	flag.StringVar(&variableName,
		"variable",
		missingParameter,
		"output variable name",
	)
	flag.StringVar(&outputFileName,
		"output",
		missingParameter,
		"output file name",
	)
	flag.Parse()
	inputFileName := flag.Arg(0)
	if packageName == "" || packageName == missingParameter {
		fmt.Println("invalid or missing package name")
		os.Exit(1)
	}
	if variableName == "" || variableName == missingParameter {
		fmt.Println("invalid or missing constant name")
		os.Exit(1)
	}
	if outputFileName == "" || outputFileName == missingParameter {
		fmt.Println("invalid or missing output file name")
		os.Exit(1)
	}

	var inputContentBytes []byte
	if inputFileName == "-" {
		inputContentBytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		inputContentBytes, err = ioutil.ReadFile(inputFileName)
	}
	if err != nil {
		fmt.Printf("invalid or missing input file or file name: %s\n", err.Error())
		os.Exit(1)
	}
	encodedInput := base64.StdEncoding.EncodeToString(inputContentBytes)

	var outputFile *os.File
	if outputFileName == "-" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.OpenFile(outputFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		defer func() {
			_ = outputFile.Close()
		}()
		if err != nil {
			fmt.Printf("unable to open output file: %s\n", err.Error())
			os.Exit(1)
		}
	}

	outputDataLines := []string{
		fmt.Sprintf("package %s", packageName),
		"",
		"import (",
		fmt.Sprintf("\t%#v", "encoding/base64"),
		")",
		"",
		fmt.Sprintf("%s, _ := base64.StdEncoding.DecodeString(%#v +", variableName, ""),
	}
	for len(encodedInput) > 0 {
		lineLength := outputLinesMaxLength
		if len(encodedInput) < outputLinesMaxLength {
			lineLength = len(encodedInput)
		}
		outputLine := fmt.Sprintf("\t%#v +", encodedInput[0:lineLength])
		outputDataLines = append(outputDataLines, outputLine)
		encodedInput = encodedInput[lineLength:]
	}
	outputDataLines = append(outputDataLines, []string{
		fmt.Sprintf("\t%#v", ""),
		")",
	}...)
	for _, outputLine := range outputDataLines {
		_, err := fmt.Fprintln(outputFile, outputLine)
		if err != nil {
			fmt.Printf("error writing to output file line %#v: %s\n", outputLine, err.Error())
			os.Exit(1)
		}
	}
}
