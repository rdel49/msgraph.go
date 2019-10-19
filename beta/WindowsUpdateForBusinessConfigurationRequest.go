// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WindowsUpdateForBusinessConfigurationRequestBuilder is request builder for WindowsUpdateForBusinessConfiguration
type WindowsUpdateForBusinessConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsUpdateForBusinessConfigurationRequest
func (b *WindowsUpdateForBusinessConfigurationRequestBuilder) Request() *WindowsUpdateForBusinessConfigurationRequest {
	return &WindowsUpdateForBusinessConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsUpdateForBusinessConfigurationRequest is request for WindowsUpdateForBusinessConfiguration
type WindowsUpdateForBusinessConfigurationRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsUpdateForBusinessConfiguration
func (r *WindowsUpdateForBusinessConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsUpdateForBusinessConfiguration, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WindowsUpdateForBusinessConfiguration
func (r *WindowsUpdateForBusinessConfigurationRequest) Get() (*WindowsUpdateForBusinessConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WindowsUpdateForBusinessConfiguration
func (r *WindowsUpdateForBusinessConfigurationRequest) Update(reqObj *WindowsUpdateForBusinessConfiguration) (*WindowsUpdateForBusinessConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WindowsUpdateForBusinessConfiguration
func (r *WindowsUpdateForBusinessConfigurationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// DeviceUpdateStates returns request builder for WindowsUpdateState collection
func (b *WindowsUpdateForBusinessConfigurationRequestBuilder) DeviceUpdateStates() *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequestBuilder {
	bb := &WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceUpdateStates"
	return bb
}

// WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequestBuilder is request builder for WindowsUpdateState collection
type WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsUpdateState collection
func (b *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequestBuilder) Request() *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest {
	return &WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsUpdateState item
func (b *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequestBuilder) ID(id string) *WindowsUpdateStateRequestBuilder {
	bb := &WindowsUpdateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest is request for WindowsUpdateState collection
type WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsUpdateState collection
func (r *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsUpdateState, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for WindowsUpdateState collection
func (r *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]WindowsUpdateState, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WindowsUpdateState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WindowsUpdateState
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for WindowsUpdateState collection
func (r *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest) Get() ([]WindowsUpdateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WindowsUpdateState collection
func (r *WindowsUpdateForBusinessConfigurationDeviceUpdateStatesCollectionRequest) Add(reqObj *WindowsUpdateState) (*WindowsUpdateState, error) {
	return r.Do("POST", "", reqObj)
}