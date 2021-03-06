/*
 * YouTrack REST API
 *
 * YouTrack issue tracking and project management system
 *
 * API version: 2021.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package utrack

// Represents the currently logged in user in YouTrack.
type Me struct {
	Login             string        `json:"login,omitempty"`
	FullName          string        `json:"fullName,omitempty"`
	Email             string        `json:"email,omitempty"`
	JabberAccountName string        `json:"jabberAccountName,omitempty"`
	RingId            string        `json:"ringId,omitempty"`
	Guest             bool          `json:"guest,omitempty"`
	Online            bool          `json:"online,omitempty"`
	Banned            bool          `json:"banned,omitempty"`
	Tags              []IssueTag    `json:"tags,omitempty"`
	SavedQueries      []SavedQuery  `json:"savedQueries,omitempty"`
	AvatarUrl         string        `json:"avatarUrl,omitempty"`
	Profiles          *UserProfiles `json:"profiles,omitempty"`
	Id                string        `json:"id,omitempty"`
	Type_             string        `json:"$type,omitempty"`
}
