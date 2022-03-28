package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/common"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/generated"
)

func (r *booleanFilterResolver) Eq(ctx context.Context, obj *common.GenericFilter, data *bool) error {
	obj.Eq = data

	return nil
}

func (r *booleanFilterResolver) NotEq(ctx context.Context, obj *common.GenericFilter, data *bool) error {
	obj.NotEq = data

	return nil
}

func (r *intFilterResolver) Eq(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Eq = data

	return nil
}

func (r *intFilterResolver) NotEq(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotEq = data

	return nil
}

func (r *intFilterResolver) Gte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Gte = data

	return nil
}

func (r *intFilterResolver) NotGte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotGte = data

	return nil
}

func (r *intFilterResolver) Gt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Gt = data

	return nil
}

func (r *intFilterResolver) NotGt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotGt = data

	return nil
}

func (r *intFilterResolver) Lte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Lte = data

	return nil
}

func (r *intFilterResolver) NotLte(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Lte = data

	return nil
}

func (r *intFilterResolver) Lt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.Lt = data

	return nil
}

func (r *intFilterResolver) NotLt(ctx context.Context, obj *common.GenericFilter, data *int) error {
	obj.NotLt = data

	return nil
}

func (r *intFilterResolver) In(ctx context.Context, obj *common.GenericFilter, data []*int) error {
	obj.In = data

	return nil
}

func (r *intFilterResolver) NotIn(ctx context.Context, obj *common.GenericFilter, data []*int) error {
	obj.NotIn = data

	return nil
}

func (r *stringFilterResolver) Eq(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.Eq = data

	return nil
}

func (r *stringFilterResolver) NotEq(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotEq = data

	return nil
}

func (r *stringFilterResolver) Contains(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.Contains = data

	return nil
}

func (r *stringFilterResolver) NotContains(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotContains = data

	return nil
}

func (r *stringFilterResolver) StartsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.StartsWith = data

	return nil
}

func (r *stringFilterResolver) NotStartsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotStartsWith = data

	return nil
}

func (r *stringFilterResolver) EndsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.EndsWith = data

	return nil
}

func (r *stringFilterResolver) NotEndsWith(ctx context.Context, obj *common.GenericFilter, data *string) error {
	obj.NotEndsWith = data

	return nil
}

func (r *stringFilterResolver) In(ctx context.Context, obj *common.GenericFilter, data []*string) error {
	obj.In = data

	return nil
}

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
