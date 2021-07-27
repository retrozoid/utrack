/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// A `WatchFolder` is an `IssueFolder` that let you enable notifications for a set  of issues that it enfolds. It is a common abstract ancestor for saved searches and issue tags.
type WatchFolder struct {
	Name         string     `json:"name,omitempty"`
	Id           string     `json:"id,omitempty"`
	Type_        string     `json:"$type,omitempty"`
	Owner        *User      `json:"owner,omitempty"`
	VisibleFor   *UserGroup `json:"visibleFor,omitempty"`
	UpdateableBy *UserGroup `json:"updateableBy,omitempty"`
}
