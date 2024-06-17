package gen

import (
	"bytes"
	"reflect"
	"strings"
	"sync"
	"text/template"

	"gorm.io/gorm/schema"
)

const tmplStr = `// Code generated by ModelTags. DO NOT EDIT.
package {{ .PkgName }}

import "emperror.dev/errors"

// Err{{ $.ObjectName }}UnsupportedGormColumn will be thrown when an unsupported Gorm column will be found in transform function.
var Err{{ $.ObjectName }}UnsupportedGormColumn = errors.Sentinel("unsupported gorm column")

// Err{{ $.ObjectName }}UnsupportedJSONKey will be thrown when an unsupported JSON key will be found in transform function.
var Err{{ $.ObjectName }}UnsupportedJSONKey = errors.Sentinel("unsupported json key")
{{ range $key, $value := .GormMap }}
// {{ $.ObjectName }} {{ $key }} Gorm Column Name
const {{ $.ObjectName }}{{ $key }}GormColumnName = "{{ $value }}"
{{ end -}}
{{ range $key, $value := .JSONMap }}
// {{ $.ObjectName }} {{ $key }} JSON Key Name
const {{ $.ObjectName }}{{ $key }}JSONKeyName = "{{ $value }}"
{{ end }}
// Transform {{ .ObjectName }} Gorm Column To JSON Key
func Transform{{ .ObjectName }}GormColumnToJSONKey(gormColumn string) (string, error) {
	switch gormColumn {
	{{ range $key, $value := .GormMap }}
	{{- if not (eq (index $.JSONMap $key) "") -}}
	case {{ $.ObjectName }}{{ $key }}GormColumnName:
		return {{ $.ObjectName }}{{ $key }}JSONKeyName, nil
	{{ end -}}
	{{ end -}}
	default:
		return "", errors.WithStack(Err{{ $.ObjectName }}UnsupportedGormColumn)
	}
}

// Transform {{ .ObjectName }} JSON Key To Gorm Column
func Transform{{ .ObjectName }}JSONKeyToGormColumn(jsonKey string) (string, error) {
	switch jsonKey {
	{{ range $key, $value := .JSONMap }}
	{{- if not (eq (index $.GormMap $key) "") -}}
	case {{ $.ObjectName }}{{ $key }}JSONKeyName:
		return {{ $.ObjectName }}{{ $key }}GormColumnName, nil
	{{ end -}}
	{{ end -}}
	default:
		return "", errors.WithStack(Err{{ $.ObjectName }}UnsupportedJSONKey)
	}
}

func Transform{{ .ObjectName }}JSONKeyMapToGormColumnMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := Transform{{ .ObjectName }}JSONKeyToGormColumn(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, Err{{ $.ObjectName }}UnsupportedJSONKey) {
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

func Transform{{ .ObjectName }}GormColumnMapToJSONKeyMap(
	input map[string]interface{},
	ignoreUnsupportedError bool,
) (map[string]interface{}, error) {
	// Rebuild
	m := map[string]interface{}{}
	// Loop over input
	for k, v := range input {
		r, err := Transform{{ .ObjectName }}GormColumnToJSONKey(k)
		// Check error
		if err != nil {
			// Check if ignore is enabled and error is matching
			if ignoreUnsupportedError && errors.Is(err, Err{{ $.ObjectName }}UnsupportedGormColumn) {
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
`

type MainPkg struct {
	GormMap    map[string]string
	JSONMap    map[string]string
	PkgName    string
	ObjectName string
}

func saveJSONKeys(rType reflect.Type, jsonMap map[string]string) {
	// Loop over object field
	for i := 0; i < rType.NumField(); i++ {
		// Get field
		fieldType := rType.Field(i)
		// Get tag value for field
		jsonTagValue := fieldType.Tag.Get("json")
		// Check if tag isn't set or marked to ignore
		if jsonTagValue == "-" {
			continue
		}

		// Check if it is anonymous
		if fieldType.Anonymous {
			// Need to go deeper
			saveJSONKeys(fieldType.Type, jsonMap)

			continue
		}

		// Check if it is empty
		if jsonTagValue == "" {
			jsonMap[rType.Field(i).Name] = rType.Field(i).Name
		} else {
			// Split over ,
			jsonTagValueSplit := strings.Split(jsonTagValue, ",")
			// Get first value as it is the key name
			jsonMap[rType.Field(i).Name] = jsonTagValueSplit[0]
		}
	}
}

func generatePackageData(pkgName string, obj interface{}) (*MainPkg, error) {
	// Manage Gorm
	// Parse object to get schema
	s, err := schema.Parse(obj, &sync.Map{}, schema.NamingStrategy{})
	// Check error
	if err != nil {
		return nil, err
	}

	// Init gorm map
	gormMap := make(map[string]string)
	// Loop over gorm fields
	for _, field := range s.Fields {
		dbName := field.DBName
		// Check if db field is set
		if dbName != "" {
			gormMap[field.Name] = dbName
		}
	}

	// Get reflect type of object
	rType := reflect.TypeOf(obj)

	// Init json map
	jsonMap := make(map[string]string)

	// Save json keys
	saveJSONKeys(rType, jsonMap)

	return &MainPkg{
		PkgName:    pkgName,
		ObjectName: rType.Name(),
		GormMap:    gormMap,
		JSONMap:    jsonMap,
	}, nil
}

func Generate(pkgName string, obj interface{}) (*bytes.Buffer, error) {
	// Init template
	tmpl, err := template.New("test").Parse(tmplStr)
	// Check error
	if err != nil {
		return nil, err
	}

	// Create buffer
	buf := &bytes.Buffer{}
	// Get main package data
	mainPkg, err := generatePackageData(pkgName, obj)
	// Check error
	if err != nil {
		return nil, err
	}

	// Execute template
	err = tmpl.Execute(buf, mainPkg)
	// Check error
	if err != nil {
		return nil, err
	}

	return buf, nil
}
