// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// Person undocumented
type Person struct {
	// Entity is the base model of Person
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// Birthday undocumented
	Birthday *string `json:"birthday,omitempty"`
	// PersonNotes undocumented
	PersonNotes *string `json:"personNotes,omitempty"`
	// IsFavorite undocumented
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// EmailAddresses undocumented
	EmailAddresses []RankedEmailAddress `json:"emailAddresses,omitempty"`
	// Phones undocumented
	Phones []Phone `json:"phones,omitempty"`
	// PostalAddresses undocumented
	PostalAddresses []Location `json:"postalAddresses,omitempty"`
	// Websites undocumented
	Websites []Website `json:"websites,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// YomiCompany undocumented
	YomiCompany *string `json:"yomiCompany,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// Profession undocumented
	Profession *string `json:"profession,omitempty"`
	// Sources undocumented
	Sources []PersonDataSource `json:"sources,omitempty"`
	// MailboxType undocumented
	MailboxType *string `json:"mailboxType,omitempty"`
	// PersonType undocumented
	PersonType *string `json:"personType,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// PersonRequestBuilder is request builder for Person
type PersonRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonRequest
func (b *PersonRequestBuilder) Request() *PersonRequest {
	return &PersonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonRequest is request for Person
type PersonRequest struct{ BaseRequest }

// Get performs GET request for Person
func (r *PersonRequest) Get(ctx context.Context) (resObj *Person, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Person
func (r *PersonRequest) Update(ctx context.Context, reqObj *Person) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Person
func (r *PersonRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
