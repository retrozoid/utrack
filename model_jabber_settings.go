/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents jabber settings for this YouTrack installation.
type JabberSettings struct {
	IsEnabled   bool   `json:"isEnabled,omitempty"`
	Host        string `json:"host,omitempty"`
	Port        int32  `json:"port,omitempty"`
	Login       string `json:"login,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
	Id          string `json:"id,omitempty"`
	Type_       string `json:"$type,omitempty"`
}
