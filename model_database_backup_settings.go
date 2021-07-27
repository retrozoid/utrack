/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents database backup settings of the YouTrack instance.
type DatabaseBackupSettings struct {
	Location           string        `json:"location,omitempty"`
	FilesToKeep        int32         `json:"filesToKeep,omitempty"`
	CronExpression     string        `json:"cronExpression,omitempty"`
	ArchiveFormat      string        `json:"archiveFormat,omitempty"`
	IsOn               bool          `json:"isOn,omitempty"`
	AvailableDiskSpace int64         `json:"availableDiskSpace,omitempty"`
	NotifiedUsers      []User        `json:"notifiedUsers,omitempty"`
	BackupStatus       *BackupStatus `json:"backupStatus,omitempty"`
	Id                 string        `json:"id,omitempty"`
	Type_              string        `json:"$type,omitempty"`
}