// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConversationThreadReplyRequestParameter undocumented
type ConversationThreadReplyRequestParameter struct {
	// Post undocumented
	Post *Post `json:"Post,omitempty"`
}

//
type ConversationThreadReplyRequestBuilder struct{ BaseRequestBuilder }

// Reply action undocumented
func (b *ConversationThreadRequestBuilder) Reply(reqObj *ConversationThreadReplyRequestParameter) *ConversationThreadReplyRequestBuilder {
	bb := &ConversationThreadReplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ConversationThreadReplyRequest struct{ BaseRequest }

//
func (b *ConversationThreadReplyRequestBuilder) Request() *ConversationThreadReplyRequest {
	return &ConversationThreadReplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ConversationThreadReplyRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *ConversationThreadReplyRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}