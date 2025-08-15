package common

import (
	"reflect"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/errors"
)

// GenericFilter is a structure that will handle filters other than Date.
// This must be used as a pointer in other structures to be used automatically in filters.
// Moreover, a tag containing the database field must be declared.
// Example:
//
//	type Filter struct {
//	 AND []*Filter
//	 OR []*Filter
//	 Field1 *GenericFilter `dbfield:"field_1"`
//	}
//
// .
type GenericFilter struct {
	// Allow to test equality to
	Eq any `mapstructure:"eq"`
	// Allow to test non equality to
	NotEq any `mapstructure:"notEq"`
	// Allow to test greater or equal than
	Gte any `mapstructure:"gte"`
	// Allow to test not greater or equal than
	NotGte any `mapstructure:"notGte"`
	// Allow to test greater than
	Gt any `mapstructure:"gt"`
	// Allow to test not greater than
	NotGt any `mapstructure:"notGt"`
	// Allow to test less or equal than
	Lte any `mapstructure:"lte"`
	// Allow to test not less or equal than
	NotLte any `mapstructure:"notLte"`
	// Allow to test less than
	Lt any `mapstructure:"lt"`
	// Allow to test not less than
	NotLt any `mapstructure:"notLt"`
	// Allow to test if a string contains another string.
	// Contains must be a string
	Contains any `mapstructure:"contains"`
	// Allow to test if a string isn't containing another string.
	// NotContains must be a string
	NotContains any `mapstructure:"notContains"`
	// Allow to test if a string starts with another string.
	// StartsWith with must be a string
	StartsWith any `mapstructure:"startsWith"`
	// Allow to test if a string isn't starting with another string.
	// NotStartsWith must be a string
	NotStartsWith any `mapstructure:"notStartsWith"`
	// Allow to test if a string ends with another string.
	// EndsWith with must be a string
	EndsWith any `mapstructure:"endsWith"`
	// Allow to test if a string isn't ending with another string.
	// NotEndsWith must be a string
	NotEndsWith any `mapstructure:"notEndsWith"`
	// Allow to test if value is in array
	In any `mapstructure:"in"`
	// Allow to test if value isn't in array
	NotIn any `mapstructure:"notIn"`
	// Allow to test if value is null
	IsNull bool `mapstructure:"isNull"`
	// Allow to test if value is not null
	IsNotNull bool `mapstructure:"isNotNull"`
	// Allow to apply "upper()" function on field
	FieldUppercase bool `mapstructure:"fieldUppercase"`
	// Allow to apply "lower()" function on field
	FieldLowercase bool `mapstructure:"fieldLowercase"`
	// Allow to apply "upper()" function on values
	ValueUppercase bool `mapstructure:"valueUppercase"`
	// Allow to apply "lower()" function on values
	ValueLowercase bool `mapstructure:"valueLowercase"`
	// Allow case insensitive search.
	// That will automatically set FieldLowercase and ValueLowercase fields and generate correct SQL query.
	CaseInsensitive bool `mapstructure:"caseInsensitive"`
}

// DateFilter is a structure that will handle filters for dates.
// This must be used as a pointer in other structures to be used automatically in filters.
// Moreover, a tag containing the database field must be declared.
// Example:
//
//	type Filter struct {
//	 AND []*Filter
//	 OR []*Filter
//	 Field1 *DateFilter `dbfield:"field_1"`
//	}
//
// .
type DateFilter struct {
	// Allow to test equality to
	Eq any `mapstructure:"eq"`
	// Allow to test non equality to
	NotEq any `mapstructure:"notEq"`
	// Allow to test greater or equal than
	Gte any `mapstructure:"gte"`
	// Allow to test not greater or equal than
	NotGte any `mapstructure:"notGte"`
	// Allow to test greater than
	Gt any `mapstructure:"gt"`
	// Allow to test not greater than
	NotGt any `mapstructure:"notGt"`
	// Allow to test less or equal than
	Lte any `mapstructure:"lte"`
	// Allow to test not less or equal than
	NotLte any `mapstructure:"notLte"`
	// Allow to test less than
	Lt any `mapstructure:"lt"`
	// Allow to test not less than
	NotLt any `mapstructure:"notLt"`
	// Allow to test if value is in array
	In any `mapstructure:"in"`
	// Allow to test if value isn't in array
	NotIn any `mapstructure:"notIn"`
	// Allow to test if value is null
	IsNull bool `mapstructure:"isNull"`
	// Allow to test if value is not null
	IsNotNull bool `mapstructure:"isNotNull"`
}

