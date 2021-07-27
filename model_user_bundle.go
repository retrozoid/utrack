/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents a set of values that contains users. You can add to the set both individual user accounts and groups of users.
type UserBundle struct {
	IsUpdateable    bool        `json:"isUpdateable,omitempty"`
	Id              string      `json:"id,omitempty"`
	Type_           string      `json:"$type,omitempty"`
	Groups          []UserGroup `json:"groups,omitempty"`
	Individuals     []User      `json:"individuals,omitempty"`
	AggregatedUsers []User      `json:"aggregatedUsers,omitempty"`
}