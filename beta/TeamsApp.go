// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// TeamsApp undocumented
type TeamsApp struct {
	// Entity is the base model of TeamsApp
	Entity
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DistributionMethod undocumented
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
	// AppDefinitions undocumented
	AppDefinitions []TeamsAppDefinition `json:"appDefinitions,omitempty"`
}

// TeamsAppRequestBuilder is request builder for TeamsApp
type TeamsAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAppRequest
func (b *TeamsAppRequestBuilder) Request() *TeamsAppRequest {
	return &TeamsAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAppRequest is request for TeamsApp
type TeamsAppRequest struct{ BaseRequest }

// Get performs GET request for TeamsApp
func (r *TeamsAppRequest) Get(ctx context.Context) (resObj *TeamsApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsApp
func (r *TeamsAppRequest) Update(ctx context.Context, reqObj *TeamsApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsApp
func (r *TeamsAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AppDefinitions returns request builder for TeamsAppDefinition collection
func (b *TeamsAppRequestBuilder) AppDefinitions() *TeamsAppAppDefinitionsCollectionRequestBuilder {
	bb := &TeamsAppAppDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appDefinitions"
	return bb
}

// TeamsAppAppDefinitionsCollectionRequestBuilder is request builder for TeamsAppDefinition collection
type TeamsAppAppDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppDefinition collection
func (b *TeamsAppAppDefinitionsCollectionRequestBuilder) Request() *TeamsAppAppDefinitionsCollectionRequest {
	return &TeamsAppAppDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppDefinition item
func (b *TeamsAppAppDefinitionsCollectionRequestBuilder) ID(id string) *TeamsAppDefinitionRequestBuilder {
	bb := &TeamsAppDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamsAppAppDefinitionsCollectionRequest is request for TeamsAppDefinition collection
type TeamsAppAppDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAppDefinition collection
func (r *TeamsAppAppDefinitionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TeamsAppDefinition, error) {
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
	var values []TeamsAppDefinition
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamsAppDefinition
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
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

// Get performs GET request for TeamsAppDefinition collection
func (r *TeamsAppAppDefinitionsCollectionRequest) Get(ctx context.Context) ([]TeamsAppDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TeamsAppDefinition collection
func (r *TeamsAppAppDefinitionsCollectionRequest) Add(ctx context.Context, reqObj *TeamsAppDefinition) (resObj *TeamsAppDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
