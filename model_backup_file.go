/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the backup file.
type BackupFile struct {
	Name         string `json:"name,omitempty"`
	Size         int64  `json:"size,omitempty"`
	CreationDate int64  `json:"creationDate,omitempty"`
	Link         string `json:"link,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}
