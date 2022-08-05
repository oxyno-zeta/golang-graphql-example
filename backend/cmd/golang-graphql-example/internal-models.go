package main

import "strings"

type targetDefinition struct {
	// Run target function
	// This function must crash the app if necessary
	Run func(sv *services)
	// Is it a primary target to perform before any other not primary ?
	Primary bool
	// Should be considered in "all" target ?
	InAllTarget bool
	// That flag will declare that this target will need that main must block the end
	// Example case: server listening.
	BlockMain bool
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "string list representation"
}

func (i *arrayFlags) Set(value string) error {
	// Try to split on "," to allow --OPTION=foo,bar
	values := strings.Split(value, ",")

	*i = append(*i, values...)

	return nil
}
