/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// CreateCheckIn Data required to create a check-in record for a host.
type CreateCheckIn struct {
	// A set of string facts about a host.
	CanonicalFacts map[string]interface{} `json:"canonical_facts"`
	// Defines how far in the future the host becomes stale (in minutes).
	CheckinFrequency int32 `json:"checkin_frequency,omitempty"`
}
