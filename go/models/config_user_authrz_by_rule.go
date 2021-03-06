package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// ConfigUserAuthrzByRule config user authrz by rule
// swagger:model ConfigUserAuthrzByRule
type ConfigUserAuthrzByRule struct {

	// Comma separated list of policies assigned to the user. Field introduced in 18.2.7.
	Policies *string `json:"policies,omitempty"`

	// assigned roles.
	Roles *string `json:"roles,omitempty"`

	// matching rule string.
	Rule *string `json:"rule,omitempty"`

	// assigned tenants.
	Tenants *string `json:"tenants,omitempty"`

	// Request user.
	User *string `json:"user,omitempty"`
}
