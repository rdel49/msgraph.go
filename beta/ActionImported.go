// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rdel49/msgraph.go/jsonx"
)

// ImportedAppleDeviceIdentityCollectionImportAppleDeviceIdentityListRequestParameter undocumented
type ImportedAppleDeviceIdentityCollectionImportAppleDeviceIdentityListRequestParameter struct {
	// ImportedAppleDeviceIdentities undocumented
	ImportedAppleDeviceIdentities []ImportedAppleDeviceIdentity `json:"importedAppleDeviceIdentities,omitempty"`
	// OverwriteImportedDeviceIdentities undocumented
	OverwriteImportedDeviceIdentities *bool `json:"overwriteImportedDeviceIdentities,omitempty"`
}

// ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestParameter undocumented
type ImportedDeviceIdentityCollectionImportDeviceIdentityListRequestParameter struct {
	// ImportedDeviceIdentities undocumented
	ImportedDeviceIdentities []ImportedDeviceIdentity `json:"importedDeviceIdentities,omitempty"`
	// OverwriteImportedDeviceIdentities undocumented
	OverwriteImportedDeviceIdentities *bool `json:"overwriteImportedDeviceIdentities,omitempty"`
}

// ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestParameter undocumented
type ImportedDeviceIdentityCollectionSearchExistingIdentitiesRequestParameter struct {
	// ImportedDeviceIdentities undocumented
	ImportedDeviceIdentities []ImportedDeviceIdentity `json:"importedDeviceIdentities,omitempty"`
}

// ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestParameter undocumented
type ImportedWindowsAutopilotDeviceIdentityCollectionImportRequestParameter struct {
	// ImportedWindowsAutopilotDeviceIdentities undocumented
	ImportedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentity `json:"importedWindowsAutopilotDeviceIdentities,omitempty"`
}

// DeviceIdentities returns request builder for ImportedWindowsAutopilotDeviceIdentity collection
func (b *ImportedWindowsAutopilotDeviceIdentityUploadRequestBuilder) DeviceIdentities() *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequestBuilder {
	bb := &ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceIdentities"
	return bb
}

// ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequestBuilder is request builder for ImportedWindowsAutopilotDeviceIdentity collection
type ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ImportedWindowsAutopilotDeviceIdentity collection
func (b *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequestBuilder) Request() *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest {
	return &ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ImportedWindowsAutopilotDeviceIdentity item
func (b *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequestBuilder) ID(id string) *ImportedWindowsAutopilotDeviceIdentityRequestBuilder {
	bb := &ImportedWindowsAutopilotDeviceIdentityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest is request for ImportedWindowsAutopilotDeviceIdentity collection
type ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ImportedWindowsAutopilotDeviceIdentity collection
func (r *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ImportedWindowsAutopilotDeviceIdentity, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ImportedWindowsAutopilotDeviceIdentity
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ImportedWindowsAutopilotDeviceIdentity
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ImportedWindowsAutopilotDeviceIdentity collection, max N pages
func (r *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest) GetN(ctx context.Context, n int) ([]ImportedWindowsAutopilotDeviceIdentity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ImportedWindowsAutopilotDeviceIdentity collection
func (r *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest) Get(ctx context.Context) ([]ImportedWindowsAutopilotDeviceIdentity, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ImportedWindowsAutopilotDeviceIdentity collection
func (r *ImportedWindowsAutopilotDeviceIdentityUploadDeviceIdentitiesCollectionRequest) Add(ctx context.Context, reqObj *ImportedWindowsAutopilotDeviceIdentity) (resObj *ImportedWindowsAutopilotDeviceIdentity, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
