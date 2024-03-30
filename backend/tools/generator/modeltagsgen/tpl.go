package main

const tmplStr = `
package main

import (
	"fmt"
	"os"

	"github.com/oxyno-zeta/golang-graphql-example/tools/generator/modeltagsgen/gen"

	_pkg "{{ .PkgPath }}"
)

var obj interface{} = _pkg.{{ .StructName }}{}
var pkgName = "{{ .PkgName }}"
var outputStdout = false
var outputFileName = "{{ .LowercaseStructName }}_{{ .OutputSuffix }}.go"

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
`
