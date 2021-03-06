/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Default settings for the build-type field.
type BuildBundleCustomFieldDefaults struct {
	Bundle        *BuildBundle         `json:"bundle,omitempty"`
	DefaultValues []BuildBundleElement `json:"defaultValues,omitempty"`
}
