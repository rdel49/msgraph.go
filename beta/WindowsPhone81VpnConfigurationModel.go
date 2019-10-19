// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhone81VpnConfiguration By providing the configurations in this profile you can instruct the Windows Phone 8.1 to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type WindowsPhone81VpnConfiguration struct {
	Windows81VpnConfiguration
	// BypassVpnOnCompanyWifi Bypass VPN on company Wi-Fi.
	BypassVpnOnCompanyWifi *bool `json:"bypassVpnOnCompanyWifi,omitempty"`
	// BypassVpnOnHomeWifi Bypass VPN on home Wi-Fi.
	BypassVpnOnHomeWifi *bool `json:"bypassVpnOnHomeWifi,omitempty"`
	// AuthenticationMethod Authentication method.
	AuthenticationMethod *VpnAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// RememberUserCredentials Remember user credentials.
	RememberUserCredentials *bool `json:"rememberUserCredentials,omitempty"`
	// DNSSuffixSearchList DNS suffix search list.
	DNSSuffixSearchList []string `json:"dnsSuffixSearchList,omitempty"`
	// IdentityCertificate undocumented
	IdentityCertificate *WindowsPhone81CertificateProfileBase `json:"identityCertificate,omitempty"`
}