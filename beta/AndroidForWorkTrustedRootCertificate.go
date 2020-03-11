// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidForWorkTrustedRootCertificate Android For Work Trusted Root Certificate configuration profile
type AndroidForWorkTrustedRootCertificate struct {
	// DeviceConfiguration is the base model of AndroidForWorkTrustedRootCertificate
	DeviceConfiguration
	// TrustedRootCertificate Trusted Root Certificate
	TrustedRootCertificate *Binary `json:"trustedRootCertificate,omitempty"`
	// CertFileName File name to display in UI.
	CertFileName *string `json:"certFileName,omitempty"`
}

// AndroidForWorkTrustedRootCertificateRequestBuilder is request builder for AndroidForWorkTrustedRootCertificate
type AndroidForWorkTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkTrustedRootCertificateRequest
func (b *AndroidForWorkTrustedRootCertificateRequestBuilder) Request() *AndroidForWorkTrustedRootCertificateRequest {
	return &AndroidForWorkTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkTrustedRootCertificateRequest is request for AndroidForWorkTrustedRootCertificate
type AndroidForWorkTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for AndroidForWorkTrustedRootCertificate
func (r *AndroidForWorkTrustedRootCertificateRequest) Get(ctx context.Context) (resObj *AndroidForWorkTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidForWorkTrustedRootCertificate
func (r *AndroidForWorkTrustedRootCertificateRequest) Update(ctx context.Context, reqObj *AndroidForWorkTrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidForWorkTrustedRootCertificate
func (r *AndroidForWorkTrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
