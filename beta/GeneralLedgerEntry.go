// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// GeneralLedgerEntry undocumented
type GeneralLedgerEntry struct {
	// Entity is the base model of GeneralLedgerEntry
	Entity
	// PostingDate undocumented
	PostingDate *Date `json:"postingDate,omitempty"`
	// DocumentNumber undocumented
	DocumentNumber *string `json:"documentNumber,omitempty"`
	// DocumentType undocumented
	DocumentType *string `json:"documentType,omitempty"`
	// AccountID undocumented
	AccountID *UUID `json:"accountId,omitempty"`
	// AccountNumber undocumented
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DebitAmount undocumented
	DebitAmount *int `json:"debitAmount,omitempty"`
	// CreditAmount undocumented
	CreditAmount *int `json:"creditAmount,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Account undocumented
	Account *Account `json:"account,omitempty"`
}

// GeneralLedgerEntryRequestBuilder is request builder for GeneralLedgerEntry
type GeneralLedgerEntryRequestBuilder struct{ BaseRequestBuilder }

// Request returns GeneralLedgerEntryRequest
func (b *GeneralLedgerEntryRequestBuilder) Request() *GeneralLedgerEntryRequest {
	return &GeneralLedgerEntryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GeneralLedgerEntryRequest is request for GeneralLedgerEntry
type GeneralLedgerEntryRequest struct{ BaseRequest }

// Get performs GET request for GeneralLedgerEntry
func (r *GeneralLedgerEntryRequest) Get(ctx context.Context) (resObj *GeneralLedgerEntry, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GeneralLedgerEntry
func (r *GeneralLedgerEntryRequest) Update(ctx context.Context, reqObj *GeneralLedgerEntry) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GeneralLedgerEntry
func (r *GeneralLedgerEntryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Account is navigation property
func (b *GeneralLedgerEntryRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}
