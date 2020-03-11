// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AadUserConversationMember undocumented
type AadUserConversationMember struct {
	// ConversationMember is the base model of AadUserConversationMember
	ConversationMember
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// User undocumented
	User *User `json:"user,omitempty"`
}

// AadUserConversationMemberRequestBuilder is request builder for AadUserConversationMember
type AadUserConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns AadUserConversationMemberRequest
func (b *AadUserConversationMemberRequestBuilder) Request() *AadUserConversationMemberRequest {
	return &AadUserConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AadUserConversationMemberRequest is request for AadUserConversationMember
type AadUserConversationMemberRequest struct{ BaseRequest }

// Get performs GET request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Get(ctx context.Context) (resObj *AadUserConversationMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Update(ctx context.Context, reqObj *AadUserConversationMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// User is navigation property
func (b *AadUserConversationMemberRequestBuilder) User() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/user"
	return bb
}
