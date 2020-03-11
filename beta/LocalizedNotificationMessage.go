// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// LocalizedNotificationMessage The text content of a Notification Message Template for the specified locale.
type LocalizedNotificationMessage struct {
	// Entity is the base model of LocalizedNotificationMessage
	Entity
	// LastModifiedDateTime DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Locale The Locale for which this message is destined.
	Locale *string `json:"locale,omitempty"`
	// Subject The Message Template Subject.
	Subject *string `json:"subject,omitempty"`
	// MessageTemplate The Message Template content.
	MessageTemplate *string `json:"messageTemplate,omitempty"`
	// IsDefault Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
	IsDefault *bool `json:"isDefault,omitempty"`
}

// LocalizedNotificationMessageRequestBuilder is request builder for LocalizedNotificationMessage
type LocalizedNotificationMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns LocalizedNotificationMessageRequest
func (b *LocalizedNotificationMessageRequestBuilder) Request() *LocalizedNotificationMessageRequest {
	return &LocalizedNotificationMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LocalizedNotificationMessageRequest is request for LocalizedNotificationMessage
type LocalizedNotificationMessageRequest struct{ BaseRequest }

// Get performs GET request for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) Get(ctx context.Context) (resObj *LocalizedNotificationMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) Update(ctx context.Context, reqObj *LocalizedNotificationMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LocalizedNotificationMessage
func (r *LocalizedNotificationMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
