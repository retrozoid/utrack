/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents a change in the list of tags of an issue.
type TagsActivityItem struct {
	Target  *Issue     `json:"target,omitempty"`
	Removed []IssueTag `json:"removed,omitempty"`
	Added   []IssueTag `json:"added,omitempty"`
}
