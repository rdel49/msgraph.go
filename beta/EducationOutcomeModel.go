// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationOutcome undocumented
type EducationOutcome struct {
	Entity
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}