package utils

// Pagination information.
type PageInfo struct {
	// Shortcut to first edge cursor in the result chunk.
	StartCursor *string `json:"startCursor"`
	// Shortcut to last edge cursor in the result chunk.
	EndCursor *string `json:"endCursor"`
	// Has a next page ?
	HasNextPage bool `json:"hasNextPage"`
	// Has a previous page ?
	HasPreviousPage bool `json:"hasPreviousPage"`
}
