// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidWorkProfileEnterpriseWiFiConfiguration By providing the configurations in this profile you can instruct the Android Work Profile device to connect to desired Wi-Fi endpoint. By specifying the authentication method and security types expected by Wi-Fi endpoint you can make the Wi-Fi connection seamless for end user.
type AndroidWorkProfileEnterpriseWiFiConfiguration struct {
	// AndroidWorkProfileWiFiConfiguration is the base model of AndroidWorkProfileEnterpriseWiFiConfiguration
	AndroidWorkProfileWiFiConfiguration
	// EapType Indicates the type of EAP protocol set on the Wi-Fi endpoint (router).
	EapType *AndroidEapType `json:"eapType,omitempty"`
	// AuthenticationMethod Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS.
	AuthenticationMethod *WiFiAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// InnerAuthenticationProtocolForEapTtls Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password.
	InnerAuthenticationProtocolForEapTtls *NonEapAuthenticationMethodForEapTtlsType `json:"innerAuthenticationProtocolForEapTtls,omitempty"`
	// InnerAuthenticationProtocolForPeap Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password.
	InnerAuthenticationProtocolForPeap *NonEapAuthenticationMethodForPeap `json:"innerAuthenticationProtocolForPeap,omitempty"`
	// OuterIdentityPrivacyTemporaryValue Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
	OuterIdentityPrivacyTemporaryValue *string `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
	// RootCertificateForServerValidation undocumented
	RootCertificateForServerValidation *AndroidWorkProfileTrustedRootCertificate `json:"rootCertificateForServerValidation,omitempty"`
	// IdentityCertificateForClientAuthentication undocumented
	IdentityCertificateForClientAuthentication *AndroidWorkProfileCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`
}

// AndroidWorkProfileEnterpriseWiFiConfigurationRequestBuilder is request builder for AndroidWorkProfileEnterpriseWiFiConfiguration
type AndroidWorkProfileEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidWorkProfileEnterpriseWiFiConfigurationRequest
func (b *AndroidWorkProfileEnterpriseWiFiConfigurationRequestBuilder) Request() *AndroidWorkProfileEnterpriseWiFiConfigurationRequest {
	return &AndroidWorkProfileEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidWorkProfileEnterpriseWiFiConfigurationRequest is request for AndroidWorkProfileEnterpriseWiFiConfiguration
type AndroidWorkProfileEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AndroidWorkProfileEnterpriseWiFiConfiguration
func (r *AndroidWorkProfileEnterpriseWiFiConfigurationRequest) Get(ctx context.Context) (resObj *AndroidWorkProfileEnterpriseWiFiConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidWorkProfileEnterpriseWiFiConfiguration
func (r *AndroidWorkProfileEnterpriseWiFiConfigurationRequest) Update(ctx context.Context, reqObj *AndroidWorkProfileEnterpriseWiFiConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidWorkProfileEnterpriseWiFiConfiguration
func (r *AndroidWorkProfileEnterpriseWiFiConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *AndroidWorkProfileEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *AndroidWorkProfileCertificateProfileBaseRequestBuilder {
	bb := &AndroidWorkProfileCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *AndroidWorkProfileEnterpriseWiFiConfigurationRequestBuilder) RootCertificateForServerValidation() *AndroidWorkProfileTrustedRootCertificateRequestBuilder {
	bb := &AndroidWorkProfileTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}
