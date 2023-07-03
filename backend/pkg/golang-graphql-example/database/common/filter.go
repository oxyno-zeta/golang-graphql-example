package common

import (
	"fmt"
	"reflect"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
	"gorm.io/gorm"
)

// AND field.
const andFieldName = "AND"

// OR field.
const orFieldName = "OR"

func ManageFilter(filter interface{}, db *gorm.DB) (*gorm.DB, error) {
	return manageFilter(filter, db, false)
}

func manageFilter(filter interface{}, originalDB *gorm.DB, skipInputNotObject bool) (*gorm.DB, error) {
	// Get reflect value of filter object
	rVal := reflect.ValueOf(filter)
	// Get kind of filter
	rKind := rVal.Kind()
	// Check if filter isn't nil
	if rKind == reflect.Invalid || (rKind == reflect.Ptr && rVal.IsNil()) {
		// Stop here
		return originalDB, nil
	}
	// Check if kind is supported
	if rKind != reflect.Struct && rKind != reflect.Ptr {
		// Check if skip input not an object is enabled
		// This is used in recursive calls in order to avoid errors when OR or AND cases aren't an object supported
		if skipInputNotObject {
			return originalDB, nil
		}

		// No skip => Error
		return nil, errors.NewInvalidInputError("filter must be an object")
	}

	// Build result
	res := originalDB
	// Indirect value
	indirect := reflect.Indirect(rVal)
	// Get indirect data
	indData := indirect.Interface()
	// Get type of indirect value
	typeOfIndi := reflect.TypeOf(indData)

	// Create a temporary db object
	// This is made to ensure groups with ()
	// Without this, there is a high risk of not having them well placed
	var tmpDB *gorm.DB
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
			return nil, errors.NewInvalidInputError(
				fmt.Sprintf("field %s with filter tag must be a *GenericFilter or implement GenericFilterBuilder interface", fType.Name),
			)
		}
		// Test if field is nil
		if fVal.IsNil() {
			// Skip field because of nil
			continue
		}
		// Get value from field
		val := fVal.Interface()

		// Try to cast it as GenericFilterBuilder
		v1, castGFB := val.(GenericFilterBuilder)
		// Check that type is supported
		if !castGFB {
			return nil, errors.NewInvalidInputError(
				fmt.Sprintf("field %s with filter tag must be a *GenericFilter or implement GenericFilterBuilder interface", fType.Name),
			)
		}

		// Cast value
		v, err := v1.GetGenericFilter()
		// Check error
		if err != nil {
			return nil, err
		}
		// Manage filter request
		// Call manage filter request WITH the original db in order to create a pure subquery
		// See here: https://gorm.io/docs/advanced_query.html#Group-Conditions
		res2, err := manageFilterRequest(tagVal, v, originalDB)
		// Check error
		if err != nil {
			return nil, err
		}

		// Manage result
		// Check if it is the first time
		if tmpDB == nil {
			// Init with result
			tmpDB = res2
		} else {
			// Simply add the value with the where to the previous one
			tmpDB = tmpDB.Where(res2)
		}
	}

	// Check if it is set
	if tmpDB != nil {
		// Init the result with this
		res = tmpDB
	}

	// Manage AND cases
	// Check in type that AND key exists
	_, exists := typeOfIndi.FieldByName(andFieldName)
	// Check if it exists
	if exists {
		// AND field is detected
		// Get field AND
		andRVal := indirect.FieldByName(andFieldName)
		// Check that type is a slice
		if andRVal.Kind() == reflect.Slice {
			// Create a temporary db object
			// This is made to ensure groups with ()
			// Without this, there is a high risk of not having them well placed
			var tmpDB *gorm.DB
			// Loop over items in array
			for i := 0; i < andRVal.Len(); i++ {
				// Get element at index
				andElementRVal := andRVal.Index(i)
				// Get value behind
				andElement := andElementRVal.Interface()
				// Call manage filter WITH the original db in order to create a pure subquery
				// See here: https://gorm.io/docs/advanced_query.html#Group-Conditions
				res2, err := manageFilter(andElement, originalDB, true)
				// Check error
				if err != nil {
					return nil, err
				}
				// Manage result
				// Check if it is the first time
				if tmpDB == nil {
					// Init with result
					tmpDB = res2
				} else {
					// Simply add the value with the where to the previous one
					tmpDB = tmpDB.Where(res2)
				}
			}
			// Add result to existing filter if it exists
			if tmpDB != nil {
				res = res.Where(tmpDB)
			}
		}
	}

	// Manage OR cases
	// Check in type that OR key exists
	_, exists = typeOfIndi.FieldByName(orFieldName)
	// Check if it exists
	if exists {
		// OR field is detected
		// Get field OR
		orRVal := indirect.FieldByName(orFieldName)

		// Check that type is a slice
		if orRVal.Kind() == reflect.Slice {
			// Get array length
			lgt := orRVal.Len()
			// Check length in order to ignore it it is 0
			if lgt != 0 {
				// Create a temporary db object
				// This is made to ensure OR filters are grouped together in request
				// Otherwise, we can have the situation of XX AND YY OR ZZ
				// instead of XX AND (YY OR ZZ) or (XX AND YY) OR ZZ
				// In fact, we ensure groups with () with this solution
				var tmpDB *gorm.DB
				// Array is populated
				// Loop over elements
				for i := 0; i < lgt; i++ {
					// Get element
					elemRVal := orRVal.Index(i)
					// Get data behind
					elem := elemRVal.Interface()
					// Call manage filter WITH the original db in order to create a pure subquery
					// See here: https://gorm.io/docs/advanced_query.html#Group-Conditions
					res2, err := manageFilter(elem, originalDB, true)
					// Check error
					if err != nil {
						return nil, err
					}
					// Manage result
					// Check if it is the first time
					if tmpDB == nil {
						// Init with result but as it it is a OR case, add to a clean db object to ensure no perturbation
						tmpDB = originalDB.Or(res2)
					} else {
						// Simply add the value with the or to the previous one
						tmpDB = tmpDB.Or(res2)
					}
				}

				// Add result to existing filter if it exists
				if tmpDB != nil {
					res = res.Where(tmpDB)
				}
			}
		}
	}

	// Return
	return res, nil
}

