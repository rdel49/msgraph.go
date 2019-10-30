// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TimeOffReasonRequestBuilder is request builder for TimeOffReason
type TimeOffReasonRequestBuilder struct{ BaseRequestBuilder }

// Request returns TimeOffReasonRequest
func (b *TimeOffReasonRequestBuilder) Request() *TimeOffReasonRequest {
	return &TimeOffReasonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TimeOffReasonRequest is request for TimeOffReason
type TimeOffReasonRequest struct{ BaseRequest }

// Get performs GET request for TimeOffReason
func (r *TimeOffReasonRequest) Get(ctx context.Context) (resObj *TimeOffReason, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TimeOffReason
func (r *TimeOffReasonRequest) Update(ctx context.Context, reqObj *TimeOffReason) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TimeOffReason
func (r *TimeOffReasonRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
