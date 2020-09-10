package pagination

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
)

func manageSortOrder(sort interface{}, db *gorm.DB) (*gorm.DB, error) {
	// Check if sort isn't nil
	if sort == nil {
		// Stop here
		return db, nil
	}

	// Create result
	res := db
	// Get reflect value of sort object
	rVal := reflect.ValueOf(sort)
	// Get kind of sort
	rKind := rVal.Kind()
	// Check if kind is supported
	if rKind != reflect.Struct && rKind != reflect.Ptr {
		return nil, errors.NewInternalServerError("sort must be an object")
	}

	// Indirect value
	indirect := reflect.Indirect(rVal)
	// Get indirect data
	indData := indirect.Interface()
	// Get type of indirect value
	typeOfIndi := reflect.TypeOf(indData)
	// Get type of searched enum
	supportedEnumType := reflect.TypeOf(&SortOrderEnumAsc)

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
			return nil, errors.NewInternalServerError("field with sort tag must be a *SortOrderEnum")
		}
		// Get field value
		fVal := indirect.Field(i)
		// Test if field is nil
		if fVal.IsNil() {
			// Skip field because of nil
			continue
		}
		// Get value from field
		val := fVal.Interface()
		// Cast value to Sort Order Enum
		enu := val.(*SortOrderEnum)
		// Apply order
		res = res.Order(fmt.Sprintf("%s %s", tagVal, enu.String()))
	}

	return res, nil
}
