// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// UserExperienceAnalyticsBaseline The user experience analytics baseline entity contains baseline values against which to compare the user experience analytics scores.
type UserExperienceAnalyticsBaseline struct {
	// Entity is the base model of UserExperienceAnalyticsBaseline
	Entity
	// DisplayName The name of the user experience analytics baseline.
	DisplayName *string `json:"displayName,omitempty"`
	// IsBuiltIn Signifies if the current baseline is the commercial median baseline or a custom baseline.
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
	// CreatedDateTime The date the custom baseline was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DeviceBootPerformanceMetrics undocumented
	DeviceBootPerformanceMetrics *UserExperienceAnalyticsCategory `json:"deviceBootPerformanceMetrics,omitempty"`
	// BestPracticesMetrics undocumented
	BestPracticesMetrics *UserExperienceAnalyticsCategory `json:"bestPracticesMetrics,omitempty"`
}

// UserExperienceAnalyticsBaselineRequestBuilder is request builder for UserExperienceAnalyticsBaseline
type UserExperienceAnalyticsBaselineRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsBaselineRequest
func (b *UserExperienceAnalyticsBaselineRequestBuilder) Request() *UserExperienceAnalyticsBaselineRequest {
	return &UserExperienceAnalyticsBaselineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsBaselineRequest is request for UserExperienceAnalyticsBaseline
type UserExperienceAnalyticsBaselineRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsBaseline, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsBaseline) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BestPracticesMetrics is navigation property
func (b *UserExperienceAnalyticsBaselineRequestBuilder) BestPracticesMetrics() *UserExperienceAnalyticsCategoryRequestBuilder {
	bb := &UserExperienceAnalyticsCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/bestPracticesMetrics"
	return bb
}

// DeviceBootPerformanceMetrics is navigation property
func (b *UserExperienceAnalyticsBaselineRequestBuilder) DeviceBootPerformanceMetrics() *UserExperienceAnalyticsCategoryRequestBuilder {
	bb := &UserExperienceAnalyticsCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceBootPerformanceMetrics"
	return bb
}
