// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WebAccount undocumented
type WebAccount struct {
	// ItemFacet is the base model of WebAccount
	ItemFacet
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// Service undocumented
	Service *ServiceInformation `json:"service,omitempty"`
	// StatusMessage undocumented
	StatusMessage *string `json:"statusMessage,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

// WebAccountRequestBuilder is request builder for WebAccount
type WebAccountRequestBuilder struct{ BaseRequestBuilder }

// Request returns WebAccountRequest
func (b *WebAccountRequestBuilder) Request() *WebAccountRequest {
	return &WebAccountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WebAccountRequest is request for WebAccount
type WebAccountRequest struct{ BaseRequest }

// Get performs GET request for WebAccount
func (r *WebAccountRequest) Get(ctx context.Context) (resObj *WebAccount, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WebAccount
func (r *WebAccountRequest) Update(ctx context.Context, reqObj *WebAccount) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WebAccount
func (r *WebAccountRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
