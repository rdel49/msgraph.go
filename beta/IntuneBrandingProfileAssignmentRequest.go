// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IntuneBrandingProfileAssignmentRequestBuilder is request builder for IntuneBrandingProfileAssignment
type IntuneBrandingProfileAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns IntuneBrandingProfileAssignmentRequest
func (b *IntuneBrandingProfileAssignmentRequestBuilder) Request() *IntuneBrandingProfileAssignmentRequest {
	return &IntuneBrandingProfileAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IntuneBrandingProfileAssignmentRequest is request for IntuneBrandingProfileAssignment
type IntuneBrandingProfileAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for IntuneBrandingProfileAssignment
func (r *IntuneBrandingProfileAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *IntuneBrandingProfileAssignment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for IntuneBrandingProfileAssignment
func (r *IntuneBrandingProfileAssignmentRequest) Get() (*IntuneBrandingProfileAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for IntuneBrandingProfileAssignment
func (r *IntuneBrandingProfileAssignmentRequest) Update(reqObj *IntuneBrandingProfileAssignment) (*IntuneBrandingProfileAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for IntuneBrandingProfileAssignment
func (r *IntuneBrandingProfileAssignmentRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}