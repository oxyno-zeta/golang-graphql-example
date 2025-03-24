package config

import (
	"reflect"

	"emperror.dev/errors"
)

// Load credential configs here.
func (impl *managerimpl) loadAllCredentials(out *Config) ([]*CredentialConfig, error) {
	return internalLoadAllCredentials(out, impl.credentialConfigPathList)
}

func internalLoadAllCredentials(
	out interface{},
	credentialConfigPathList [][]string,
) ([]*CredentialConfig, error) {
	// Initialize answer
	result := make([]*CredentialConfig, 0)

	// Loop over credential config path list
	for _, cPath := range credentialConfigPathList {
		// Save path value
		pVal := reflect.ValueOf(out).Elem()
		// Save length
		pLen := len(cPath)

		// Loop over path
		for i, p := range cPath {
			// Get path value
			pVal = pVal.FieldByName(p)
			// Check if new value doesn't exist
			if !pVal.IsValid() || (pVal.Kind() == reflect.Pointer && pVal.IsNil()) {
				// Stop here
				break
			}

			// Check if it is a pointer and not the last key
			if pVal.Kind() == reflect.Pointer && i != pLen-1 {
				// Remove ptr
				pVal = pVal.Elem()
			}
		}

		// Check if value exists
		if pVal.IsValid() && !pVal.IsNil() {
			// Check if it is an array or slice
			if pVal.Kind() == reflect.Array || pVal.Kind() == reflect.Slice {
				// Loop over it
				for i := 0; i < pVal.Len(); i++ {
					// Get value
					pVal2 := pVal.Index(i)
					// Check if value exists
					if pVal2.IsValid() && !pVal2.IsNil() {
						// Get value
						v := pVal2.Interface()

						vv, err := internalLoadCredential(v)
						// Check error
						if err != nil {
							return nil, err
						}
						// Append result
						result = append(result, vv)
					}
				}
			} else {
				// Direct object case
				// Get value
				v := pVal.Interface()

				vv, err := internalLoadCredential(v)
				// Check error
				if err != nil {
					return nil, err
				}
				// Append result
				result = append(result, vv)
			}
		}
	}

	return result, nil
}

func internalLoadCredential(in interface{}) (*CredentialConfig, error) {
	// Cast
	vv, ok := in.(*CredentialConfig)
	// Check if cast is ok
	if !ok {
		return nil, errors.New("cannot cast to *CredentialConfig")
	}

	// Load database credential
	err := loadCredential(vv)
	if err != nil {
		return nil, err
	}

	return vv, nil
}

func getCredentialConfigPathList() ([][]string, error) {
	return getRecursivelyCredentialConfigPathList([]string{}, reflect.TypeOf(Config{}))
}

func getRecursivelyCredentialConfigPathList(keys []string, r reflect.Type) ([][]string, error) {
	// Init result
	res := [][]string{}

	// Create type to save it and use it
	credCfgType := reflect.TypeOf(CredentialConfig{})

	// Loop over fields
	for i := 0; i < r.NumField(); i++ {
		// Get field
		field := r.Field(i)
		// Get type
		fieldType := field.Type

		// Check if it is a pointer
		if fieldType.Kind() == reflect.Pointer {
			// Remove pointer
			fieldType = fieldType.Elem()
		}

		// Check if it is an array or slice
		if fieldType.Kind() == reflect.Array || fieldType.Kind() == reflect.Slice {
			// Get slice/array element
			fieldType = fieldType.Elem()
			// Check if it is a pointer
			if fieldType.Kind() == reflect.Pointer {
				// Remove pointer
				fieldType = fieldType.Elem()
			}
		}

		// Check if it isn't a struct or pointer
		if fieldType.Kind() != reflect.Struct {
			continue
		}

		// Get field name
		fieldName := field.Name

		// Check if it is a CredentialConfig
		// If yes, save path and continue
		if fieldType.AssignableTo(credCfgType) {
			// Force recreate a slice
			inter := append([]string{}, keys...)
			inter = append(inter, fieldName)
			// Save
			res = append(res, inter)

			continue
		}

		// Create new array by forcing array recreation
		inter := append([]string{}, keys...)
		// Analyze sub type
		intRes, err := getRecursivelyCredentialConfigPathList(
			append(inter, fieldName),
			fieldType,
		)
		// Check error
		if err != nil {
			return nil, err
		}
		// Check if intermediate result exists
		if len(intRes) != 0 {
			res = append(res, intRes...)
		}
	}

	return res, nil
}
