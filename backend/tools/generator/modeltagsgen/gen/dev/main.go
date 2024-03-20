package main

import (
	"fmt"
	"os"

	"github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen"
)

var obj interface{} = Dev1{}
var pkgName = "main"
var outputStdout = false
var outputFileName = "Dev1_modeltags.go"

func main() {
	buf, err := gen.Generate(pkgName, obj)
	// Check error
	if err != nil {
		panic(err)
	}

	if outputStdout {
		fmt.Println(buf.String())
	} else {
		err = os.WriteFile(outputFileName, buf.Bytes(), 0600)
		// Check error
		if err != nil {
			panic(err)
		}
	}
}
