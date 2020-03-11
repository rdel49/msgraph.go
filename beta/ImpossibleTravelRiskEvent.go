// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// ImpossibleTravelRiskEvent undocumented
type ImpossibleTravelRiskEvent struct {
	// LocatedRiskEvent is the base model of ImpossibleTravelRiskEvent
	LocatedRiskEvent
	// UserAgent undocumented
	UserAgent *string `json:"userAgent,omitempty"`
	// DeviceInformation undocumented
	DeviceInformation *string `json:"deviceInformation,omitempty"`
	// IsAtypicalLocation undocumented
	IsAtypicalLocation *bool `json:"isAtypicalLocation,omitempty"`
	// PreviousSigninDateTime undocumented
	PreviousSigninDateTime *time.Time `json:"previousSigninDateTime,omitempty"`
	// PreviousLocation undocumented
	PreviousLocation *SignInLocation `json:"previousLocation,omitempty"`
	// PreviousIPAddress undocumented
	PreviousIPAddress *string `json:"previousIpAddress,omitempty"`
}

// ImpossibleTravelRiskEventRequestBuilder is request builder for ImpossibleTravelRiskEvent
type ImpossibleTravelRiskEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImpossibleTravelRiskEventRequest
func (b *ImpossibleTravelRiskEventRequestBuilder) Request() *ImpossibleTravelRiskEventRequest {
	return &ImpossibleTravelRiskEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImpossibleTravelRiskEventRequest is request for ImpossibleTravelRiskEvent
type ImpossibleTravelRiskEventRequest struct{ BaseRequest }

// Get performs GET request for ImpossibleTravelRiskEvent
func (r *ImpossibleTravelRiskEventRequest) Get(ctx context.Context) (resObj *ImpossibleTravelRiskEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ImpossibleTravelRiskEvent
func (r *ImpossibleTravelRiskEventRequest) Update(ctx context.Context, reqObj *ImpossibleTravelRiskEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ImpossibleTravelRiskEvent
func (r *ImpossibleTravelRiskEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
