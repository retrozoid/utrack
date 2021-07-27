/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents list of command and related comment in YouTrack. Can be used to either apply a command or get command suggestions.
type CommandList struct {
	Comment      string             `json:"comment,omitempty"`
	Visibility   *CommandVisibility `json:"visibility,omitempty"`
	Query        string             `json:"query,omitempty"`
	Caret        int32              `json:"caret,omitempty"`
	Silent       bool               `json:"silent,omitempty"`
	UsesMarkdown bool               `json:"usesMarkdown,omitempty"`
	RunAs        string             `json:"runAs,omitempty"`
	Commands     []ParsedCommand    `json:"commands,omitempty"`
	Issues       []Issue            `json:"issues,omitempty"`
	Suggestions  []Suggestion       `json:"suggestions,omitempty"`
	Id           string             `json:"id,omitempty"`
	Type_        string             `json:"$type,omitempty"`
}
