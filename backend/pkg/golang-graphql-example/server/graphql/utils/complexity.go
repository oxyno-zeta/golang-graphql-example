package utils

//
// All calculation in here are based on: https://shopify.engineering/rate-limiting-graphql-apis-calculating-query-complexity
//

const (
	baseMutationComplexity          = 10
	baseQuerySimpleStructComplexity = 1
	baseQueryConnectionComplexity   = 2
)

// CalculateMutationComplexity will calculate a mutation complexity.
func CalculateMutationComplexity(childComplexity int) int {
	return baseMutationComplexity + childComplexity
}

// CalculateQuerySimpleStructComplexity will calculate a query simple structure complexity.
func CalculateQuerySimpleStructComplexity(childComplexity int) int {
	return childComplexity + baseQuerySimpleStructComplexity
}

// CalculateQueryConnectionComplexity will calculate a query connection complexity.
func CalculateQueryConnectionComplexity(
	childComplexity int,
	after *string,
	before *string,
	first *int,
	last *int,
) int {
	// Initialize size
	var size int

	// Try to get the size of asked connection

	// Check if before and last are set
	// Or just check if first is set
	if before != nil && *before != "" && last != nil {
		size = *last
	} else if first != nil {
		size = *first
	}

	// Check if size is size is empty
	if size == 0 {
		size = defaultDefaultPageSize
	}

	// Result
	return childComplexity*size + baseQueryConnectionComplexity
}
