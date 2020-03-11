// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookTableSort undocumented
type WorkbookTableSort struct {
	// Entity is the base model of WorkbookTableSort
	Entity
	// Fields undocumented
	Fields []WorkbookSortField `json:"fields,omitempty"`
	// MatchCase undocumented
	MatchCase *bool `json:"matchCase,omitempty"`
	// Method undocumented
	Method *string `json:"method,omitempty"`
}

// WorkbookTableSortApplyRequestParameter undocumented
type WorkbookTableSortApplyRequestParameter struct {
	// Fields undocumented
	Fields []WorkbookSortField `json:"fields,omitempty"`
	// MatchCase undocumented
	MatchCase *bool `json:"matchCase,omitempty"`
	// Method undocumented
	Method *string `json:"method,omitempty"`
}

// WorkbookTableSortClearRequestParameter undocumented
type WorkbookTableSortClearRequestParameter struct {
}

// WorkbookTableSortReapplyRequestParameter undocumented
type WorkbookTableSortReapplyRequestParameter struct {
}

// WorkbookTableSortRequestBuilder is request builder for WorkbookTableSort
type WorkbookTableSortRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookTableSortRequest
func (b *WorkbookTableSortRequestBuilder) Request() *WorkbookTableSortRequest {
	return &WorkbookTableSortRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookTableSortRequest is request for WorkbookTableSort
type WorkbookTableSortRequest struct{ BaseRequest }

// Get performs GET request for WorkbookTableSort
func (r *WorkbookTableSortRequest) Get(ctx context.Context) (resObj *WorkbookTableSort, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookTableSort
func (r *WorkbookTableSortRequest) Update(ctx context.Context, reqObj *WorkbookTableSort) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookTableSort
func (r *WorkbookTableSortRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type WorkbookTableSortApplyRequestBuilder struct{ BaseRequestBuilder }

// Apply action undocumented
func (b *WorkbookTableSortRequestBuilder) Apply(reqObj *WorkbookTableSortApplyRequestParameter) *WorkbookTableSortApplyRequestBuilder {
	bb := &WorkbookTableSortApplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/apply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableSortApplyRequest struct{ BaseRequest }

//
func (b *WorkbookTableSortApplyRequestBuilder) Request() *WorkbookTableSortApplyRequest {
	return &WorkbookTableSortApplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableSortApplyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type WorkbookTableSortClearRequestBuilder struct{ BaseRequestBuilder }

// Clear action undocumented
func (b *WorkbookTableSortRequestBuilder) Clear(reqObj *WorkbookTableSortClearRequestParameter) *WorkbookTableSortClearRequestBuilder {
	bb := &WorkbookTableSortClearRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/clear"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableSortClearRequest struct{ BaseRequest }

//
func (b *WorkbookTableSortClearRequestBuilder) Request() *WorkbookTableSortClearRequest {
	return &WorkbookTableSortClearRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableSortClearRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type WorkbookTableSortReapplyRequestBuilder struct{ BaseRequestBuilder }

// Reapply action undocumented
func (b *WorkbookTableSortRequestBuilder) Reapply(reqObj *WorkbookTableSortReapplyRequestParameter) *WorkbookTableSortReapplyRequestBuilder {
	bb := &WorkbookTableSortReapplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reapply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableSortReapplyRequest struct{ BaseRequest }

//
func (b *WorkbookTableSortReapplyRequestBuilder) Request() *WorkbookTableSortReapplyRequest {
	return &WorkbookTableSortReapplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableSortReapplyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
