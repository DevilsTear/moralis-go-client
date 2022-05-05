/*
 * Moralis API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type NftOwnerCollection struct {
	// The syncing status of the address [SYNCING/SYNCED]
	Status string `json:"status,omitempty"`
	// The total number of matches for this query
	Total int32 `json:"total,omitempty"`
	// The page of the current result
	Page int32 `json:"page,omitempty"`
	// The number of results per page
	PageSize int32 `json:"page_size,omitempty"`
	// The cursor to get to the next page
	Cursor string     `json:"cursor,omitempty"`
	Result []NftOwner `json:"result,omitempty"`
}
