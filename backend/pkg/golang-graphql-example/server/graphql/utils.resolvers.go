package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
)

// Eq is the resolver for the eq field.
func (r *booleanFilterResolver) Eq(ctx context.Context, obj *common.GenericFilter, data *bool) error {
	obj.Eq = data

	return nil
}

// NotEq is the resolver for the notEq field.
func (r *booleanFilterResolver) NotEq(ctx context.Context, obj *common.GenericFilter, data *bool) error {
	obj.NotEq = data

	return nil
}

// Eq is the resolver for the eq field.
func (r *intFilterResolver) Eq(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Eq = data

	return nil
}

// NotEq is the resolver for the notEq field.
func (r *intFilterResolver) NotEq(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotEq = data

	return nil
}

// Gte is the resolver for the gte field.
func (r *intFilterResolver) Gte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Gte = data

	return nil
}

// NotGte is the resolver for the notGte field.
func (r *intFilterResolver) NotGte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotGte = data

	return nil
}

// Gt is the resolver for the gt field.
func (r *intFilterResolver) Gt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Gt = data

	return nil
}

// NotGt is the resolver for the notGt field.
func (r *intFilterResolver) NotGt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotGt = data

	return nil
}

// Lte is the resolver for the lte field.
func (r *intFilterResolver) Lte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Lte = data

	return nil
}

// NotLte is the resolver for the notLte field.
func (r *intFilterResolver) NotLte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Lte = data

	return nil
}

// Lt is the resolver for the lt field.
func (r *intFilterResolver) Lt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Lt = data

	return nil
}

// NotLt is the resolver for the notLt field.
func (r *intFilterResolver) NotLt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotLt = data

	return nil
}

// In is the resolver for the in field.
func (r *intFilterResolver) In(ctx context.Context, obj *common.GenericFilter, data []*int) error {
	obj.In = data

	return nil
}

// NotIn is the resolver for the notIn field.
func (r *intFilterResolver) NotIn(ctx context.Context, obj *common.GenericFilter, data []*int) error {
	obj.NotIn = data

	return nil
}

// Eq is the resolver for the eq field.
func (r *stringFilterResolver) Eq(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.Eq = data

	return nil
}

// NotEq is the resolver for the notEq field.
func (r *stringFilterResolver) NotEq(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotEq = data

	return nil
}

// Contains is the resolver for the contains field.
func (r *stringFilterResolver) Contains(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.Contains = data

	return nil
}

// NotContains is the resolver for the notContains field.
func (r *stringFilterResolver) NotContains(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotContains = data

	return nil
}

// StartsWith is the resolver for the startsWith field.
func (r *stringFilterResolver) StartsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.StartsWith = data

	return nil
}

// NotStartsWith is the resolver for the notStartsWith field.
func (r *stringFilterResolver) NotStartsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotStartsWith = data

	return nil
}

// EndsWith is the resolver for the endsWith field.
func (r *stringFilterResolver) EndsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.EndsWith = data

	return nil
}

// NotEndsWith is the resolver for the notEndsWith field.
func (r *stringFilterResolver) NotEndsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotEndsWith = data

	return nil
}

// In is the resolver for the in field.
func (r *stringFilterResolver) In(ctx context.Context, obj *common.GenericFilter, data []*string) error {
	obj.In = data

	return nil
}

// NotIn is the resolver for the notIn field.
func (r *stringFilterResolver) NotIn(ctx context.Context, obj *common.GenericFilter, data []*string) error {
	obj.NotIn = data

	return nil
}

// BooleanFilter returns generated.BooleanFilterResolver implementation.
func (r *Resolver) BooleanFilter() generated.BooleanFilterResolver { return &booleanFilterResolver{r} }

// IntFilter returns generated.IntFilterResolver implementation.
func (r *Resolver) IntFilter() generated.IntFilterResolver { return &intFilterResolver{r} }

// StringFilter returns generated.StringFilterResolver implementation.
func (r *Resolver) StringFilter() generated.StringFilterResolver { return &stringFilterResolver{r} }

type booleanFilterResolver struct{ *Resolver }
type intFilterResolver struct{ *Resolver }
type stringFilterResolver struct{ *Resolver }
