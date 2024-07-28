package main

import (
	"context"
	"strings"
)

type targetDefinition struct {
	// Run target function
	// This function must crash the app if necessary
	Run func(targets []string, sv *services)
	// Is it a primary target to perform before any other not primary ?
	Primary bool
	// Should be considered in "all" target ?
	InAllTarget bool
}

type daemonDefinition struct {
	// Run target function
	// This function must crash the app if necessary
	Run func(ctx context.Context, targets []string, sv *services)
}

type arrayFlags []string

func (*arrayFlags) String() string {
	return "string list representation"
}

func (i *arrayFlags) Set(value string) error {
	// Try to split on "," to allow --OPTION=foo,bar
	values := strings.Split(value, ",")

	*i = append(*i, values...)

	return nil
}
