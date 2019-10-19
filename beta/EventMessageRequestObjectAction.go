// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EventMessageRequestObjectAcceptRequestParameter undocumented
type EventMessageRequestObjectAcceptRequestParameter struct {
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventMessageRequestObjectDeclineRequestParameter undocumented
type EventMessageRequestObjectDeclineRequestParameter struct {
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

// EventMessageRequestObjectTentativelyAcceptRequestParameter undocumented
type EventMessageRequestObjectTentativelyAcceptRequestParameter struct {
	// SendResponse undocumented
	SendResponse *bool `json:"SendResponse,omitempty"`
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
}

//
type EventMessageRequestObjectAcceptRequestBuilder struct{ BaseRequestBuilder }

// Accept action undocumented
func (b *EventMessageRequestObjectRequestBuilder) Accept(reqObj *EventMessageRequestObjectAcceptRequestParameter) *EventMessageRequestObjectAcceptRequestBuilder {
	bb := &EventMessageRequestObjectAcceptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/accept"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EventMessageRequestObjectAcceptRequest struct{ BaseRequest }

//
func (b *EventMessageRequestObjectAcceptRequestBuilder) Request() *EventMessageRequestObjectAcceptRequest {
	return &EventMessageRequestObjectAcceptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EventMessageRequestObjectAcceptRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *EventMessageRequestObjectAcceptRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type EventMessageRequestObjectDeclineRequestBuilder struct{ BaseRequestBuilder }

// Decline action undocumented
func (b *EventMessageRequestObjectRequestBuilder) Decline(reqObj *EventMessageRequestObjectDeclineRequestParameter) *EventMessageRequestObjectDeclineRequestBuilder {
	bb := &EventMessageRequestObjectDeclineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/decline"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EventMessageRequestObjectDeclineRequest struct{ BaseRequest }

//
func (b *EventMessageRequestObjectDeclineRequestBuilder) Request() *EventMessageRequestObjectDeclineRequest {
	return &EventMessageRequestObjectDeclineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EventMessageRequestObjectDeclineRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *EventMessageRequestObjectDeclineRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type EventMessageRequestObjectTentativelyAcceptRequestBuilder struct{ BaseRequestBuilder }

// TentativelyAccept action undocumented
func (b *EventMessageRequestObjectRequestBuilder) TentativelyAccept(reqObj *EventMessageRequestObjectTentativelyAcceptRequestParameter) *EventMessageRequestObjectTentativelyAcceptRequestBuilder {
	bb := &EventMessageRequestObjectTentativelyAcceptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/tentativelyAccept"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EventMessageRequestObjectTentativelyAcceptRequest struct{ BaseRequest }

//
func (b *EventMessageRequestObjectTentativelyAcceptRequestBuilder) Request() *EventMessageRequestObjectTentativelyAcceptRequest {
	return &EventMessageRequestObjectTentativelyAcceptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EventMessageRequestObjectTentativelyAcceptRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *EventMessageRequestObjectTentativelyAcceptRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}