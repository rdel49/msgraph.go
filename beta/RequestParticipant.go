// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ParticipantRequestBuilder is request builder for Participant
type ParticipantRequestBuilder struct{ BaseRequestBuilder }

// Request returns ParticipantRequest
func (b *ParticipantRequestBuilder) Request() *ParticipantRequest {
	return &ParticipantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ParticipantRequest is request for Participant
type ParticipantRequest struct{ BaseRequest }

// Get performs GET request for Participant
func (r *ParticipantRequest) Get(ctx context.Context) (resObj *Participant, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Participant
func (r *ParticipantRequest) Update(ctx context.Context, reqObj *Participant) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Participant
func (r *ParticipantRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ParticipantCollectionInviteRequestBuilder struct{ BaseRequestBuilder }

// Invite action undocumented
func (b *CallParticipantsCollectionRequestBuilder) Invite(reqObj *ParticipantCollectionInviteRequestParameter) *ParticipantCollectionInviteRequestBuilder {
	bb := &ParticipantCollectionInviteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/invite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantCollectionInviteRequest struct{ BaseRequest }

//
func (b *ParticipantCollectionInviteRequestBuilder) Request() *ParticipantCollectionInviteRequest {
	return &ParticipantCollectionInviteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantCollectionInviteRequest) Post(ctx context.Context) (resObj *InviteParticipantsOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ParticipantCollectionMuteAllRequestBuilder struct{ BaseRequestBuilder }

// MuteAll action undocumented
func (b *CallParticipantsCollectionRequestBuilder) MuteAll(reqObj *ParticipantCollectionMuteAllRequestParameter) *ParticipantCollectionMuteAllRequestBuilder {
	bb := &ParticipantCollectionMuteAllRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/muteAll"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantCollectionMuteAllRequest struct{ BaseRequest }

//
func (b *ParticipantCollectionMuteAllRequestBuilder) Request() *ParticipantCollectionMuteAllRequest {
	return &ParticipantCollectionMuteAllRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantCollectionMuteAllRequest) Post(ctx context.Context) (resObj *MuteParticipantsOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ParticipantMuteRequestBuilder struct{ BaseRequestBuilder }

// Mute action undocumented
func (b *ParticipantRequestBuilder) Mute(reqObj *ParticipantMuteRequestParameter) *ParticipantMuteRequestBuilder {
	bb := &ParticipantMuteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/mute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantMuteRequest struct{ BaseRequest }

//
func (b *ParticipantMuteRequestBuilder) Request() *ParticipantMuteRequest {
	return &ParticipantMuteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantMuteRequest) Post(ctx context.Context) (resObj *MuteParticipantOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}