// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceCompliancePolicyStateRequestBuilder is request builder for DeviceCompliancePolicyState
type DeviceCompliancePolicyStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceCompliancePolicyStateRequest
func (b *DeviceCompliancePolicyStateRequestBuilder) Request() *DeviceCompliancePolicyStateRequest {
	return &DeviceCompliancePolicyStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceCompliancePolicyStateRequest is request for DeviceCompliancePolicyState
type DeviceCompliancePolicyStateRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceCompliancePolicyState
func (r *DeviceCompliancePolicyStateRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceCompliancePolicyState, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceCompliancePolicyState
func (r *DeviceCompliancePolicyStateRequest) Get() (*DeviceCompliancePolicyState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceCompliancePolicyState
func (r *DeviceCompliancePolicyStateRequest) Update(reqObj *DeviceCompliancePolicyState) (*DeviceCompliancePolicyState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceCompliancePolicyState
func (r *DeviceCompliancePolicyStateRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}