/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents a change in the `author` attribute of a work item.
type WorkItemAuthorActivityItem struct {
	Target  *IssueWorkItem `json:"target,omitempty"`
	Removed *User          `json:"removed,omitempty"`
	Added   *User          `json:"added,omitempty"`
}
