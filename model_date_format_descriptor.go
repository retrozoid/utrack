/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents date format.
type DateFormatDescriptor struct {
	Presentation string `json:"presentation,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	DatePattern  string `json:"datePattern,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}
