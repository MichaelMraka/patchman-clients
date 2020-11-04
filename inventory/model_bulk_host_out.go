/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// BulkHostOut struct for BulkHostOut
type BulkHostOut struct {
	// List of hosts that were created, updated or failed to be processed.
	Data []BulkHostOutDetails `json:"data,omitempty"`
	// Number of items in the \"data\" list that contain errors.
	Errors int32 `json:"errors,omitempty"`
	// Total number of items in the \"data\" list.
	Total int32 `json:"total,omitempty"`
}
