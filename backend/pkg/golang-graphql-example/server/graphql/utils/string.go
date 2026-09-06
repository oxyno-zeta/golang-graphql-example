package utils

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/microcosm-cc/bluemonday"
)

// Build the policy once and reuse it without modification across requests.
var stringPolicy = bluemonday.UGCPolicy()

// UnmarshalString sanitizes GraphQL String inputs before resolver execution.
// StrictPolicy removes HTML and encodes HTML-sensitive text. Do not unescape
// its result: doing so can turn encoded markup back into executable HTML.
func UnmarshalString(v any) (string, error) {
	s, err := graphql.UnmarshalString(v)
	if err != nil {
		return "", err
	}

	return stringPolicy.Sanitize(s), nil
}

// MarshalString preserves response values, including introspection descriptions.
// Clients must still escape values appropriately for their rendering context.
func MarshalString(s string) graphql.Marshaler {
	return graphql.MarshalString(s)
}
