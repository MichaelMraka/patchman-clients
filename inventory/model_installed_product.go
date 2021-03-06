/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// InstalledProduct Representation of one installed product
type InstalledProduct struct {
	// the product ID
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// subscription status for product
	Status string `json:"status,omitempty"`
}
