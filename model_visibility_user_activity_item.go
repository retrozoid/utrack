/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the event when a user adds or removes a user to/from the Visibility settings of the target entity.
type VisibilityUserActivityItem struct {
	TargetMember    string `json:"targetMember,omitempty"`
	TargetSubMember string `json:"targetSubMember,omitempty"`
	Removed         []User `json:"removed,omitempty"`
	Added           []User `json:"added,omitempty"`
}