/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the issue custom field of the `ownedField` type that can only have a single value.
type SingleOwnedIssueCustomField struct {
	ProjectCustomField *ProjectCustomField `json:"projectCustomField,omitempty"`
	Value              *OwnedBundleElement `json:"value,omitempty"`
}
