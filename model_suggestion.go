/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents query suggestion.
type Suggestion struct {
	CompletionStart int32  `json:"completionStart,omitempty"`
	CompletionEnd   int32  `json:"completionEnd,omitempty"`
	MatchingStart   int32  `json:"matchingStart,omitempty"`
	MatchingEnd     int32  `json:"matchingEnd,omitempty"`
	Caret           int32  `json:"caret,omitempty"`
	Description     string `json:"description,omitempty"`
	Option          string `json:"option,omitempty"`
	Prefix          string `json:"prefix,omitempty"`
	Suffix          string `json:"suffix,omitempty"`
	Group           string `json:"group,omitempty"`
	Icon            string `json:"icon,omitempty"`
	AuxiliaryIcon   string `json:"auxiliaryIcon,omitempty"`
	ClassName       string `json:"className,omitempty"`
	Id              string `json:"id,omitempty"`
	Type_           string `json:"$type,omitempty"`
}