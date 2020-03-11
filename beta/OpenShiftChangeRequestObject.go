// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// OpenShiftChangeRequestObject undocumented
type OpenShiftChangeRequestObject struct {
	// ScheduleChangeRequestObject is the base model of OpenShiftChangeRequestObject
	ScheduleChangeRequestObject
	// OpenShiftID undocumented
	OpenShiftID *string `json:"openShiftId,omitempty"`
}

// OpenShiftChangeRequestObjectRequestBuilder is request builder for OpenShiftChangeRequestObject
type OpenShiftChangeRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns OpenShiftChangeRequestObjectRequest
func (b *OpenShiftChangeRequestObjectRequestBuilder) Request() *OpenShiftChangeRequestObjectRequest {
	return &OpenShiftChangeRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OpenShiftChangeRequestObjectRequest is request for OpenShiftChangeRequestObject
type OpenShiftChangeRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for OpenShiftChangeRequestObject
func (r *OpenShiftChangeRequestObjectRequest) Get(ctx context.Context) (resObj *OpenShiftChangeRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OpenShiftChangeRequestObject
func (r *OpenShiftChangeRequestObjectRequest) Update(ctx context.Context, reqObj *OpenShiftChangeRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OpenShiftChangeRequestObject
func (r *OpenShiftChangeRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
