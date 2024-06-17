//go:build unit

package gen

import (
	"testing"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	type Object1 struct {
		database.Base
		Name string
	}
	type Object2 struct {
		database.Base
		Name string `json:"-"`
	}
	type Object3 struct {
		database.Base
		Name string `json:"name" gorm:"column:full__name"`
	}
	type Object4 struct {
		database.Base
		Name string `json:"name" gorm:"-"`
	}
	type Object5 struct {
		database.Base
		Name string `json:"name,omitempty"`
	}
	type args struct {
		pkgName string
		obj     interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple without json tag",
			args: args{
				pkgName: "fake",
				obj:     Object1{},
			},
			want: `// Code generated by ModelTags. DO NOT EDIT.
package fake

import "emperror.dev/errors"

// ErrObject1UnsupportedGormColumn will be thrown when an unsupported Gorm column will be found in transform function.
var ErrObject1UnsupportedGormColumn = errors.Sentinel("unsupported gorm column")

// ErrObject1UnsupportedJSONKey will be thrown when an unsupported JSON key will be found in transform function.
var ErrObject1UnsupportedJSONKey = errors.Sentinel("unsupported json key")

// Object1 CreatedAt Gorm Column Name
const Object1CreatedAtGormColumnName = "created_at"

// Object1 DeletedAt Gorm Column Name
const Object1DeletedAtGormColumnName = "deleted_at"

// Object1 ID Gorm Column Name
const Object1IDGormColumnName = "id"

// Object1 Name Gorm Column Name
const Object1NameGormColumnName = "name"

// Object1 UpdatedAt Gorm Column Name
const Object1UpdatedAtGormColumnName = "updated_at"

// Object1 CreatedAt JSON Key Name
const Object1CreatedAtJSONKeyName = "createdAt"

// Object1 DeletedAt JSON Key Name
const Object1DeletedAtJSONKeyName = "deletedAt"

// Object1 ID JSON Key Name
const Object1IDJSONKeyName = "id"

// Object1 Name JSON Key Name
const Object1NameJSONKeyName = "Name"

// Object1 UpdatedAt JSON Key Name
const Object1UpdatedAtJSONKeyName = "updatedAt"

// Transform Object1 Gorm Column To JSON Key
func TransformObject1GormColumnToJSONKey(gormColumn string) (string, error) {
	switch gormColumn {
	case Object1CreatedAtGormColumnName:
		return Object1CreatedAtJSONKeyName, nil
	case Object1DeletedAtGormColumnName:
		return Object1DeletedAtJSONKeyName, nil
	case Object1IDGormColumnName:
		return Object1IDJSONKeyName, nil
	case Object1NameGormColumnName:
		return Object1NameJSONKeyName, nil
	case Object1UpdatedAtGormColumnName:
		return Object1UpdatedAtJSONKeyName, nil
	default:
		return "", errors.WithStack(ErrObject1UnsupportedGormColumn)
	}
}

// Transform Object1 JSON Key To Gorm Column
func TransformObject1JSONKeyToGormColumn(jsonKey string) (string, error) {
	switch jsonKey {
	case Object1CreatedAtJSONKeyName:
		return Object1CreatedAtGormColumnName, nil
	case Object1DeletedAtJSONKeyName:
		return Object1DeletedAtGormColumnName, nil
	case Object1IDJSONKeyName:
		return Object1IDGormColumnName, nil
	case Object1NameJSONKeyName:
		return Object1NameGormColumnName, nil
	case Object1UpdatedAtJSONKeyName:
		return Object1UpdatedAtGormColumnName, nil
	default:
		return "", errors.WithStack(ErrObject1UnsupportedJSONKey)
	}
}

// Transform Object1 JSON Key map To Gorm Column map
func TransformObject1JSONKeyMapToGormColumnMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject1JSONKeyToGormColumn(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject1UnsupportedJSONKey) {
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

// Transform Object1 Gorm Column map To JSON Key map
func TransformObject1GormColumnMapToJSONKeyMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject1GormColumnToJSONKey(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject1UnsupportedGormColumn) {
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
`,
		},
		{
			name: "ignore a json key",
			args: args{
				pkgName: "fake",
				obj:     Object2{},
			},
			want: `// Code generated by ModelTags. DO NOT EDIT.
package fake

import "emperror.dev/errors"

// ErrObject2UnsupportedGormColumn will be thrown when an unsupported Gorm column will be found in transform function.
var ErrObject2UnsupportedGormColumn = errors.Sentinel("unsupported gorm column")

// ErrObject2UnsupportedJSONKey will be thrown when an unsupported JSON key will be found in transform function.
var ErrObject2UnsupportedJSONKey = errors.Sentinel("unsupported json key")

// Object2 CreatedAt Gorm Column Name
const Object2CreatedAtGormColumnName = "created_at"

// Object2 DeletedAt Gorm Column Name
const Object2DeletedAtGormColumnName = "deleted_at"

// Object2 ID Gorm Column Name
const Object2IDGormColumnName = "id"

// Object2 Name Gorm Column Name
const Object2NameGormColumnName = "name"

// Object2 UpdatedAt Gorm Column Name
const Object2UpdatedAtGormColumnName = "updated_at"

// Object2 CreatedAt JSON Key Name
const Object2CreatedAtJSONKeyName = "createdAt"

// Object2 DeletedAt JSON Key Name
const Object2DeletedAtJSONKeyName = "deletedAt"

// Object2 ID JSON Key Name
const Object2IDJSONKeyName = "id"

// Object2 UpdatedAt JSON Key Name
const Object2UpdatedAtJSONKeyName = "updatedAt"

// Transform Object2 Gorm Column To JSON Key
func TransformObject2GormColumnToJSONKey(gormColumn string) (string, error) {
	switch gormColumn {
	case Object2CreatedAtGormColumnName:
		return Object2CreatedAtJSONKeyName, nil
	case Object2DeletedAtGormColumnName:
		return Object2DeletedAtJSONKeyName, nil
	case Object2IDGormColumnName:
		return Object2IDJSONKeyName, nil
	case Object2UpdatedAtGormColumnName:
		return Object2UpdatedAtJSONKeyName, nil
	default:
		return "", errors.WithStack(ErrObject2UnsupportedGormColumn)
	}
}

// Transform Object2 JSON Key To Gorm Column
func TransformObject2JSONKeyToGormColumn(jsonKey string) (string, error) {
	switch jsonKey {
	case Object2CreatedAtJSONKeyName:
		return Object2CreatedAtGormColumnName, nil
	case Object2DeletedAtJSONKeyName:
		return Object2DeletedAtGormColumnName, nil
	case Object2IDJSONKeyName:
		return Object2IDGormColumnName, nil
	case Object2UpdatedAtJSONKeyName:
		return Object2UpdatedAtGormColumnName, nil
	default:
		return "", errors.WithStack(ErrObject2UnsupportedJSONKey)
	}
}

// Transform Object2 JSON Key map To Gorm Column map
func TransformObject2JSONKeyMapToGormColumnMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject2JSONKeyToGormColumn(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject2UnsupportedJSONKey) {
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

// Transform Object2 Gorm Column map To JSON Key map
func TransformObject2GormColumnMapToJSONKeyMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject2GormColumnToJSONKey(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject2UnsupportedGormColumn) {
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
`,
		},
		{
			name: "json is present but gorm key is different",
			args: args{
				pkgName: "fake",
				obj:     Object3{},
			},
			want: `// Code generated by ModelTags. DO NOT EDIT.
package fake

import "emperror.dev/errors"

// ErrObject3UnsupportedGormColumn will be thrown when an unsupported Gorm column will be found in transform function.
var ErrObject3UnsupportedGormColumn = errors.Sentinel("unsupported gorm column")

// ErrObject3UnsupportedJSONKey will be thrown when an unsupported JSON key will be found in transform function.
var ErrObject3UnsupportedJSONKey = errors.Sentinel("unsupported json key")

// Object3 CreatedAt Gorm Column Name
const Object3CreatedAtGormColumnName = "created_at"

// Object3 DeletedAt Gorm Column Name
const Object3DeletedAtGormColumnName = "deleted_at"

// Object3 ID Gorm Column Name
const Object3IDGormColumnName = "id"

// Object3 Name Gorm Column Name
const Object3NameGormColumnName = "full__name"

// Object3 UpdatedAt Gorm Column Name
const Object3UpdatedAtGormColumnName = "updated_at"

// Object3 CreatedAt JSON Key Name
const Object3CreatedAtJSONKeyName = "createdAt"

// Object3 DeletedAt JSON Key Name
const Object3DeletedAtJSONKeyName = "deletedAt"

// Object3 ID JSON Key Name
const Object3IDJSONKeyName = "id"

// Object3 Name JSON Key Name
const Object3NameJSONKeyName = "name"

// Object3 UpdatedAt JSON Key Name
const Object3UpdatedAtJSONKeyName = "updatedAt"

// Transform Object3 Gorm Column To JSON Key
func TransformObject3GormColumnToJSONKey(gormColumn string) (string, error) {
	switch gormColumn {
	case Object3CreatedAtGormColumnName:
		return Object3CreatedAtJSONKeyName, nil
	case Object3DeletedAtGormColumnName:
		return Object3DeletedAtJSONKeyName, nil
	case Object3IDGormColumnName:
		return Object3IDJSONKeyName, nil
	case Object3NameGormColumnName:
		return Object3NameJSONKeyName, nil
	case Object3UpdatedAtGormColumnName:
		return Object3UpdatedAtJSONKeyName, nil
	default:
		return "", errors.WithStack(ErrObject3UnsupportedGormColumn)
	}
}

// Transform Object3 JSON Key To Gorm Column
func TransformObject3JSONKeyToGormColumn(jsonKey string) (string, error) {
	switch jsonKey {
	case Object3CreatedAtJSONKeyName:
		return Object3CreatedAtGormColumnName, nil
	case Object3DeletedAtJSONKeyName:
		return Object3DeletedAtGormColumnName, nil
	case Object3IDJSONKeyName:
		return Object3IDGormColumnName, nil
	case Object3NameJSONKeyName:
		return Object3NameGormColumnName, nil
	case Object3UpdatedAtJSONKeyName:
		return Object3UpdatedAtGormColumnName, nil
	default:
		return "", errors.WithStack(ErrObject3UnsupportedJSONKey)
	}
}

// Transform Object3 JSON Key map To Gorm Column map
func TransformObject3JSONKeyMapToGormColumnMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject3JSONKeyToGormColumn(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject3UnsupportedJSONKey) {
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

// Transform Object3 Gorm Column map To JSON Key map
func TransformObject3GormColumnMapToJSONKeyMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject3GormColumnToJSONKey(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject3UnsupportedGormColumn) {
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
`,
		},
		{
			name: "ignore a gorm key but json is present",
			args: args{
				pkgName: "fake",
				obj:     Object4{},
			},
			want: `// Code generated by ModelTags. DO NOT EDIT.
package fake

import "emperror.dev/errors"

// ErrObject4UnsupportedGormColumn will be thrown when an unsupported Gorm column will be found in transform function.
var ErrObject4UnsupportedGormColumn = errors.Sentinel("unsupported gorm column")

// ErrObject4UnsupportedJSONKey will be thrown when an unsupported JSON key will be found in transform function.
var ErrObject4UnsupportedJSONKey = errors.Sentinel("unsupported json key")

// Object4 CreatedAt Gorm Column Name
const Object4CreatedAtGormColumnName = "created_at"

// Object4 DeletedAt Gorm Column Name
const Object4DeletedAtGormColumnName = "deleted_at"

// Object4 ID Gorm Column Name
const Object4IDGormColumnName = "id"

// Object4 UpdatedAt Gorm Column Name
const Object4UpdatedAtGormColumnName = "updated_at"

// Object4 CreatedAt JSON Key Name
const Object4CreatedAtJSONKeyName = "createdAt"

// Object4 DeletedAt JSON Key Name
const Object4DeletedAtJSONKeyName = "deletedAt"

// Object4 ID JSON Key Name
const Object4IDJSONKeyName = "id"

// Object4 Name JSON Key Name
const Object4NameJSONKeyName = "name"

// Object4 UpdatedAt JSON Key Name
const Object4UpdatedAtJSONKeyName = "updatedAt"

// Transform Object4 Gorm Column To JSON Key
func TransformObject4GormColumnToJSONKey(gormColumn string) (string, error) {
	switch gormColumn {
	case Object4CreatedAtGormColumnName:
		return Object4CreatedAtJSONKeyName, nil
	case Object4DeletedAtGormColumnName:
		return Object4DeletedAtJSONKeyName, nil
	case Object4IDGormColumnName:
		return Object4IDJSONKeyName, nil
	case Object4UpdatedAtGormColumnName:
		return Object4UpdatedAtJSONKeyName, nil
	default:
		return "", errors.WithStack(ErrObject4UnsupportedGormColumn)
	}
}

// Transform Object4 JSON Key To Gorm Column
func TransformObject4JSONKeyToGormColumn(jsonKey string) (string, error) {
	switch jsonKey {
	case Object4CreatedAtJSONKeyName:
		return Object4CreatedAtGormColumnName, nil
	case Object4DeletedAtJSONKeyName:
		return Object4DeletedAtGormColumnName, nil
	case Object4IDJSONKeyName:
		return Object4IDGormColumnName, nil
	case Object4UpdatedAtJSONKeyName:
		return Object4UpdatedAtGormColumnName, nil
	default:
		return "", errors.WithStack(ErrObject4UnsupportedJSONKey)
	}
}

// Transform Object4 JSON Key map To Gorm Column map
func TransformObject4JSONKeyMapToGormColumnMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject4JSONKeyToGormColumn(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject4UnsupportedJSONKey) {
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

// Transform Object4 Gorm Column map To JSON Key map
func TransformObject4GormColumnMapToJSONKeyMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject4GormColumnToJSONKey(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject4UnsupportedGormColumn) {
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
`,
		},
		{
			name: "simple with json omitempty",
			args: args{
				pkgName: "fake",
				obj:     Object5{},
			},
			want: `// Code generated by ModelTags. DO NOT EDIT.
package fake

import "emperror.dev/errors"

// ErrObject5UnsupportedGormColumn will be thrown when an unsupported Gorm column will be found in transform function.
var ErrObject5UnsupportedGormColumn = errors.Sentinel("unsupported gorm column")

// ErrObject5UnsupportedJSONKey will be thrown when an unsupported JSON key will be found in transform function.
var ErrObject5UnsupportedJSONKey = errors.Sentinel("unsupported json key")

// Object5 CreatedAt Gorm Column Name
const Object5CreatedAtGormColumnName = "created_at"

// Object5 DeletedAt Gorm Column Name
const Object5DeletedAtGormColumnName = "deleted_at"

// Object5 ID Gorm Column Name
const Object5IDGormColumnName = "id"

// Object5 Name Gorm Column Name
const Object5NameGormColumnName = "name"

// Object5 UpdatedAt Gorm Column Name
const Object5UpdatedAtGormColumnName = "updated_at"

// Object5 CreatedAt JSON Key Name
const Object5CreatedAtJSONKeyName = "createdAt"

// Object5 DeletedAt JSON Key Name
const Object5DeletedAtJSONKeyName = "deletedAt"

// Object5 ID JSON Key Name
const Object5IDJSONKeyName = "id"

// Object5 Name JSON Key Name
const Object5NameJSONKeyName = "name"

// Object5 UpdatedAt JSON Key Name
const Object5UpdatedAtJSONKeyName = "updatedAt"

// Transform Object5 Gorm Column To JSON Key
func TransformObject5GormColumnToJSONKey(gormColumn string) (string, error) {
	switch gormColumn {
	case Object5CreatedAtGormColumnName:
		return Object5CreatedAtJSONKeyName, nil
	case Object5DeletedAtGormColumnName:
		return Object5DeletedAtJSONKeyName, nil
	case Object5IDGormColumnName:
		return Object5IDJSONKeyName, nil
	case Object5NameGormColumnName:
		return Object5NameJSONKeyName, nil
	case Object5UpdatedAtGormColumnName:
		return Object5UpdatedAtJSONKeyName, nil
	default:
		return "", errors.WithStack(ErrObject5UnsupportedGormColumn)
	}
}

// Transform Object5 JSON Key To Gorm Column
func TransformObject5JSONKeyToGormColumn(jsonKey string) (string, error) {
	switch jsonKey {
	case Object5CreatedAtJSONKeyName:
		return Object5CreatedAtGormColumnName, nil
	case Object5DeletedAtJSONKeyName:
		return Object5DeletedAtGormColumnName, nil
	case Object5IDJSONKeyName:
		return Object5IDGormColumnName, nil
	case Object5NameJSONKeyName:
		return Object5NameGormColumnName, nil
	case Object5UpdatedAtJSONKeyName:
		return Object5UpdatedAtGormColumnName, nil
	default:
		return "", errors.WithStack(ErrObject5UnsupportedJSONKey)
	}
}

// Transform Object5 JSON Key map To Gorm Column map
func TransformObject5JSONKeyMapToGormColumnMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject5JSONKeyToGormColumn(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject5UnsupportedJSONKey) {
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

// Transform Object5 Gorm Column map To JSON Key map
func TransformObject5GormColumnMapToJSONKeyMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := TransformObject5GormColumnToJSONKey(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, ErrObject5UnsupportedGormColumn) {
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
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.pkgName, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got.String())
		})
	}
}
