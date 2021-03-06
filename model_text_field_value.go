/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents a value of the text field. Returns both source and rendered text.
type TextFieldValue struct {
	Text         string `json:"text,omitempty"`
	MarkdownText string `json:"markdownText,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}
