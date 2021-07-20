package authorization

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authentication"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
	goerrors "github.com/pkg/errors"
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
	// Check if configuration is empty
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
		return false, goerrors.WithStack(err)
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

func (s *service) requestOPAServer(ctx context.Context, opaCfg *config.OPAServerAuthorization, body []byte) (bool, error) {
	// Get trace from context
	trace := tracing.GetTraceFromContext(ctx)
	// Generate child trace
	childTrace := trace.GetChildTrace("opa-server.request")
	defer childTrace.Finish()
	// Add data
	childTrace.SetTag("opa.uri", opaCfg.URL)

	// Change NewRequest to NewRequestWithContext and pass context it
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, opaCfg.URL, bytes.NewBuffer(body))
	// Check error
	if err != nil {
		return false, goerrors.WithStack(err)
	}
	// Add content type
	req.Header.Add("Content-Type", "application/json")
	// Forward trace on request
	err = childTrace.InjectInHTTPHeader(req.Header)
	// Check error
	if err != nil {
		return false, err
	}
	// Making request to OPA server
	resp, err := http.DefaultClient.Do(req)
	// Check error
	if err != nil {
		return false, goerrors.WithStack(err)
	}
	// Defer closing body
	defer resp.Body.Close()

	// Prepare answer
	var answer opaAnswer
	// Decode answer
	err = json.NewDecoder(resp.Body).Decode(&answer)
	// Check error
	if err != nil {
		return false, goerrors.WithStack(err)
	}

	return answer.Result, nil
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
