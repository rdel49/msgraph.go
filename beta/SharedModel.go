// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Shared undocumented
type Shared struct {
	// Owner undocumented
	Owner *IdentitySet `json:"owner,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// SharedBy undocumented
	SharedBy *IdentitySet `json:"sharedBy,omitempty"`
	// SharedDateTime undocumented
	SharedDateTime *time.Time `json:"sharedDateTime,omitempty"`
}