// GenericFilterBuilder is an interface that must be implemented in order to work automatic filter.
// This is done like this in order to add more fields in GenericFilter without the need of upgrading
// all code in other to be compatible.
type GenericFilterBuilder interface {
	GetGenericFilter() (*GenericFilter, error)
}

func (g *GenericFilter) GetGenericFilter() (*GenericFilter, error) { return g, nil }

func (d *DateFilter) GetGenericFilter() (*GenericFilter, error) {
	// Create result
	res := &GenericFilter{}

	// Eq case
	if reflect.ValueOf(d.Eq).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.Eq)
		// Check error
		if err != nil {
			return nil, err
		}

		res.Eq = t
	}

	// Not Eq case
	if reflect.ValueOf(d.NotEq).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.NotEq)
		// Check error
		if err != nil {
			return nil, err
		}

		res.NotEq = t
	}

	// Gte case
	if reflect.ValueOf(d.Gte).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.Gte)
		// Check error
		if err != nil {
			return nil, err
		}

		res.Gte = t
	}

	// Not Gte case
	if reflect.ValueOf(d.NotGte).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.NotGte)
		// Check error
		if err != nil {
			return nil, err
		}

		res.NotGte = t
	}

	// Gt case
	if reflect.ValueOf(d.Gt).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.Gt)
		// Check error
		if err != nil {
			return nil, err
		}

		res.Gt = t
	}

	// Not Gt case
	if reflect.ValueOf(d.NotGt).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.NotGt)
		// Check error
		if err != nil {
			return nil, err
		}

		res.NotGt = t
	}

	// Lte case
	if reflect.ValueOf(d.Lte).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.Lte)
		// Check error
		if err != nil {
			return nil, err
		}

		res.Lte = t
	}

	// Not Lte case
	if reflect.ValueOf(d.NotLte).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.NotLte)
		// Check error
		if err != nil {
			return nil, err
		}

		res.NotLte = t
	}

	// Lt case
	if reflect.ValueOf(d.Lt).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.Lt)
		// Check error
		if err != nil {
			return nil, err
		}

		res.Lt = t
	}

	// Not Lt case
	if reflect.ValueOf(d.NotLt).IsValid() {
		// Parse time
		t, err := parseOrGetTime(d.NotLt)
		// Check error
		if err != nil {
			return nil, err
		}

		res.NotLt = t
	}

	// In case
	if reflect.ValueOf(d.In).IsValid() {
		// Parse time
		t, err := parseOrGetTimes(d.In)
		// Check error
		if err != nil {
			return nil, err
		}

		res.In = t
	}

	// Not In case
	if reflect.ValueOf(d.NotIn).IsValid() {
		// Parse time
		t, err := parseOrGetTimes(d.NotIn)
		// Check error
		if err != nil {
			return nil, err
		}

		res.NotIn = t
	}

	// Apply is null
	res.IsNull = d.IsNull
	// Apply is not null
	res.IsNotNull = d.IsNotNull

	return res, nil
}

func parseOrGetTime(x any) (*time.Time, error) {
	// Get value in reflect mode
	val := reflect.Indirect(reflect.ValueOf(x))
	// Get interface value
	valInt := val.Interface()

	// Switch on type
	switch v := valInt.(type) {
	case string:
		// Parse date
		t, err := time.Parse(time.RFC3339Nano, v)
		// Check error
		if err != nil {
			// In this particular case, display error in public message in order to help api user to detect the error
			// and consider that error as an invalid input error.
			return nil, errors.NewInvalidInputErrorWithError(err, errors.WithPublicError(err))
		}

		// Force utc
		t = t.UTC()

		return &t, nil
	case time.Time:
		t := v
		// Force utc
		t = t.UTC()

		return &t, nil
	default:
		return nil, errors.NewInternalServerError("date filter value not supported")
	}
}

func parseOrGetTimes(x any) ([]*time.Time, error) {
	// Prepare result
	res := make([]*time.Time, 0)

	// Get value in reflect mode
	val := reflect.Indirect(reflect.ValueOf(x))

	// Check that value is an array
	if val.Kind() != reflect.Slice {
		return nil, errors.NewInternalServerError("date filter input must be a slice")
	}

	// Loop over all values
	for i := 0; i < val.Len(); i++ {
		// Get value
		v := val.Index(i).Interface()
		// Parse time
		t, err := parseOrGetTime(v)
		// Check error
		if err != nil {
			return nil, err
		}
		// Append
		res = append(res, t)
	}

	return res, nil
}
