/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rbac
// PolicyPagination struct for PolicyPagination
type PolicyPagination struct {
	Meta PaginationMeta `json:"meta,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
	Data []PolicyExtended `json:"data"`
}