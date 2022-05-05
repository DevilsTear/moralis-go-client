/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type NftTransferCollection struct {
	// The total number of matches for this query
	Total int32 `json:"total"`
	// The page of the current result
	Page int32 `json:"page"`
	// The number of results per page
	PageSize int32 `json:"page_size"`
	// The cursor to get to the next page
	Cursor string        `json:"cursor"`
	Result []NftTransfer `json:"result"`
	// Indicator if the block exists
	BlockExists bool `json:"block_exists,omitempty"`
	// Indicator if the block is fully indexed
	IndexComplete bool `json:"index_complete,omitempty"`
}
