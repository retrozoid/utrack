/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents a build - a single element of a builds bundle.
type BuildBundleElement struct {
	Name          string      `json:"name,omitempty"`
	Bundle        *Bundle     `json:"bundle,omitempty"`
	Description   string      `json:"description,omitempty"`
	Ordinal       int32       `json:"ordinal,omitempty"`
	Color         *FieldStyle `json:"color,omitempty"`
	HasRunningJob bool        `json:"hasRunningJob,omitempty"`
	Id            string      `json:"id,omitempty"`
	Type_         string      `json:"$type,omitempty"`
	AssembleDate  int64       `json:"assembleDate,omitempty"`
}