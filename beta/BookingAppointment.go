// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// BookingAppointment Represents a booked appointment of a service by a customer in a business.
type BookingAppointment struct {
	// Entity is the base model of BookingAppointment
	Entity
	// SelfServiceAppointmentID undocumented
	SelfServiceAppointmentID *string `json:"selfServiceAppointmentId,omitempty"`
	// CustomerID The id of the booking customer associated with this appointment.
	CustomerID *string `json:"customerId,omitempty"`
	// CustomerName undocumented
	CustomerName *string `json:"customerName,omitempty"`
	// CustomerEmailAddress undocumented
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	// CustomerPhone undocumented
	CustomerPhone *string `json:"customerPhone,omitempty"`
	// CustomerLocation undocumented
	CustomerLocation *Location `json:"customerLocation,omitempty"`
	// CustomerNotes Notes from the customer associated with this appointment.
	CustomerNotes *string `json:"customerNotes,omitempty"`
	// ServiceID The id of the booking service associated with this appointment.
	ServiceID *string `json:"serviceId,omitempty"`
	// ServiceName The name of the booking service associated with this appointment.
	ServiceName *string `json:"serviceName,omitempty"`
	// Start undocumented
	Start *DateTimeTimeZone `json:"start,omitempty"`
	// End undocumented
	End *DateTimeTimeZone `json:"end,omitempty"`
	// Duration undocumented
	Duration *Duration `json:"duration,omitempty"`
	// PreBuffer undocumented
	PreBuffer *Duration `json:"preBuffer,omitempty"`
	// PostBuffer undocumented
	PostBuffer *Duration `json:"postBuffer,omitempty"`
	// ServiceLocation undocumented
	ServiceLocation *Location `json:"serviceLocation,omitempty"`
	// PriceType undocumented
	PriceType *BookingPriceType `json:"priceType,omitempty"`
	// Price undocumented
	Price *float64 `json:"price,omitempty"`
	// ServiceNotes undocumented
	ServiceNotes *string `json:"serviceNotes,omitempty"`
	// Reminders undocumented
	Reminders []BookingReminder `json:"reminders,omitempty"`
	// OptOutOfCustomerEmail undocumented
	OptOutOfCustomerEmail *bool `json:"optOutOfCustomerEmail,omitempty"`
	// StaffMemberIDs undocumented
	StaffMemberIDs []string `json:"staffMemberIds,omitempty"`
	// InvoiceAmount undocumented
	InvoiceAmount *float64 `json:"invoiceAmount,omitempty"`
	// InvoiceDate undocumented
	InvoiceDate *DateTimeTimeZone `json:"invoiceDate,omitempty"`
	// InvoiceID undocumented
	InvoiceID *string `json:"invoiceId,omitempty"`
	// InvoiceStatus undocumented
	InvoiceStatus *BookingInvoiceStatus `json:"invoiceStatus,omitempty"`
	// InvoiceURL undocumented
	InvoiceURL *string `json:"invoiceUrl,omitempty"`
}

// BookingAppointmentCancelRequestParameter undocumented
type BookingAppointmentCancelRequestParameter struct {
	// CancellationMessage undocumented
	CancellationMessage *string `json:"cancellationMessage,omitempty"`
}

// BookingAppointmentRequestBuilder is request builder for BookingAppointment
type BookingAppointmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingAppointmentRequest
func (b *BookingAppointmentRequestBuilder) Request() *BookingAppointmentRequest {
	return &BookingAppointmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingAppointmentRequest is request for BookingAppointment
type BookingAppointmentRequest struct{ BaseRequest }

// Get performs GET request for BookingAppointment
func (r *BookingAppointmentRequest) Get(ctx context.Context) (resObj *BookingAppointment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingAppointment
func (r *BookingAppointmentRequest) Update(ctx context.Context, reqObj *BookingAppointment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingAppointment
func (r *BookingAppointmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type BookingAppointmentCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumented
func (b *BookingAppointmentRequestBuilder) Cancel(reqObj *BookingAppointmentCancelRequestParameter) *BookingAppointmentCancelRequestBuilder {
	bb := &BookingAppointmentCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type BookingAppointmentCancelRequest struct{ BaseRequest }

//
func (b *BookingAppointmentCancelRequestBuilder) Request() *BookingAppointmentCancelRequest {
	return &BookingAppointmentCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *BookingAppointmentCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
