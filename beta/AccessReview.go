// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AccessReview undocumented
type AccessReview struct {
	// Entity is the base model of AccessReview
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// CreatedBy undocumented
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`
	// BusinessFlowTemplateID undocumented
	BusinessFlowTemplateID *string `json:"businessFlowTemplateId,omitempty"`
	// ReviewerType undocumented
	ReviewerType *string `json:"reviewerType,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Settings undocumented
	Settings *AccessReviewSettings `json:"settings,omitempty"`
	// ReviewedEntity undocumented
	ReviewedEntity *Identity `json:"reviewedEntity,omitempty"`
	// Reviewers undocumented
	Reviewers []AccessReviewReviewer `json:"reviewers,omitempty"`
	// Decisions undocumented
	Decisions []AccessReviewDecision `json:"decisions,omitempty"`
	// MyDecisions undocumented
	MyDecisions []AccessReviewDecision `json:"myDecisions,omitempty"`
	// Instances undocumented
	Instances []AccessReview `json:"instances,omitempty"`
}

// AccessReviewStopRequestParameter undocumented
type AccessReviewStopRequestParameter struct {
}

// AccessReviewSendReminderRequestParameter undocumented
type AccessReviewSendReminderRequestParameter struct {
}

// AccessReviewResetDecisionsRequestParameter undocumented
type AccessReviewResetDecisionsRequestParameter struct {
}

// AccessReviewApplyDecisionsRequestParameter undocumented
type AccessReviewApplyDecisionsRequestParameter struct {
}

