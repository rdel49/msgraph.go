// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OnPremisesPublishingProfileRequestBuilder is request builder for OnPremisesPublishingProfile
type OnPremisesPublishingProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesPublishingProfileRequest
func (b *OnPremisesPublishingProfileRequestBuilder) Request() *OnPremisesPublishingProfileRequest {
	return &OnPremisesPublishingProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesPublishingProfileRequest is request for OnPremisesPublishingProfile
type OnPremisesPublishingProfileRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesPublishingProfile
func (r *OnPremisesPublishingProfileRequest) Get(ctx context.Context) (resObj *OnPremisesPublishingProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesPublishingProfile
func (r *OnPremisesPublishingProfileRequest) Update(ctx context.Context, reqObj *OnPremisesPublishingProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesPublishingProfile
func (r *OnPremisesPublishingProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AgentGroups returns request builder for OnPremisesAgentGroup collection
func (b *OnPremisesPublishingProfileRequestBuilder) AgentGroups() *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agentGroups"
	return bb
}

// OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder is request builder for OnPremisesAgentGroup collection
type OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgentGroup collection
func (b *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileAgentGroupsCollectionRequest {
	return &OnPremisesPublishingProfileAgentGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgentGroup item
func (b *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder) ID(id string) *OnPremisesAgentGroupRequestBuilder {
	bb := &OnPremisesAgentGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileAgentGroupsCollectionRequest is request for OnPremisesAgentGroup collection
type OnPremisesPublishingProfileAgentGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnPremisesAgentGroup, error) {
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
	var values []OnPremisesAgentGroup
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
			value  []OnPremisesAgentGroup
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

// Get performs GET request for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgentGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgentGroup) (resObj *OnPremisesAgentGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Agents returns request builder for OnPremisesAgent collection
func (b *OnPremisesPublishingProfileRequestBuilder) Agents() *OnPremisesPublishingProfileAgentsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileAgentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agents"
	return bb
}

// OnPremisesPublishingProfileAgentsCollectionRequestBuilder is request builder for OnPremisesAgent collection
type OnPremisesPublishingProfileAgentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgent collection
func (b *OnPremisesPublishingProfileAgentsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileAgentsCollectionRequest {
	return &OnPremisesPublishingProfileAgentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgent item
func (b *OnPremisesPublishingProfileAgentsCollectionRequestBuilder) ID(id string) *OnPremisesAgentRequestBuilder {
	bb := &OnPremisesAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileAgentsCollectionRequest is request for OnPremisesAgent collection
type OnPremisesPublishingProfileAgentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnPremisesAgent, error) {
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
	var values []OnPremisesAgent
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
			value  []OnPremisesAgent
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

// Get performs GET request for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgent) (resObj *OnPremisesAgent, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PublishedResources returns request builder for PublishedResource collection
func (b *OnPremisesPublishingProfileRequestBuilder) PublishedResources() *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/publishedResources"
	return bb
}

// OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder is request builder for PublishedResource collection
type OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PublishedResource collection
func (b *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder) Request() *OnPremisesPublishingProfilePublishedResourcesCollectionRequest {
	return &OnPremisesPublishingProfilePublishedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PublishedResource item
func (b *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder) ID(id string) *PublishedResourceRequestBuilder {
	bb := &PublishedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfilePublishedResourcesCollectionRequest is request for PublishedResource collection
type OnPremisesPublishingProfilePublishedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]PublishedResource, error) {
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
	var values []PublishedResource
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
			value  []PublishedResource
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

// Get performs GET request for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Get(ctx context.Context) ([]PublishedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Add(ctx context.Context, reqObj *PublishedResource) (resObj *PublishedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
