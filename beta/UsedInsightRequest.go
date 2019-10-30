// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// UsedInsightRequestBuilder is request builder for UsedInsight
type UsedInsightRequestBuilder struct{ BaseRequestBuilder }

// Request returns UsedInsightRequest
func (b *UsedInsightRequestBuilder) Request() *UsedInsightRequest {
	return &UsedInsightRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UsedInsightRequest is request for UsedInsight
type UsedInsightRequest struct{ BaseRequest }

// Get performs GET request for UsedInsight
func (r *UsedInsightRequest) Get(ctx context.Context) (resObj *UsedInsight, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UsedInsight
func (r *UsedInsightRequest) Update(ctx context.Context, reqObj *UsedInsight) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UsedInsight
func (r *UsedInsightRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Resource is navigation property
func (b *UsedInsightRequestBuilder) Resource() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}
