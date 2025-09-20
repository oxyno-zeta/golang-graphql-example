package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "graphqlgen",
		Short: "used to generate daos from configuration",
		Run: func(_ *cobra.Command, args []string) {
			// Run
			err := run()
			// Check error
			if err != nil {
				hardExit(err)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		hardExit(err)
	}
}

func hardExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func run() error {
	// Load configuration
	cfg, err := loadConfig()
	// Check error
	if err != nil {
		return err
	}

	// Generate
	err = generate(cfg)
	// Check error
	if err != nil {
		return err
	}

	return nil
}
