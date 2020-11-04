/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// ActiveTag Information about a host tag
type ActiveTag struct {
	// The number of hosts with the given tag. If the value is null this indicates that the count is unknown.
	Count *int32 `json:"count"`
	Tag StructuredTag `json:"tag"`
}
