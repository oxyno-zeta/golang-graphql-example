// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/model"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/server/graphql/utils"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type TodoResolver interface {
	ID(ctx context.Context, obj *models.Todo) (string, error)
	CreatedAt(ctx context.Context, obj *models.Todo, format *utils.DateFormat) (string, error)
	UpdatedAt(ctx context.Context, obj *models.Todo, format *utils.DateFormat) (string, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Todo_createdAt_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	arg0, err := ec.field_Todo_createdAt_argsFormat(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["format"] = arg0
	return args, nil
}
func (ec *executionContext) field_Todo_createdAt_argsFormat(
	ctx context.Context,
	rawArgs map[string]interface{},
) (*utils.DateFormat, error) {
	// We won't call the directive if the argument is null.
	// Set call_argument_directives_with_null to true to call directives
	// even if the argument is null.
	_, ok := rawArgs["format"]
	if !ok {
		var zeroVal *utils.DateFormat
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("format"))
	if tmp, ok := rawArgs["format"]; ok {
		return ec.unmarshalODateFormat2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋutilsᚐDateFormat(ctx, tmp)
	}

	var zeroVal *utils.DateFormat
	return zeroVal, nil
}

func (ec *executionContext) field_Todo_updatedAt_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	arg0, err := ec.field_Todo_updatedAt_argsFormat(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["format"] = arg0
	return args, nil
}
func (ec *executionContext) field_Todo_updatedAt_argsFormat(
	ctx context.Context,
	rawArgs map[string]interface{},
) (*utils.DateFormat, error) {
	// We won't call the directive if the argument is null.
	// Set call_argument_directives_with_null to true to call directives
	// even if the argument is null.
	_, ok := rawArgs["format"]
	if !ok {
		var zeroVal *utils.DateFormat
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("format"))
	if tmp, ok := rawArgs["format"]; ok {
		return ec.unmarshalODateFormat2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋutilsᚐDateFormat(ctx, tmp)
	}

	var zeroVal *utils.DateFormat
	return zeroVal, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Todo_id(ctx context.Context, field graphql.CollectedField, obj *models.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Todo().ID(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Todo_createdAt(ctx context.Context, field graphql.CollectedField, obj *models.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_createdAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Todo().CreatedAt(rctx, obj, fc.Args["format"].(*utils.DateFormat))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_createdAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Todo_createdAt_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Todo_updatedAt(ctx context.Context, field graphql.CollectedField, obj *models.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_updatedAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Todo().UpdatedAt(rctx, obj, fc.Args["format"].(*utils.DateFormat))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_updatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Todo_updatedAt_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Todo_text(ctx context.Context, field graphql.CollectedField, obj *models.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_text(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Text, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_text(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Todo_done(ctx context.Context, field graphql.CollectedField, obj *models.Todo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Todo_done(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Done, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Todo_done(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Todo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TodoConnection_edges(ctx context.Context, field graphql.CollectedField, obj *model.TodoConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TodoConnection_edges(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Edges, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.TodoEdge)
	fc.Result = res
	return ec.marshalOTodoEdge2ᚕᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋmodelᚐTodoEdge(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TodoConnection_edges(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TodoConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "cursor":
				return ec.fieldContext_TodoEdge_cursor(ctx, field)
			case "node":
				return ec.fieldContext_TodoEdge_node(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type TodoEdge", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _TodoConnection_pageInfo(ctx context.Context, field graphql.CollectedField, obj *model.TodoConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TodoConnection_pageInfo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*utils.PageInfo)
	fc.Result = res
	return ec.marshalNPageInfo2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋutilsᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TodoConnection_pageInfo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TodoConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "hasNextPage":
				return ec.fieldContext_PageInfo_hasNextPage(ctx, field)
			case "hasPreviousPage":
				return ec.fieldContext_PageInfo_hasPreviousPage(ctx, field)
			case "startCursor":
				return ec.fieldContext_PageInfo_startCursor(ctx, field)
			case "endCursor":
				return ec.fieldContext_PageInfo_endCursor(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PageInfo", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _TodoEdge_cursor(ctx context.Context, field graphql.CollectedField, obj *model.TodoEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TodoEdge_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Cursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TodoEdge_cursor(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TodoEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TodoEdge_node(ctx context.Context, field graphql.CollectedField, obj *model.TodoEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TodoEdge_node(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Node, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*models.Todo)
	fc.Result = res
	return ec.marshalOTodo2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TodoEdge_node(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TodoEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Todo_id(ctx, field)
			case "createdAt":
				return ec.fieldContext_Todo_createdAt(ctx, field)
			case "updatedAt":
				return ec.fieldContext_Todo_updatedAt(ctx, field)
			case "text":
				return ec.fieldContext_Todo_text(ctx, field)
			case "done":
				return ec.fieldContext_Todo_done(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Todo", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputNewTodo(ctx context.Context, obj interface{}) (model.NewTodo, error) {
	var it model.NewTodo
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"text"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "text":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("text"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Text = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputTodoFilter(ctx context.Context, obj interface{}) (models.Filter, error) {
	var it models.Filter
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"AND", "OR", "createdAt", "updatedAt", "text", "done"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "AND":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("AND"))
			data, err := ec.unmarshalOTodoFilter2ᚕᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐFilterᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.AND = data
		case "OR":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("OR"))
			data, err := ec.unmarshalOTodoFilter2ᚕᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐFilterᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.OR = data
		case "createdAt":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("createdAt"))
			data, err := ec.unmarshalODateFilter2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐDateFilter(ctx, v)
			if err != nil {
				return it, err
			}
			it.CreatedAt = data
		case "updatedAt":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("updatedAt"))
			data, err := ec.unmarshalODateFilter2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐDateFilter(ctx, v)
			if err != nil {
				return it, err
			}
			it.UpdatedAt = data
		case "text":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("text"))
			data, err := ec.unmarshalOStringFilter2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐGenericFilter(ctx, v)
			if err != nil {
				return it, err
			}
			it.Text = data
		case "done":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("done"))
			data, err := ec.unmarshalOBooleanFilter2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐGenericFilter(ctx, v)
			if err != nil {
				return it, err
			}
			it.Done = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputTodoSortOrder(ctx context.Context, obj interface{}) (models.SortOrder, error) {
	var it models.SortOrder
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"createdAt", "updatedAt", "text", "done"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "createdAt":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("createdAt"))
			data, err := ec.unmarshalOSortOrderEnum2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐSortOrderEnum(ctx, v)
			if err != nil {
				return it, err
			}
			it.CreatedAt = data
		case "updatedAt":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("updatedAt"))
			data, err := ec.unmarshalOSortOrderEnum2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐSortOrderEnum(ctx, v)
			if err != nil {
				return it, err
			}
			it.UpdatedAt = data
		case "text":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("text"))
			data, err := ec.unmarshalOSortOrderEnum2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐSortOrderEnum(ctx, v)
			if err != nil {
				return it, err
			}
			it.Text = data
		case "done":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("done"))
			data, err := ec.unmarshalOSortOrderEnum2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋdatabaseᚋcommonᚐSortOrderEnum(ctx, v)
			if err != nil {
				return it, err
			}
			it.Done = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpdateTodo(ctx context.Context, obj interface{}) (model.UpdateTodo, error) {
	var it model.UpdateTodo
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "text"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNID2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "text":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("text"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Text = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var todoImplementors = []string{"Todo"}

func (ec *executionContext) _Todo(ctx context.Context, sel ast.SelectionSet, obj *models.Todo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, todoImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Todo")
		case "id":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Todo_id(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "createdAt":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Todo_createdAt(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "updatedAt":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Todo_updatedAt(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "text":
			out.Values[i] = ec._Todo_text(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "done":
			out.Values[i] = ec._Todo_done(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var todoConnectionImplementors = []string{"TodoConnection"}

func (ec *executionContext) _TodoConnection(ctx context.Context, sel ast.SelectionSet, obj *model.TodoConnection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, todoConnectionImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("TodoConnection")
		case "edges":
			out.Values[i] = ec._TodoConnection_edges(ctx, field, obj)
		case "pageInfo":
			out.Values[i] = ec._TodoConnection_pageInfo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var todoEdgeImplementors = []string{"TodoEdge"}

func (ec *executionContext) _TodoEdge(ctx context.Context, sel ast.SelectionSet, obj *model.TodoEdge) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, todoEdgeImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("TodoEdge")
		case "cursor":
			out.Values[i] = ec._TodoEdge_cursor(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "node":
			out.Values[i] = ec._TodoEdge_node(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNNewTodo2githubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋmodelᚐNewTodo(ctx context.Context, v interface{}) (model.NewTodo, error) {
	res, err := ec.unmarshalInputNewTodo(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNTodo2githubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐTodo(ctx context.Context, sel ast.SelectionSet, v models.Todo) graphql.Marshaler {
	return ec._Todo(ctx, sel, &v)
}

func (ec *executionContext) marshalNTodo2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐTodo(ctx context.Context, sel ast.SelectionSet, v *models.Todo) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Todo(ctx, sel, v)
}

func (ec *executionContext) unmarshalNTodoFilter2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐFilter(ctx context.Context, v interface{}) (*models.Filter, error) {
	res, err := ec.unmarshalInputTodoFilter(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOTodo2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐTodo(ctx context.Context, sel ast.SelectionSet, v *models.Todo) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Todo(ctx, sel, v)
}

func (ec *executionContext) marshalOTodoConnection2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋmodelᚐTodoConnection(ctx context.Context, sel ast.SelectionSet, v *model.TodoConnection) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._TodoConnection(ctx, sel, v)
}

func (ec *executionContext) marshalOTodoEdge2ᚕᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋmodelᚐTodoEdge(ctx context.Context, sel ast.SelectionSet, v []*model.TodoEdge) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOTodoEdge2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋmodelᚐTodoEdge(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

func (ec *executionContext) marshalOTodoEdge2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋmodelᚐTodoEdge(ctx context.Context, sel ast.SelectionSet, v *model.TodoEdge) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._TodoEdge(ctx, sel, v)
}

func (ec *executionContext) unmarshalOTodoFilter2ᚕᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐFilterᚄ(ctx context.Context, v interface{}) ([]*models.Filter, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*models.Filter, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNTodoFilter2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐFilter(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalOTodoFilter2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐFilter(ctx context.Context, v interface{}) (*models.Filter, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputTodoFilter(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOTodoSortOrder2ᚕᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐSortOrder(ctx context.Context, v interface{}) ([]*models.SortOrder, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*models.SortOrder, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalOTodoSortOrder2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐSortOrder(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalOTodoSortOrder2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋbusinessᚋtodosᚋmodelsᚐSortOrder(ctx context.Context, v interface{}) (*models.SortOrder, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputTodoSortOrder(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOUpdateTodo2ᚖgithubᚗcomᚋoxynoᚑzetaᚋgolangᚑgraphqlᚑexampleᚋpkgᚋgolangᚑgraphqlᚑexampleᚋserverᚋgraphqlᚋmodelᚐUpdateTodo(ctx context.Context, v interface{}) (*model.UpdateTodo, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputUpdateTodo(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