// AccessReviewRequestBuilder is request builder for AccessReview
type AccessReviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessReviewRequest
func (b *AccessReviewRequestBuilder) Request() *AccessReviewRequest {
	return &AccessReviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessReviewRequest is request for AccessReview
type AccessReviewRequest struct{ BaseRequest }

// Get performs GET request for AccessReview
func (r *AccessReviewRequest) Get(ctx context.Context) (resObj *AccessReview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessReview
func (r *AccessReviewRequest) Update(ctx context.Context, reqObj *AccessReview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessReview
func (r *AccessReviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Decisions returns request builder for AccessReviewDecision collection
func (b *AccessReviewRequestBuilder) Decisions() *AccessReviewDecisionsCollectionRequestBuilder {
	bb := &AccessReviewDecisionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/decisions"
	return bb
}

// AccessReviewDecisionsCollectionRequestBuilder is request builder for AccessReviewDecision collection
type AccessReviewDecisionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessReviewDecision collection
func (b *AccessReviewDecisionsCollectionRequestBuilder) Request() *AccessReviewDecisionsCollectionRequest {
	return &AccessReviewDecisionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessReviewDecision item
func (b *AccessReviewDecisionsCollectionRequestBuilder) ID(id string) *AccessReviewDecisionRequestBuilder {
	bb := &AccessReviewDecisionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessReviewDecisionsCollectionRequest is request for AccessReviewDecision collection
type AccessReviewDecisionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessReviewDecision collection
func (r *AccessReviewDecisionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessReviewDecision, error) {
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
	var values []AccessReviewDecision
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
			value  []AccessReviewDecision
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

// Get performs GET request for AccessReviewDecision collection
func (r *AccessReviewDecisionsCollectionRequest) Get(ctx context.Context) ([]AccessReviewDecision, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessReviewDecision collection
func (r *AccessReviewDecisionsCollectionRequest) Add(ctx context.Context, reqObj *AccessReviewDecision) (resObj *AccessReviewDecision, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Instances returns request builder for AccessReview collection
func (b *AccessReviewRequestBuilder) Instances() *AccessReviewInstancesCollectionRequestBuilder {
	bb := &AccessReviewInstancesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/instances"
	return bb
}

// AccessReviewInstancesCollectionRequestBuilder is request builder for AccessReview collection
type AccessReviewInstancesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessReview collection
func (b *AccessReviewInstancesCollectionRequestBuilder) Request() *AccessReviewInstancesCollectionRequest {
	return &AccessReviewInstancesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessReview item
func (b *AccessReviewInstancesCollectionRequestBuilder) ID(id string) *AccessReviewRequestBuilder {
	bb := &AccessReviewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessReviewInstancesCollectionRequest is request for AccessReview collection
type AccessReviewInstancesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessReview collection
func (r *AccessReviewInstancesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessReview, error) {
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
	var values []AccessReview
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
			value  []AccessReview
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

// Get performs GET request for AccessReview collection
func (r *AccessReviewInstancesCollectionRequest) Get(ctx context.Context) ([]AccessReview, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessReview collection
func (r *AccessReviewInstancesCollectionRequest) Add(ctx context.Context, reqObj *AccessReview) (resObj *AccessReview, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MyDecisions returns request builder for AccessReviewDecision collection
func (b *AccessReviewRequestBuilder) MyDecisions() *AccessReviewMyDecisionsCollectionRequestBuilder {
	bb := &AccessReviewMyDecisionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/myDecisions"
	return bb
}

// AccessReviewMyDecisionsCollectionRequestBuilder is request builder for AccessReviewDecision collection
type AccessReviewMyDecisionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessReviewDecision collection
func (b *AccessReviewMyDecisionsCollectionRequestBuilder) Request() *AccessReviewMyDecisionsCollectionRequest {
	return &AccessReviewMyDecisionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessReviewDecision item
func (b *AccessReviewMyDecisionsCollectionRequestBuilder) ID(id string) *AccessReviewDecisionRequestBuilder {
	bb := &AccessReviewDecisionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessReviewMyDecisionsCollectionRequest is request for AccessReviewDecision collection
type AccessReviewMyDecisionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessReviewDecision collection
func (r *AccessReviewMyDecisionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessReviewDecision, error) {
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
	var values []AccessReviewDecision
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
			value  []AccessReviewDecision
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

// Get performs GET request for AccessReviewDecision collection
func (r *AccessReviewMyDecisionsCollectionRequest) Get(ctx context.Context) ([]AccessReviewDecision, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessReviewDecision collection
func (r *AccessReviewMyDecisionsCollectionRequest) Add(ctx context.Context, reqObj *AccessReviewDecision) (resObj *AccessReviewDecision, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Reviewers returns request builder for AccessReviewReviewer collection
func (b *AccessReviewRequestBuilder) Reviewers() *AccessReviewReviewersCollectionRequestBuilder {
	bb := &AccessReviewReviewersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/reviewers"
	return bb
}

// AccessReviewReviewersCollectionRequestBuilder is request builder for AccessReviewReviewer collection
type AccessReviewReviewersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessReviewReviewer collection
func (b *AccessReviewReviewersCollectionRequestBuilder) Request() *AccessReviewReviewersCollectionRequest {
	return &AccessReviewReviewersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessReviewReviewer item
func (b *AccessReviewReviewersCollectionRequestBuilder) ID(id string) *AccessReviewReviewerRequestBuilder {
	bb := &AccessReviewReviewerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessReviewReviewersCollectionRequest is request for AccessReviewReviewer collection
type AccessReviewReviewersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessReviewReviewer collection
func (r *AccessReviewReviewersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessReviewReviewer, error) {
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
	var values []AccessReviewReviewer
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
			value  []AccessReviewReviewer
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

// Get performs GET request for AccessReviewReviewer collection
func (r *AccessReviewReviewersCollectionRequest) Get(ctx context.Context) ([]AccessReviewReviewer, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessReviewReviewer collection
func (r *AccessReviewReviewersCollectionRequest) Add(ctx context.Context, reqObj *AccessReviewReviewer) (resObj *AccessReviewReviewer, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type AccessReviewStopRequestBuilder struct{ BaseRequestBuilder }

// Stop action undocumented
func (b *AccessReviewRequestBuilder) Stop(reqObj *AccessReviewStopRequestParameter) *AccessReviewStopRequestBuilder {
	bb := &AccessReviewStopRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/stop"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AccessReviewStopRequest struct{ BaseRequest }

//
func (b *AccessReviewStopRequestBuilder) Request() *AccessReviewStopRequest {
	return &AccessReviewStopRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AccessReviewStopRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type AccessReviewSendReminderRequestBuilder struct{ BaseRequestBuilder }

// SendReminder action undocumented
func (b *AccessReviewRequestBuilder) SendReminder(reqObj *AccessReviewSendReminderRequestParameter) *AccessReviewSendReminderRequestBuilder {
	bb := &AccessReviewSendReminderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendReminder"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AccessReviewSendReminderRequest struct{ BaseRequest }

//
func (b *AccessReviewSendReminderRequestBuilder) Request() *AccessReviewSendReminderRequest {
	return &AccessReviewSendReminderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AccessReviewSendReminderRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type AccessReviewResetDecisionsRequestBuilder struct{ BaseRequestBuilder }

// ResetDecisions action undocumented
func (b *AccessReviewRequestBuilder) ResetDecisions(reqObj *AccessReviewResetDecisionsRequestParameter) *AccessReviewResetDecisionsRequestBuilder {
	bb := &AccessReviewResetDecisionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/resetDecisions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AccessReviewResetDecisionsRequest struct{ BaseRequest }

//
func (b *AccessReviewResetDecisionsRequestBuilder) Request() *AccessReviewResetDecisionsRequest {
	return &AccessReviewResetDecisionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AccessReviewResetDecisionsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type AccessReviewApplyDecisionsRequestBuilder struct{ BaseRequestBuilder }

// ApplyDecisions action undocumented
func (b *AccessReviewRequestBuilder) ApplyDecisions(reqObj *AccessReviewApplyDecisionsRequestParameter) *AccessReviewApplyDecisionsRequestBuilder {
	bb := &AccessReviewApplyDecisionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyDecisions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AccessReviewApplyDecisionsRequest struct{ BaseRequest }

//
func (b *AccessReviewApplyDecisionsRequestBuilder) Request() *AccessReviewApplyDecisionsRequest {
	return &AccessReviewApplyDecisionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AccessReviewApplyDecisionsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
