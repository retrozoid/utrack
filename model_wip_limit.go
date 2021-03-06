/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents WIP limits for particular column. If they are not satisfied, the column will be highlighted in UI.
type WipLimit struct {
	Max    int32        `json:"max,omitempty"`
	Min    int32        `json:"min,omitempty"`
	Column *AgileColumn `json:"column,omitempty"`
	Id     string       `json:"id,omitempty"`
	Type_  string       `json:"$type,omitempty"`
}
