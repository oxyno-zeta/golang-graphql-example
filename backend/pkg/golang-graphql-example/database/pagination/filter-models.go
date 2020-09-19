package pagination

// GenericFilter is a structure that will handle filters.
// This must be used as a pointer in other structures to be used automatically in filters.
// Moreover, a tag containing the database field must be declared.
// Example:
// type Filter struct {
// 	Field1 *GenericFilter `dbfield:"field_1"`
// }
// .
type GenericFilter struct {
	// Allow to test equality to
	Eq interface{}
	// Allow to test non equality to
	NotEq interface{}
	// Allow to test greater or equal than
	Gte interface{}
	// Allow to test not greater or equal than
	NotGte interface{}
	// Allow to test greater than
	Gt interface{}
	// Allow to test not greater than
	NotGt interface{}
	// Allow to test less or equal than
	Lte interface{}
	// Allow to test not less or equal than
	NotLte interface{}
	// Allow to test less than
	Lt interface{}
	// Allow to test not less than
	NotLt interface{}
	// Allow to test if a string contains another string.
	// Contains must be a string
	Contains interface{}
	// Allow to test if a string isn't containing another string.
	// NotContains must be a string
	NotContains interface{}
	// Allow to test if a string starts with another string.
	// StartsWith with must be a string
	StartsWith interface{}
	// Allow to test if a string isn't starting with another string.
	// NotStartsWith must be a string
	NotStartsWith interface{}
	// Allow to test if a string ends with another string.
	// EndsWith with must be a string
	EndsWith interface{}
	// Allow to test if a string isn't ending with another string.
	// NotEndsWith must be a string
	NotEndsWith interface{}
	// Allow to test if value is in array
	In interface{}
	// Allow to test if value isn't in array
	NotIn interface{}
}
