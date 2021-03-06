/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the Visual settings of the YouTrack service.
type AppearanceSettings struct {
	TimeZone        *TimeZoneDescriptor   `json:"timeZone,omitempty"`
	DateFieldFormat *DateFormatDescriptor `json:"dateFieldFormat,omitempty"`
	Logo            *Logo                 `json:"logo,omitempty"`
	Id              string                `json:"id,omitempty"`
	Type_           string                `json:"$type,omitempty"`
}
