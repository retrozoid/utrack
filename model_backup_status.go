/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the current status of the backup process.
type BackupStatus struct {
	BackupInProgress bool         `json:"backupInProgress,omitempty"`
	BackupCancelled  bool         `json:"backupCancelled,omitempty"`
	BackupError      *BackupError `json:"backupError,omitempty"`
	Id               string       `json:"id,omitempty"`
	Type_            string       `json:"$type,omitempty"`
}
