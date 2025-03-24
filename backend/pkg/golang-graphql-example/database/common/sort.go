package common

import (
	"fmt"
	"reflect"

	gerrors "emperror.dev/errors"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"gorm.io/gorm"
)

// ErrSortListMustNotHaveMultipleFields error.
var ErrSortListMustNotHaveMultipleFields = gerrors.New(
	"sort list objects mustn't have multiple fields with sort values into the same object",
)

// Supported enum type for testing purpose.
var supportedEnumType = reflect.TypeOf(new(SortOrderEnum))

func ManageSortOrder(sort interface{}, db *gorm.DB) (*gorm.DB, error) {
	return manageSortOrder(sort, db)
}

func manageSortOrder(sort interface{}, db *gorm.DB) (*gorm.DB, error) {
	// Get reflect value of sort object
	rVal := reflect.ValueOf(sort)
	// Get kind of sort
	rKind := rVal.Kind()
	// Check nil
	if rKind == reflect.Invalid || (rKind == reflect.Ptr && rVal.IsNil()) {
		// Stop here
		return manageDefaultSort(db), nil
	}

	// Check if it is a slice
	if rKind == reflect.Array || rKind == reflect.Slice {
		return manageListSortOrder(&rVal, db)
	}

	// Manage object as default
	res, sortApplied, err := manageObjectSortOrder(rKind, &rVal, false, db)
	// Check error
	if err != nil {
		return nil, err
	}
	// Check if one sort was applied or not in order to put the default one
	if !sortApplied {
		res = manageDefaultSort(res)
	}

	// Default
	return res, nil
}

func manageListSortOrder(rVal *reflect.Value, db *gorm.DB) (*gorm.DB, error) {
	// Create result
	res := db
	// Initialize error
	var err error
	// Initialize sort applied marker
	sortApplied := false

	// Loop over slice
	for i := 0; i < rVal.Len(); i++ {
		// Get value
		rElem := rVal.Index(i)
		// Initialize sort applied result
		sortAppliedRes := false //nolint:wastedassign // False positive
		// Manage object
		res, sortAppliedRes, err = manageObjectSortOrder(rElem.Kind(), &rElem, true, res)
		// Check error
		if err != nil {
			return nil, err
		}
		// Save sort applied
		sortApplied = sortApplied || sortAppliedRes
	}

	// Check if one sort was applied or not in order to put the default one
	if !sortApplied {
		res = manageDefaultSort(res)
	}

	return res, nil
}

func manageObjectSortOrder(
	rKind reflect.Kind,
	rVal *reflect.Value,
	refuseMultipleField bool,
	db *gorm.DB,
) (*gorm.DB, bool, error) {
	// Create result
	res := db
	// Check if kind is supported
	if rKind != reflect.Struct && rKind != reflect.Ptr {
		return nil, false, errors.NewInvalidInputError("sort must be an object")
	}

	// Indirect value
	indirect := reflect.Indirect(*rVal)
	// Get indirect data
	indData := indirect.Interface()
	// Get type of indirect value
	typeOfIndi := reflect.TypeOf(indData)
	// Variable to know at the end if one sort was applied
	sortApplied := false

	// Loop over all num fields
	for i := 0; i < indirect.NumField(); i++ {
		// Get field type
		fType := typeOfIndi.Field(i)
		// Get tag on field
		tagVal := fType.Tag.Get(dbColTagName)
		// Check that field have a tag set and correct
		if tagVal == "" || tagVal == "-" {
			// Skip this value
			continue
		}
		// Check that type is supported
		if fType.Type != supportedEnumType {
			return nil, false, errors.NewInvalidInputError(
				fmt.Sprintf("field %s with sort tag must be a *SortOrderEnum", fType.Name),
			)
		}
		// Get field value
		fVal := indirect.Field(i)
		// Test if field is nil
		if fVal.IsNil() {
			// Skip field because of nil
			continue
		}
		// Check if sort have been already applied
		if refuseMultipleField && sortApplied {
			return nil, false, errors.NewInvalidInputErrorWithError(
				ErrSortListMustNotHaveMultipleFields,
				errors.WithPublicError(ErrSortListMustNotHaveMultipleFields),
			)
		}

		// Get value from field
		val := fVal.Interface()
		// Cast value to Sort Order Enum
		enu, ok := val.(*SortOrderEnum)
		// Check if it is ok or not
		if !ok {
			return nil, false, gerrors.Errorf("%v isn't a valid SortOrderEnum value", val)
		}
		// Apply order
		res = res.Order(fmt.Sprintf("%s %s", tagVal, enu.String()))
		// Store sort applied
		sortApplied = true
	}

	return res, sortApplied, nil
}

func manageDefaultSort(db *gorm.DB) *gorm.DB {
	return db.Order("created_at DESC")
}
