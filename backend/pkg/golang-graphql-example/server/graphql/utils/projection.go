package utils

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"github.com/samber/lo"
	"github.com/thoas/go-funk"
)

const graphqlFieldTagKey = "graphqlfield"

func ManageConnectionNodeProjection(
	ctx context.Context,
	projectionOut interface{},
) error {
	return ManageDepthProjection(ctx, projectionOut, []string{"edges", "node"})
}

func ManageDepthProjection(
	ctx context.Context,
	projectionOut interface{},
	fieldChain []string,
) error {
	// Validate projection out
	err := validateProjectionOut(projectionOut)
	// Check error
	if err != nil {
		return err
	}

	// Get operation context
	octx := graphql.GetOperationContext(ctx)
	// Get graphql fields
	fields := graphql.CollectFieldsCtx(ctx, nil)

	// Dive to get collected fields under chain
	collectedFields := diveToGraphqlCollectedField(octx, fields, fieldChain)

	// Start projection on this path
	err = manageGraphqlProjection(
		collectedFields,
		projectionOut,
	)
	// Check error
	if err != nil {
		return err
	}

	// Default
	return nil
}

func ManageSimpleProjection(
	ctx context.Context,
	projectionOut interface{},
) error {
	return ManageDepthProjection(ctx, projectionOut, []string{})
}

func diveToGraphqlCollectedField(
	octx *graphql.OperationContext,
	fields []graphql.CollectedField,
	fieldChain []string,
) []graphql.CollectedField {
	// Check if field chain have values
	if len(fieldChain) == 0 {
		return fields
	}

	// Pop first item on chain
	f, fieldChain := fieldChain[0], fieldChain[1:]

	// Find collected field
	fieldC, ok := lo.Find(fields, func(fieldC graphql.CollectedField) bool {
		return fieldC.Name == f
	})
	// Check if it is found
	if !ok {
		return nil
	}

	// Get collected fields from field selections
	subFieldCollections := graphql.CollectFields(octx, fieldC.Selections, nil)

	// Dive
	return diveToGraphqlCollectedField(octx, subFieldCollections, fieldChain)
}

func validateProjectionOut(projectionOut interface{}) error {
	// Check if input is nil
	if projectionOut == nil {
		return errors.NewInvalidInputError("projection output cannot be nil")
	}

	// Get projection type
	projOutType := reflect.TypeOf(projectionOut)
	// Check if projection out is a pointer
	if projOutType.Kind() != reflect.Ptr {
		return errors.NewInvalidInputError("projection output must be a pointer to an object")
	}
	// Get projection out value
	projOutVal := reflect.ValueOf(projectionOut)
	// Indirect value
	indVal := reflect.Indirect(projOutVal)
	// Check indirect value type
	if indVal.Kind() != reflect.Struct {
		return errors.NewInvalidInputError("projection output must be a pointer to an object")
	}

	// Default
	return nil
}

func manageGraphqlProjection(
	gfields []graphql.CollectedField,
	projectionOut interface{},
) error {
	// Get reflect ptr value
	pOutPtrVal := reflect.ValueOf(projectionOut)
	// Get reflect value
	pOutVal := reflect.Indirect(pOutPtrVal)
	// Get reflect ptr type
	pOutPtrType := reflect.TypeOf(projectionOut)
	// Get reflect type
	pOutType := pOutPtrType.Elem()

	// Loop over projection struct fields
	for i := 0; i < pOutType.NumField(); i++ {
		// Get field
		fieldType := pOutType.Field(i)
		// Get tag value for graphql field
		tagValue := fieldType.Tag.Get(graphqlFieldTagKey)
		// Check if tag exists or ignored
		if tagValue == "" || tagValue == "-" {
			// Continue to next field
			continue
		}

		// Split tag value for multiple fields
		tagValues := strings.Split(tagValue, ",")

		// Loop over tag values
		for _, value := range tagValues {
			// Check if tag isn't marked as ignored
			if value == "-" {
				// Continue to next field
				continue
			}

			// Manage field projection
			err := manageFieldProjection(
				gfields,
				value,
				&pOutVal,
				i,
				&fieldType,
			)
			// Check error
			if err != nil {
				return err
			}
		}
	}

	// Default
	return nil
}

func manageFieldProjection(
	gfields []graphql.CollectedField,
	tagValue string,
	pOutVal *reflect.Value,
	fieldNumber int,
	fieldType *reflect.StructField,
) error {
	// Check if field is asked in graphql
	gfieldInt := funk.Find(gfields, func(gfield graphql.CollectedField) bool {
		return gfield.Name == tagValue
	})
	// Check if field isn't found
	if gfieldInt == nil {
		// Field isn't found => continue to next field
		return nil
	}

	// Check if field is a boolean
	if fieldType.Type.Kind() == reflect.Bool {
		pOutVal.Field(fieldNumber).SetBool(true)
		// Stop here
		return nil
	}

	// Field is found but type isn't supported
	return errors.NewInvalidInputError(fmt.Sprintf("field %s must be a boolean", fieldType.Name))
}