func GenerateQueryTemplate(
	operation string,
	value string,
	fieldUppercase bool,
	fieldLowercase bool,
	valueUppercase bool,
	valueLowercase bool,
) string {
	first := "%s"
	second := value

	// Manage field
	if fieldLowercase {
		first = "lower(" + first + ")"
	} else if fieldUppercase {
		first = "upper(" + first + ")"
	}

	// Manage value
	if valueLowercase {
		second = "lower(" + second + ")"
	} else if valueUppercase {
		second = "upper(" + second + ")"
	}

	// Result
	return first + " " + operation + " " + second
}

func manageFilterRequest(dbCol string, v *GenericFilter, db *gorm.DB) (*gorm.DB, error) {
	// Create result
	dbRes := db
	// Check Equal case
	if v.Eq != nil {
		tpl := GenerateQueryTemplate("=", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Where(fmt.Sprintf(tpl, dbCol), v.Eq)
	}
	// Check not equal case
	if v.NotEq != nil {
		tpl := GenerateQueryTemplate("=", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Not(fmt.Sprintf(tpl, dbCol), v.NotEq)
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

		tpl := GenerateQueryTemplate("LIKE", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Where(fmt.Sprintf(tpl, dbCol), fmt.Sprintf("%%%s%%", s))
	}
	// Check not contains case
	if v.NotContains != nil {
		// Get string value
		s, err := getStringValue(v.NotContains)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("notContains " + err.Error())
		}

		tpl := GenerateQueryTemplate("LIKE", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Not(fmt.Sprintf(tpl, dbCol), fmt.Sprintf("%%%s%%", s))
	}
	// Check starts with case
	if v.StartsWith != nil {
		// Get string value
		s, err := getStringValue(v.StartsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("startsWith " + err.Error())
		}

		tpl := GenerateQueryTemplate("LIKE", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Where(fmt.Sprintf(tpl, dbCol), fmt.Sprintf("%s%%", s))
	}
	// Check not starts with case
	if v.NotStartsWith != nil {
		// Get string value
		s, err := getStringValue(v.NotStartsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("notStartsWith " + err.Error())
		}

		tpl := GenerateQueryTemplate("LIKE", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Not(fmt.Sprintf(tpl, dbCol), fmt.Sprintf("%s%%", s))
	}
	// Check ends with case
	if v.EndsWith != nil {
		// Get string value
		s, err := getStringValue(v.EndsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("endsWith " + err.Error())
		}

		tpl := GenerateQueryTemplate("LIKE", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Where(fmt.Sprintf(tpl, dbCol), fmt.Sprintf("%%%s", s))
	}
	// Check not ends with case
	if v.NotEndsWith != nil {
		// Get string value
		s, err := getStringValue(v.NotEndsWith)
		// Check error
		if err != nil {
			return nil, errors.NewInvalidInputError("notEndsWith " + err.Error())
		}

		tpl := GenerateQueryTemplate("LIKE", "?", v.FieldUppercase, v.FieldLowercase, v.ValueUppercase, v.ValueLowercase)
		dbRes = dbRes.Not(fmt.Sprintf(tpl, dbCol), fmt.Sprintf("%%%s", s))
	}
	// Check in case
	if v.In != nil {
		tpl := GenerateQueryTemplate("IN", "(?)", v.FieldUppercase, v.FieldLowercase, false, false)
		dbRes = dbRes.Where(fmt.Sprintf(tpl, dbCol), v.In)
	}
	// Check not in case
	if v.NotIn != nil {
		tpl := GenerateQueryTemplate("NOT IN", "(?)", v.FieldUppercase, v.FieldLowercase, false, false)
		dbRes = dbRes.Where(fmt.Sprintf(tpl, dbCol), v.NotIn)
	}
	// Check is null case
	if v.IsNull {
		dbRes = dbRes.Where(fmt.Sprintf("%s IS NULL", dbCol))
	}
	// Check is not null case
	if v.IsNotNull {
		dbRes = dbRes.Where(fmt.Sprintf("%s IS NOT NULL", dbCol))
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
