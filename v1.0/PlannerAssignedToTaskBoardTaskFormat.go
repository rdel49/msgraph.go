// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PlannerAssignedToTaskBoardTaskFormat undocumented
type PlannerAssignedToTaskBoardTaskFormat struct {
	// Entity is the base model of PlannerAssignedToTaskBoardTaskFormat
	Entity
	// UnassignedOrderHint undocumented
	UnassignedOrderHint *string `json:"unassignedOrderHint,omitempty"`
	// OrderHintsByAssignee undocumented
	OrderHintsByAssignee *PlannerOrderHintsByAssignee `json:"orderHintsByAssignee,omitempty"`
}

// PlannerAssignedToTaskBoardTaskFormatRequestBuilder is request builder for PlannerAssignedToTaskBoardTaskFormat
type PlannerAssignedToTaskBoardTaskFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerAssignedToTaskBoardTaskFormatRequest
func (b *PlannerAssignedToTaskBoardTaskFormatRequestBuilder) Request() *PlannerAssignedToTaskBoardTaskFormatRequest {
	return &PlannerAssignedToTaskBoardTaskFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerAssignedToTaskBoardTaskFormatRequest is request for PlannerAssignedToTaskBoardTaskFormat
type PlannerAssignedToTaskBoardTaskFormatRequest struct{ BaseRequest }

// Get performs GET request for PlannerAssignedToTaskBoardTaskFormat
func (r *PlannerAssignedToTaskBoardTaskFormatRequest) Get(ctx context.Context) (resObj *PlannerAssignedToTaskBoardTaskFormat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PlannerAssignedToTaskBoardTaskFormat
func (r *PlannerAssignedToTaskBoardTaskFormatRequest) Update(ctx context.Context, reqObj *PlannerAssignedToTaskBoardTaskFormat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PlannerAssignedToTaskBoardTaskFormat
func (r *PlannerAssignedToTaskBoardTaskFormatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
