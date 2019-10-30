// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TeamsAppInstallationRequestBuilder is request builder for TeamsAppInstallation
type TeamsAppInstallationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAppInstallationRequest
func (b *TeamsAppInstallationRequestBuilder) Request() *TeamsAppInstallationRequest {
	return &TeamsAppInstallationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAppInstallationRequest is request for TeamsAppInstallation
type TeamsAppInstallationRequest struct{ BaseRequest }

// Get performs GET request for TeamsAppInstallation
func (r *TeamsAppInstallationRequest) Get(ctx context.Context) (resObj *TeamsAppInstallation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamsAppInstallation
func (r *TeamsAppInstallationRequest) Update(ctx context.Context, reqObj *TeamsAppInstallation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamsAppInstallation
func (r *TeamsAppInstallationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsApp is navigation property
func (b *TeamsAppInstallationRequestBuilder) TeamsApp() *TeamsAppRequestBuilder {
	bb := &TeamsAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsApp"
	return bb
}

// TeamsAppDefinition is navigation property
func (b *TeamsAppInstallationRequestBuilder) TeamsAppDefinition() *TeamsAppDefinitionRequestBuilder {
	bb := &TeamsAppDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsAppDefinition"
	return bb
}
