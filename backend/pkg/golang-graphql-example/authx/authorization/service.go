package authorization

import (
	"context"
	"encoding/json"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type service struct {
	cfgManager config.Manager
}

type generalInputOPA struct {
	Input *generalInputDataOPA `json:"input"`
}

type generalInputDataOPA struct {
	User *models.OIDCUser  `json:"user"`
	Tags map[string]string `json:"tags"`
	Data *generalDataOPA   `json:"data"`
}

type generalDataOPA struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

type opaAnswer struct {
	Result bool `json:"result"`
}

func (s *service) IsAuthorized(ctx context.Context, action, resource string) (bool, error) {
	// Get logger
	logger := log.GetLoggerFromContext(ctx)
	// Get configuration to check that authorization can be calculated
	cfg := s.cfgManager.GetConfig().OPAServerAuthorization
	if cfg == nil {
		// Configuration doesn't exists, authorization is given
		return true, nil
	}

	// Get user from context
	user := authentication.GetAuthenticatedUserFromContext(ctx)

	// Create opa input
	input := &generalInputOPA{
		Input: &generalInputDataOPA{
			User: user,
			Tags: cfg.Tags,
			Data: &generalDataOPA{
				Action:   action,
				Resource: resource,
			},
		},
	}
	// Json encode body
	bb, err := json.Marshal(input)
	if err != nil {
		return false, err
	}

	authorized, err := s.requestOPAServer(ctx, cfg, bb)
	// Check error
	if err != nil {
		return false, err
	}

	// Check if user isn't authorized
	if !authorized {
		logger.Infof("User %s not authorized for action %s on resource %s", user.GetIdentifier(), action, resource)

		return false, nil
	}

	logger.Infof("User %s authorized for action %s on resource %s", user.GetIdentifier(), action, resource)

	return true, nil
}

func (s *service) CheckAuthorized(ctx context.Context, action, resource string) error {
	// Call is authorized
	res, err := s.IsAuthorized(ctx, action, resource)
	// Check error
	if err != nil {
		return err
	}

	// Check not authorized
	if !res {
		return errors.NewForbiddenError("forbidden")
	}

	return nil
}
