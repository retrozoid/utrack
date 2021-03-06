/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents a custom field of the issue.
type CustomFilterField struct {
	Name        string       `json:"name,omitempty"`
	Id          string       `json:"id,omitempty"`
	Type_       string       `json:"$type,omitempty"`
	CustomField *CustomField `json:"customField,omitempty"`
}
