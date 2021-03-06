/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the issue custom field of the `user` type that can have multiple values.
type MultiUserIssueCustomField struct {
	ProjectCustomField *ProjectCustomField `json:"projectCustomField,omitempty"`
	Value              *User               `json:"value,omitempty"`
}
