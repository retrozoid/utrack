/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents an error that appeared during the backup.
type BackupError struct {
	Date         int64  `json:"date,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}