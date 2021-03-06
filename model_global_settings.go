/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents application-wide settings.
type GlobalSettings struct {
	SystemSettings       *SystemSettings       `json:"systemSettings,omitempty"`
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
	RestSettings         *RestCorsSettings     `json:"restSettings,omitempty"`
	AppearanceSettings   *AppearanceSettings   `json:"appearanceSettings,omitempty"`
	LocaleSettings       *LocaleSettings       `json:"localeSettings,omitempty"`
	License              *License              `json:"license,omitempty"`
	Id                   string                `json:"id,omitempty"`
	Type_                string                `json:"$type,omitempty"`
}
