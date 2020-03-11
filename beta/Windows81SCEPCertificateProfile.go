// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// Windows81SCEPCertificateProfile Windows 8.1+ SCEP certificate profile
type Windows81SCEPCertificateProfile struct {
	// Windows81CertificateProfileBase is the base model of Windows81SCEPCertificateProfile
	Windows81CertificateProfileBase
	// ScepServerUrls SCEP Server Url(s).
	ScepServerUrls []string `json:"scepServerUrls,omitempty"`
	// SubjectNameFormatString Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
	SubjectNameFormatString *string `json:"subjectNameFormatString,omitempty"`
	// KeyUsage SCEP Key Usage.
	KeyUsage *KeyUsages `json:"keyUsage,omitempty"`
	// KeySize SCEP Key Size.
	KeySize *KeySize `json:"keySize,omitempty"`
	// HashAlgorithm SCEP Hash Algorithm.
	HashAlgorithm *HashAlgorithms `json:"hashAlgorithm,omitempty"`
	// SubjectAlternativeNameFormatString Custom String that defines the AAD Attribute.
	SubjectAlternativeNameFormatString *string `json:"subjectAlternativeNameFormatString,omitempty"`
	// CertificateStore Target store certificate
	CertificateStore *CertificateStore `json:"certificateStore,omitempty"`
	// RootCertificate undocumented
	RootCertificate *Windows81TrustedRootCertificate `json:"rootCertificate,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}

// Windows81SCEPCertificateProfileRequestBuilder is request builder for Windows81SCEPCertificateProfile
type Windows81SCEPCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows81SCEPCertificateProfileRequest
func (b *Windows81SCEPCertificateProfileRequestBuilder) Request() *Windows81SCEPCertificateProfileRequest {
	return &Windows81SCEPCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows81SCEPCertificateProfileRequest is request for Windows81SCEPCertificateProfile
type Windows81SCEPCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for Windows81SCEPCertificateProfile
func (r *Windows81SCEPCertificateProfileRequest) Get(ctx context.Context) (resObj *Windows81SCEPCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Windows81SCEPCertificateProfile
func (r *Windows81SCEPCertificateProfileRequest) Update(ctx context.Context, reqObj *Windows81SCEPCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Windows81SCEPCertificateProfile
func (r *Windows81SCEPCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *Windows81SCEPCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
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
	var values []ManagedDeviceCertificateState
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
			value  []ManagedDeviceCertificateState
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

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *Windows81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RootCertificate is navigation property
func (b *Windows81SCEPCertificateProfileRequestBuilder) RootCertificate() *Windows81TrustedRootCertificateRequestBuilder {
	bb := &Windows81TrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}
