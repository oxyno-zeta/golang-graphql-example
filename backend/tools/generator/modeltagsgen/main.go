package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"emperror.dev/errors"
	"github.com/spf13/cobra"
)

type inputData struct {
	PkgPath              string
	PkgName              string
	StructName           string
	LowercaseStructName  string
	OutputSuffix         string
	DisableJSON          bool
	DisableGorm          bool
	DisableStructKeyName bool
}

func main() {
	var outputSuffix string

	var disableJSON, disableGorm, disableStructKeyName bool

	var rootCmd = &cobra.Command{
		Use:   "modeltagsgen [package] [model]",
		Short: "used to generate constants from tag models",
		Args:  cobra.MinimumNArgs(2), //nolint:mnd // No const
		Run: func(_ *cobra.Command, args []string) {
			// Put args into structure
			input := &inputData{
				PkgPath:              args[0],
				StructName:           args[1],
				LowercaseStructName:  strings.ToLower(args[1]),
				OutputSuffix:         outputSuffix,
				DisableJSON:          disableJSON,
				DisableGorm:          disableGorm,
				DisableStructKeyName: disableStructKeyName,
			}
			// Validate a bit
			if !strings.Contains(input.PkgPath, "/") {
				hardExit(errors.New("pkg name must be place as arg 0"))
			}

			// Compute pkg name
			split := strings.Split(input.PkgPath, "/")
			// Save
			input.PkgName = split[len(split)-1]

			// Run
			err := run(input)
			// Check error
			if err != nil {
				hardExit(err)
			}
		},
	}

	rootCmd.PersistentFlags().StringVarP(&outputSuffix, "output-suffix", "o", "modeltags_generated", "model tags generated file name")
	rootCmd.PersistentFlags().BoolVarP(&disableJSON, "disable-json", "j", false, "disable json")
	rootCmd.PersistentFlags().BoolVarP(&disableGorm, "disable-gorm", "g", false, "disable gorm")
	rootCmd.PersistentFlags().BoolVarP(&disableStructKeyName, "disable-struct-key-name", "s", false, "disable struct key name")

	if err := rootCmd.Execute(); err != nil {
		hardExit(err)
	}
}

func hardExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func run(input *inputData) error {
	// Generate
	buff, err := generate(input)
	if err != nil {
		return err
	}

	err = os.MkdirAll("_prog", os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}

	filePath := "./_prog/prog.go"

	err = os.WriteFile(filePath, buff.Bytes(), 0600)
	if err != nil {
		return errors.WithStack(err)
	}

	cmd := exec.Command("go", "run", input.PkgPath+"/_prog")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err = cmd.Run()
	if err != nil {
		return errors.New(errb.String())
	}

	// Remove all
	err = os.RemoveAll("_prog")
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
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
