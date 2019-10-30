// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// UserExperienceAnalyticsMetricRequestBuilder is request builder for UserExperienceAnalyticsMetric
type UserExperienceAnalyticsMetricRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsMetricRequest
func (b *UserExperienceAnalyticsMetricRequestBuilder) Request() *UserExperienceAnalyticsMetricRequest {
	return &UserExperienceAnalyticsMetricRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsMetricRequest is request for UserExperienceAnalyticsMetric
type UserExperienceAnalyticsMetricRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsMetric
func (r *UserExperienceAnalyticsMetricRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsMetric, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsMetric
func (r *UserExperienceAnalyticsMetricRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsMetric) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsMetric
func (r *UserExperienceAnalyticsMetricRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
