// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeleteUserFromSharedAppleDeviceActionResult undocumented
type DeleteUserFromSharedAppleDeviceActionResult struct {
	DeviceActionResult
	// UserPrincipalName User principal name of the user to be deleted
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}