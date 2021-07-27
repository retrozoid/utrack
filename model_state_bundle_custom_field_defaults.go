/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Default settings for the state-type field.
type StateBundleCustomFieldDefaults struct {
	Bundle        *StateBundle         `json:"bundle,omitempty"`
	DefaultValues []StateBundleElement `json:"defaultValues,omitempty"`
}