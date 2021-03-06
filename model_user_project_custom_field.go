/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents project settings for the user field.
type UserProjectCustomField struct {
	Bundle        *UserBundle `json:"bundle,omitempty"`
	DefaultValues []User      `json:"defaultValues,omitempty"`
}
