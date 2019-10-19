// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SingleValueLegacyExtendedPropertyRequestBuilder is request builder for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns SingleValueLegacyExtendedPropertyRequest
func (b *SingleValueLegacyExtendedPropertyRequestBuilder) Request() *SingleValueLegacyExtendedPropertyRequest {
	return &SingleValueLegacyExtendedPropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SingleValueLegacyExtendedPropertyRequest is request for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequest struct{ BaseRequest }

// Do performs HTTP request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Do(method, path string, reqObj interface{}) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Get() (*SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Update(reqObj *SingleValueLegacyExtendedProperty) (*SingleValueLegacyExtendedProperty, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}