// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidWorkProfileTrustedRootCertificate Android Work Profile Trusted Root Certificate configuration profile
type AndroidWorkProfileTrustedRootCertificate struct {
	// DeviceConfiguration is the base model of AndroidWorkProfileTrustedRootCertificate
	DeviceConfiguration
	// TrustedRootCertificate Trusted Root Certificate
	TrustedRootCertificate *Binary `json:"trustedRootCertificate,omitempty"`
	// CertFileName File name to display in UI.
	CertFileName *string `json:"certFileName,omitempty"`
}

// AndroidWorkProfileTrustedRootCertificateRequestBuilder is request builder for AndroidWorkProfileTrustedRootCertificate
type AndroidWorkProfileTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidWorkProfileTrustedRootCertificateRequest
func (b *AndroidWorkProfileTrustedRootCertificateRequestBuilder) Request() *AndroidWorkProfileTrustedRootCertificateRequest {
	return &AndroidWorkProfileTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidWorkProfileTrustedRootCertificateRequest is request for AndroidWorkProfileTrustedRootCertificate
type AndroidWorkProfileTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for AndroidWorkProfileTrustedRootCertificate
func (r *AndroidWorkProfileTrustedRootCertificateRequest) Get(ctx context.Context) (resObj *AndroidWorkProfileTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidWorkProfileTrustedRootCertificate
func (r *AndroidWorkProfileTrustedRootCertificateRequest) Update(ctx context.Context, reqObj *AndroidWorkProfileTrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidWorkProfileTrustedRootCertificate
func (r *AndroidWorkProfileTrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
