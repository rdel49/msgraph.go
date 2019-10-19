// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ApplicationTemplateInstantiateRequestParameter undocumented
type ApplicationTemplateInstantiateRequestParameter struct {
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

//
type ApplicationTemplateInstantiateRequestBuilder struct{ BaseRequestBuilder }

// Instantiate action undocumented
func (b *ApplicationTemplateRequestBuilder) Instantiate(reqObj *ApplicationTemplateInstantiateRequestParameter) *ApplicationTemplateInstantiateRequestBuilder {
	bb := &ApplicationTemplateInstantiateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/instantiate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ApplicationTemplateInstantiateRequest struct{ BaseRequest }

//
func (b *ApplicationTemplateInstantiateRequestBuilder) Request() *ApplicationTemplateInstantiateRequest {
	return &ApplicationTemplateInstantiateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ApplicationTemplateInstantiateRequest) Do(method, path string, reqObj interface{}) (resObj *ApplicationServicePrincipal, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *ApplicationTemplateInstantiateRequest) Post() (*ApplicationServicePrincipal, error) {
	return r.Do("POST", "", r.requestObject)
}