// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// UnfamiliarLocationRiskEvent undocumented
type UnfamiliarLocationRiskEvent struct {
	// LocatedRiskEvent is the base model of UnfamiliarLocationRiskEvent
	LocatedRiskEvent
}

// UnfamiliarLocationRiskEventRequestBuilder is request builder for UnfamiliarLocationRiskEvent
type UnfamiliarLocationRiskEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnfamiliarLocationRiskEventRequest
func (b *UnfamiliarLocationRiskEventRequestBuilder) Request() *UnfamiliarLocationRiskEventRequest {
	return &UnfamiliarLocationRiskEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnfamiliarLocationRiskEventRequest is request for UnfamiliarLocationRiskEvent
type UnfamiliarLocationRiskEventRequest struct{ BaseRequest }

// Get performs GET request for UnfamiliarLocationRiskEvent
func (r *UnfamiliarLocationRiskEventRequest) Get(ctx context.Context) (resObj *UnfamiliarLocationRiskEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnfamiliarLocationRiskEvent
func (r *UnfamiliarLocationRiskEventRequest) Update(ctx context.Context, reqObj *UnfamiliarLocationRiskEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnfamiliarLocationRiskEvent
func (r *UnfamiliarLocationRiskEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
