// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// DirectoryAudit undocumented
type DirectoryAudit struct {
	// Entity is the base model of DirectoryAudit
	Entity
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// CorrelationID undocumented
	CorrelationID *string `json:"correlationId,omitempty"`
	// Result undocumented
	Result *OperationResult `json:"result,omitempty"`
	// ResultReason undocumented
	ResultReason *string `json:"resultReason,omitempty"`
	// ActivityDisplayName undocumented
	ActivityDisplayName *string `json:"activityDisplayName,omitempty"`
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// LoggedByService undocumented
	LoggedByService *string `json:"loggedByService,omitempty"`
	// OperationType undocumented
	OperationType *string `json:"operationType,omitempty"`
	// InitiatedBy undocumented
	InitiatedBy *AuditActivityInitiator `json:"initiatedBy,omitempty"`
	// TargetResources undocumented
	TargetResources []TargetResource `json:"targetResources,omitempty"`
	// AdditionalDetails undocumented
	AdditionalDetails []KeyValue `json:"additionalDetails,omitempty"`
}

// DirectoryAuditRequestBuilder is request builder for DirectoryAudit
type DirectoryAuditRequestBuilder struct{ BaseRequestBuilder }

// Request returns DirectoryAuditRequest
func (b *DirectoryAuditRequestBuilder) Request() *DirectoryAuditRequest {
	return &DirectoryAuditRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DirectoryAuditRequest is request for DirectoryAudit
type DirectoryAuditRequest struct{ BaseRequest }

// Get performs GET request for DirectoryAudit
func (r *DirectoryAuditRequest) Get(ctx context.Context) (resObj *DirectoryAudit, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DirectoryAudit
func (r *DirectoryAuditRequest) Update(ctx context.Context, reqObj *DirectoryAudit) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DirectoryAudit
func (r *DirectoryAuditRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
