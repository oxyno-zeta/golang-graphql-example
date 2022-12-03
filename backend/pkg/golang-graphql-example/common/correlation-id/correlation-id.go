package correlationid

import (
	"context"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

type contextKey struct {
	name string
}

var correlationIDCtxKey = &contextKey{name: "correlation-id"}

func Generate() (string, error) {
	// Generate uuid
	uuid, err := uuid.NewV4()
	// Check error
	if err != nil {
		return "", errors.WithStack(err)
	}

	// Default
	return uuid.String(), nil
}

func GetFromContext(ctx context.Context) string {
	requestIDObj := ctx.Value(correlationIDCtxKey)
	if requestIDObj != nil {
		return requestIDObj.(string) //nolint: forcetypeassert // Ignored
	}

	return ""
}

func SetInContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, correlationIDCtxKey, requestID)
}
