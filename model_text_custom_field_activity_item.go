/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents an activity that affects a custom field of the `text` type of an issue.
type TextCustomFieldActivityItem struct {
	Target  *Issue `json:"target,omitempty"`
	Removed string `json:"removed,omitempty"`
	Added   string `json:"added,omitempty"`
	Markup  string `json:"markup,omitempty"`
}
