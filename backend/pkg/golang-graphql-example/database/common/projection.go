package common

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
)

const maxOptionsSplitLength = 2

func ManageProjection(projection any, db *gorm.DB) (*gorm.DB, error) {
	// Create result
	res := db
	// Get reflect value of projection object
	rVal := reflect.ValueOf(projection)
	// Get kind of projection
	rKind := rVal.Kind()
	// Check if projection isn't nil
	if rKind == reflect.Invalid || (rKind == reflect.Pointer && rVal.IsNil()) {
		// Stop here
		return res, nil
	}
	// Check if kind is supported
	if rKind != reflect.Struct && rKind != reflect.Pointer {
		// No skip => Error
		return nil, errors.NewInvalidInputError("projection must be an object")
	}

	// Indirect value
	indirect := reflect.Indirect(rVal)
	// Get indirect data
	indData := indirect.Interface()
	// Get type of indirect value
	typeOfIndi := reflect.TypeOf(indData)

	// Create array of projection results
	selectArray := make([]string, 0)
	// Store boolean to avoid checking length of select array which consum
	selectArrayFilled := false

	// Loop over all num fields
	for i := 0; i < indirect.NumField(); i++ {
		// Get field type
		fType := typeOfIndi.Field(i)
		// Get tag on field
		tagVal := fType.Tag.Get(dbColTagName)
		// Try to split to get options
		tagValSplit := strings.Split(tagVal, ";")
		// Too much options check
		if len(tagValSplit) > maxOptionsSplitLength {
			return nil, errors.NewInvalidInputError(fmt.Sprintf("field %s with too much options in tag %s", fType.Name, dbColTagName))
		}
		// Override tag if needed and save tag options
		options := make([]string, 0)
		// Manage save of possible options
		if len(tagValSplit) > 1 {
			tagVal = tagValSplit[0]
			options = tagValSplit[1:]
		}
		// Check that field have a tag set and correct
		if tagVal == "" || tagVal == "-" {
			// Skip this value
			continue
		}
		// Get field value
		fVal := indirect.Field(i)
		// Check if value is a boolean or not
		if fVal.Kind() != reflect.Bool {
			return nil, errors.NewInvalidInputError(
				fmt.Sprintf("field %s with projection tag must be a boolean", fType.Name),
			)
		}
		// Get value from field
		val := fVal.Interface()
		// Cast it to boolean
		v, _ := val.(bool)

		// Check if option is present
		// ? Note: We can work like this today because we only have 1 option
		if len(options) != 0 {
			// Get option
			option := options[0]
			// Check if it is the always fetch option
			if option != dbTagValueAlwaysFetch {
				return nil, errors.NewInvalidInputError(fmt.Sprintf("field %s unsupported option in tag %s", fType.Name, dbColTagName))
			}

			// Force fetch and so override possible value
			v = true
		}

		// Manage projection if enabled
		if v {
			selectArray = append(selectArray, tagVal)
			selectArrayFilled = true
		}
	}

	// Check if projection array is filled or not
	if selectArrayFilled {
		res = res.Select(selectArray)
	}

	// Default case
	return res, nil
}
