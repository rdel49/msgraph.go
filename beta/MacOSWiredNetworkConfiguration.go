// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacOSWiredNetworkConfiguration MacOS wired network configuration profile.
type MacOSWiredNetworkConfiguration struct {
	// DeviceConfiguration is the base model of MacOSWiredNetworkConfiguration
	DeviceConfiguration
	// NetworkName Network Name
	NetworkName *string `json:"networkName,omitempty"`
	// NetworkInterface Network interface.
	NetworkInterface *WiredNetworkInterface `json:"networkInterface,omitempty"`
	// EapType Extensible Authentication Protocol (EAP). Indicates the type of EAP protocol set on the wired network.
	EapType *EapType `json:"eapType,omitempty"`
	// EapFastConfiguration EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type.
	EapFastConfiguration *EapFastConfiguration `json:"eapFastConfiguration,omitempty"`
	// TrustedServerCertificateNames Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users devices when they connect to this wired network.
	TrustedServerCertificateNames []string `json:"trustedServerCertificateNames,omitempty"`
	// AuthenticationMethod Authentication Method when EAP Type is configured to PEAP or EAP-TTLS.
	AuthenticationMethod *WiFiAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// NonEapAuthenticationMethodForEapTtls Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password.
	NonEapAuthenticationMethodForEapTtls *NonEapAuthenticationMethodForEapTtlsType `json:"nonEapAuthenticationMethodForEapTtls,omitempty"`
	// EnableOuterIdentityPrivacy Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS, EAP-FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this wired network using their real username is displayed as 'anonymous'.
	EnableOuterIdentityPrivacy *string `json:"enableOuterIdentityPrivacy,omitempty"`
	// RootCertificateForServerValidation undocumented
	RootCertificateForServerValidation *MacOSTrustedRootCertificate `json:"rootCertificateForServerValidation,omitempty"`
	// IdentityCertificateForClientAuthentication undocumented
	IdentityCertificateForClientAuthentication *MacOSCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`
}

// MacOSWiredNetworkConfigurationRequestBuilder is request builder for MacOSWiredNetworkConfiguration
type MacOSWiredNetworkConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSWiredNetworkConfigurationRequest
func (b *MacOSWiredNetworkConfigurationRequestBuilder) Request() *MacOSWiredNetworkConfigurationRequest {
	return &MacOSWiredNetworkConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSWiredNetworkConfigurationRequest is request for MacOSWiredNetworkConfiguration
type MacOSWiredNetworkConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Get(ctx context.Context) (resObj *MacOSWiredNetworkConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Update(ctx context.Context, reqObj *MacOSWiredNetworkConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *MacOSWiredNetworkConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *MacOSWiredNetworkConfigurationRequestBuilder) RootCertificateForServerValidation() *MacOSTrustedRootCertificateRequestBuilder {
	bb := &MacOSTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}
