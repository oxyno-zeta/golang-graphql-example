package pagination

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
)

// Supported generic filter type for testing purpose
var supportedGenericFilterType = reflect.TypeOf(new(GenericFilter))

func manageFilter(filter interface{}, db *gorm.DB) (*gorm.DB, error) {
	// Check if filter isn't nil
	if filter == nil {
		// Stop here
		return db, nil
	}

	// Create result
	res := db
	// Get reflect value of filter object
	rVal := reflect.ValueOf(filter)
	// Get kind of filter
	rKind := rVal.Kind()
	// Check if kind is supported
	if rKind != reflect.Struct && rKind != reflect.Ptr {
		return nil, errors.NewInternalServerError("filter must be an object")
	}

	// Indirect value
	indirect := reflect.Indirect(rVal)
	// Get indirect data
	indData := indirect.Interface()
	// Get type of indirect value
	typeOfIndi := reflect.TypeOf(indData)

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
		// Get field value
		fVal := indirect.Field(i)
		// Check if value is a pointer or not
		if fVal.Kind() != reflect.Ptr {
			return nil, errors.NewInternalServerErrorWithError(
				fmt.Errorf("field %s with filter tag must be a pointer to an object", fType.Name),
			)
		}
		// Test if field is nil
		if fVal.IsNil() {
			// Skip field because of nil
			continue
		}
		// Get value from field
		// val := fVal.Interface()

		switch fType.Type {
		case supportedGenericFilterType:
			// Cast value
			// v := val.(*GenericFilter)
			// res = manageGenericFilter(tagVal, v, res)
		default:
			return nil, errors.NewInternalServerErrorWithError(
				fmt.Errorf("unsupported field type for filter: %s", fType.Name),
			)
		}
		// // Check that type is supported
		// if fType.Type != supportedEnumType {
		// 	return nil, errors.NewInternalServerError("field with filter tag must be a *SortOrderEnum")
		// }
		// // Cast value to Sort Order Enum
		// enu := val.(*SortOrderEnum)
		// // Apply order
		// res = res.Order(fmt.Sprintf("%s %s", tagVal, enu.String()))
	}

	// Return
	return res, nil
}

func manageGenericFilter(dbCol string, v *GenericFilter, db *gorm.DB) (*gorm.DB, error) {
	// Create result
	dbRes := db
	// Check Equal case
	if v.Eq != nil {
		dbRes = dbRes.Where(fmt.Sprintf("%s = ?", dbCol), v.Eq)
	}
	// Check not equal case
	if v.NotEq != nil {
		dbRes = dbRes.Not(fmt.Sprintf("%s = ?", dbCol), v.NotEq)
	}
	// Check greater and equal than case
	if v.Gte != nil {
		dbRes = dbRes.Where(fmt.Sprintf("%s >= ?", dbCol), v.Gte)
	}
	// Check not greater and equal than case
	if v.NotGte != nil {
		dbRes = dbRes.Not(fmt.Sprintf("%s >= ?", dbCol), v.NotGte)
	}
	// Check greater than case
	if v.Gt != nil {
		dbRes = dbRes.Where(fmt.Sprintf("%s > ?", dbCol), v.Gt)
	}
	// Check not greater than case
	if v.NotGt != nil {
		dbRes = dbRes.Not(fmt.Sprintf("%s > ?", dbCol), v.NotGt)
	}
	// Check less and equal than case
	if v.Lte != nil {
		dbRes = dbRes.Where(fmt.Sprintf("%s <= ?", dbCol), v.Lte)
	}
	// Check not less and equal than case
	if v.NotLte != nil {
		dbRes = dbRes.Not(fmt.Sprintf("%s <= ?", dbCol), v.NotLte)
	}
	// Check less than case
	if v.Lt != nil {
		dbRes = dbRes.Where(fmt.Sprintf("%s < ?", dbCol), v.Lt)
	}
	// Check not less than case
	if v.NotLt != nil {
		dbRes = dbRes.Not(fmt.Sprintf("%s < ?", dbCol), v.NotLt)
	}
	// Check contains case
	if v.Contains != nil {
		// Get string value
		s, err := getStringValue(v.Contains)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("contains " + err.Error())
		}

		dbRes = dbRes.Where(fmt.Sprintf("%s LIKE ?", dbCol), fmt.Sprintf("%%%s%%", s))
	}
	// Check not contains case
	if v.NotContains != nil {
		// Get string value
		s, err := getStringValue(v.NotContains)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("notContains " + err.Error())
		}

		dbRes = dbRes.Not(fmt.Sprintf("%s LIKE ?", dbCol), fmt.Sprintf("%%%s%%", s))
	}
	// Check starts with case
	if v.StartsWith != nil {
		// Get string value
		s, err := getStringValue(v.StartsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("startsWith " + err.Error())
		}

		dbRes = dbRes.Where(fmt.Sprintf("%s LIKE ?", dbCol), fmt.Sprintf("%s%%", s))
	}
	// Check not starts with case
	if v.NotStartsWith != nil {
		// Get string value
		s, err := getStringValue(v.NotStartsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("notStartsWith " + err.Error())
		}

		dbRes = dbRes.Not(fmt.Sprintf("%s LIKE ?", dbCol), fmt.Sprintf("%s%%", s))
	}
	// Check ends with case
	if v.EndsWith != nil {
		// Get string value
		s, err := getStringValue(v.EndsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("endsWith " + err.Error())
		}

		dbRes = dbRes.Where(fmt.Sprintf("%s LIKE ?", dbCol), fmt.Sprintf("%%%s", s))
	}
	// Check not ends with case
	if v.NotEndsWith != nil {
		// Get string value
		s, err := getStringValue(v.NotEndsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("notEndsWith " + err.Error())
		}

		dbRes = dbRes.Not(fmt.Sprintf("%s LIKE ?", dbCol), fmt.Sprintf("%%%s", s))
	}
	// Check in case
	if v.In != nil {
		dbRes = dbRes.Where(fmt.Sprintf("%s IN (?)", dbCol), v.In)
	}
	// Check not in case
	if v.NotIn != nil {
		dbRes = dbRes.Where(fmt.Sprintf("%s NOT IN (?)", dbCol), v.NotIn)
	}

	// Return
	return dbRes, nil
}

func getStringValue(x interface{}) (string, error) {
	// Get reflect value
	val := reflect.ValueOf(x)
	// Check if val is a pointer
	if val.Kind() == reflect.Ptr {
		// Indirect for pointers
		val = reflect.Indirect(val)
	}
	// Check if type is acceptable
	if val.Kind() != reflect.String {
		return "", errors.NewInvalidInputError("value must be a string or *string")
	}

	return val.String(), nil
}
