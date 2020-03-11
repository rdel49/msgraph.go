// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ColumnLink undocumented
type ColumnLink struct {
	// Entity is the base model of ColumnLink
	Entity
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// ColumnLinkRequestBuilder is request builder for ColumnLink
type ColumnLinkRequestBuilder struct{ BaseRequestBuilder }

// Request returns ColumnLinkRequest
func (b *ColumnLinkRequestBuilder) Request() *ColumnLinkRequest {
	return &ColumnLinkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ColumnLinkRequest is request for ColumnLink
type ColumnLinkRequest struct{ BaseRequest }

// Get performs GET request for ColumnLink
func (r *ColumnLinkRequest) Get(ctx context.Context) (resObj *ColumnLink, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ColumnLink
func (r *ColumnLinkRequest) Update(ctx context.Context, reqObj *ColumnLink) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ColumnLink
func (r *ColumnLinkRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
