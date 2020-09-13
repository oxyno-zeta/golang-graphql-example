package pagination

type GenericFilter struct {
	Eq            interface{}
	NotEq         interface{}
	Gte           interface{}
	NotGte        interface{}
	Gt            interface{}
	NotGt         interface{}
	Lte           interface{}
	NotLte        interface{}
	Lt            interface{}
	NotLt         interface{}
	Contains      interface{}
	NotContains   interface{}
	StartsWith    interface{}
	NotStartsWith interface{}
	EndsWith      interface{}
	NotEndsWith   interface{}
	In            interface{}
	NotIn         interface{}
}
