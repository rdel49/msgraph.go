// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// CustomerPayment undocumented
type CustomerPayment struct {
	// Entity is the base model of CustomerPayment
	Entity
	// JournalDisplayName undocumented
	JournalDisplayName *string `json:"journalDisplayName,omitempty"`
	// LineNumber undocumented
	LineNumber *int `json:"lineNumber,omitempty"`
	// CustomerID undocumented
	CustomerID *UUID `json:"customerId,omitempty"`
	// CustomerNumber undocumented
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// ContactID undocumented
	ContactID *string `json:"contactId,omitempty"`
	// PostingDate undocumented
	PostingDate *Date `json:"postingDate,omitempty"`
	// DocumentNumber undocumented
	DocumentNumber *string `json:"documentNumber,omitempty"`
	// ExternalDocumentNumber undocumented
	ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
	// Amount undocumented
	Amount *int `json:"amount,omitempty"`
	// AppliesToInvoiceID undocumented
	AppliesToInvoiceID *UUID `json:"appliesToInvoiceId,omitempty"`
	// AppliesToInvoiceNumber undocumented
	AppliesToInvoiceNumber *string `json:"appliesToInvoiceNumber,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Customer undocumented
	Customer *Customer `json:"customer,omitempty"`
}

// CustomerPaymentRequestBuilder is request builder for CustomerPayment
type CustomerPaymentRequestBuilder struct{ BaseRequestBuilder }

// Request returns CustomerPaymentRequest
func (b *CustomerPaymentRequestBuilder) Request() *CustomerPaymentRequest {
	return &CustomerPaymentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CustomerPaymentRequest is request for CustomerPayment
type CustomerPaymentRequest struct{ BaseRequest }

// Get performs GET request for CustomerPayment
func (r *CustomerPaymentRequest) Get(ctx context.Context) (resObj *CustomerPayment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CustomerPayment
func (r *CustomerPaymentRequest) Update(ctx context.Context, reqObj *CustomerPayment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CustomerPayment
func (r *CustomerPaymentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Customer is navigation property
func (b *CustomerPaymentRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}
