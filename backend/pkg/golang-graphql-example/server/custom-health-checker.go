package server

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type customHealthChecker struct {
	logger log.Logger
	fn     func() error
	name   string
}

func (chc *customHealthChecker) Name() string {
	return chc.name
}

func (chc *customHealthChecker) Execute(_ context.Context) (interface{}, error) {
	// Run check
	err := chc.fn()
	// Check error
	if err != nil {
		// Log it and return
		chc.logger.Error(err)

		return nil, err
	}

	// Default
	return nil, nil //nolint: nilnil // not needed here
}
