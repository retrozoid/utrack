/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents visibility limited to several users and/or groups.
type LimitedVisibility struct {
	Id              string      `json:"id,omitempty"`
	Type_           string      `json:"$type,omitempty"`
	PermittedGroups []UserGroup `json:"permittedGroups,omitempty"`
	PermittedUsers  []User      `json:"permittedUsers,omitempty"`
}