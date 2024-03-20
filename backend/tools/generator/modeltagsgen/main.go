package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"emperror.dev/errors"
	"github.com/spf13/cobra"
)

type inputData struct {
	PkgPath      string
	PkgName      string
	StructName   string
	OutputSuffix string
}

func main() {
	var outputSuffix string

	var rootCmd = &cobra.Command{
		Use:   "modeltagsgen [package] [model]",
		Short: "used to generate constants from tag models",
		Args:  cobra.MinimumNArgs(2), //nolint:gomnd // No const
		Run: func(_ *cobra.Command, args []string) {
			// Put args into structure
			input := &inputData{
				PkgPath:      args[0],
				StructName:   args[1],
				OutputSuffix: outputSuffix,
			}
			// Validate a bit
			if !strings.Contains(input.PkgPath, "/") {
				hardExit(errors.New("pkg name must be place as arg 0"))
			}

			// Compute pkg name
			split := strings.Split(input.PkgPath, "/")
			// Save
			input.PkgName = split[len(split)-1]

			// Generate
			buff, err := generate(input)
			if err != nil {
				hardExit(err)
			}

			fmt.Println(buff.String())
		},
	}

	rootCmd.PersistentFlags().StringVarP(&outputSuffix, "output-suffix", "o", "modeltags_generated", "model tags generated file name")

	if err := rootCmd.Execute(); err != nil {
		hardExit(err)
	}
}

func hardExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func generate(input *inputData) (*bytes.Buffer, error) {
	// Init template
	tmpl, err := template.New("test").Parse(tmplStr)
	// Check error
	if err != nil {
		return nil, err
	}

	// Create buffer
	buf := &bytes.Buffer{}

	// Execute template
	err = tmpl.Execute(buf, input)
	// Check error
	if err != nil {
		return nil, err
	}

	return buf, nil
}
