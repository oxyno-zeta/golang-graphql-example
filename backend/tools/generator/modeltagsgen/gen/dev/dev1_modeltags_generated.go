// Code generated by ModelTags. DO NOT EDIT.
package dev

import "emperror.dev/errors"

// ErrDev1UnsupportedGormColumn will be thrown when an unsupported Gorm column will be found in transform function.
var ErrDev1UnsupportedGormColumn = errors.Sentinel("unsupported gorm column")

// ErrDev1UnsupportedJSONKey will be thrown when an unsupported JSON key will be found in transform function.
var ErrDev1UnsupportedJSONKey = errors.Sentinel("unsupported json key")

// Dev1 CreatedAt Gorm Column Name
const Dev1CreatedAtGormColumnName = "created_at"

// Dev1 DeletedAt Gorm Column Name
const Dev1DeletedAtGormColumnName = "deleted_at"

// Dev1 Field2 Gorm Column Name
const Dev1Field2GormColumnName = "fake2"

// Dev1 Field3 Gorm Column Name
const Dev1Field3GormColumnName = "fake3"

// Dev1 Field4 Gorm Column Name
const Dev1Field4GormColumnName = "field4"

// Dev1 ID Gorm Column Name
const Dev1IDGormColumnName = "id"

// Dev1 UpdatedAt Gorm Column Name
const Dev1UpdatedAtGormColumnName = "updated_at"

// Dev1 CreatedAt JSON Key Name
const Dev1CreatedAtJSONKeyName = "createdAt"

// Dev1 DeletedAt JSON Key Name
const Dev1DeletedAtJSONKeyName = "deletedAt"

// Dev1 Field1 JSON Key Name
const Dev1Field1JSONKeyName = "field1"

// Dev1 Field3 JSON Key Name
const Dev1Field3JSONKeyName = "field3"

// Dev1 Field4 JSON Key Name
const Dev1Field4JSONKeyName = "field4"

// Dev1 ID JSON Key Name
const Dev1IDJSONKeyName = "id"

// Dev1 UpdatedAt JSON Key Name
const Dev1UpdatedAtJSONKeyName = "updatedAt"

// Transform Dev1 Gorm Column To JSON Key
func TransformDev1GormColumnToJSONKey(gormColumn string) (string, error) {
	switch gormColumn {
	case Dev1CreatedAtGormColumnName:
		return Dev1CreatedAtJSONKeyName, nil
	case Dev1DeletedAtGormColumnName:
		return Dev1DeletedAtJSONKeyName, nil
	case Dev1Field3GormColumnName:
		return Dev1Field3JSONKeyName, nil
	case Dev1Field4GormColumnName:
		return Dev1Field4JSONKeyName, nil
	case Dev1IDGormColumnName:
		return Dev1IDJSONKeyName, nil
	case Dev1UpdatedAtGormColumnName:
		return Dev1UpdatedAtJSONKeyName, nil
	default:
		return "", errors.WithStack(ErrDev1UnsupportedGormColumn)
	}
}

// Transform Dev1 JSON Key To Gorm Column
func TransformDev1JSONKeyToGormColumn(jsonKey string) (string, error) {
	switch jsonKey {
	case Dev1CreatedAtJSONKeyName:
		return Dev1CreatedAtGormColumnName, nil
	case Dev1DeletedAtJSONKeyName:
		return Dev1DeletedAtGormColumnName, nil
	case Dev1Field3JSONKeyName:
		return Dev1Field3GormColumnName, nil
	case Dev1Field4JSONKeyName:
		return Dev1Field4GormColumnName, nil
	case Dev1IDJSONKeyName:
		return Dev1IDGormColumnName, nil
	case Dev1UpdatedAtJSONKeyName:
		return Dev1UpdatedAtGormColumnName, nil
	default:
		return "", errors.WithStack(ErrDev1UnsupportedJSONKey)
	}
}

// Transform Dev1 JSON Key map To Gorm Column map
func TransformDev1JSONKeyMapToGormColumnMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformDev1JSONKeyToGormColumn(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrDev1UnsupportedJSONKey) {
				// Continue the loop
				continue
			}

			// Return
			return nil, err
		}
		// Save
		m[r] = v
	}

	return m, nil
}

// Transform Dev1 Gorm Column map To JSON Key map
func TransformDev1GormColumnMapToJSONKeyMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformDev1GormColumnToJSONKey(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrDev1UnsupportedGormColumn) {
				// Continue the loop
				continue
			}

			// Return
			return nil, err
		}
		// Save
		m[r] = v
	}

	return m, nil
}
