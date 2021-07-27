/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the change of the boolean flag that indicates whether YouTrack Wiki or Markdown is used in the target entity as a markup language. If \"true\", then the markdown is used. Otherwise, YouTrack Wiki markup.
type UsesMarkupActivityItem struct {
	Removed bool   `json:"removed,omitempty"`
	Added   bool   `json:"added,omitempty"`
	Markup  string `json:"markup,omitempty"`
}
