// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the CallVendorRequest request.
type CallVendorRequestParams struct {
	// Object containing appropriate request data
	RequestData map[string]any `json:"requestData,omitempty"`

	// The request type to call
	RequestType *string `json:"requestType,omitempty"`

	// Name of the vendor to use
	VendorName *string `json:"vendorName,omitempty"`
}

func NewCallVendorRequestParams() *CallVendorRequestParams {
	return &CallVendorRequestParams{}
}
func (o *CallVendorRequestParams) WithRequestData(x map[string]any) *CallVendorRequestParams {
	o.RequestData = x
	return o
}
func (o *CallVendorRequestParams) WithRequestType(x string) *CallVendorRequestParams {
	o.RequestType = &x
	return o
}
func (o *CallVendorRequestParams) WithVendorName(x string) *CallVendorRequestParams {
	o.VendorName = &x
	return o
}

// Returns the associated request.
func (o *CallVendorRequestParams) GetRequestName() string {
	return "CallVendorRequest"
}

// Represents the response body for the CallVendorRequest request.
type CallVendorRequestResponse struct {
	_response

	// Echoed of `requestType`
	RequestType string `json:"requestType,omitempty"`

	// Object containing appropriate response data. {} if request does not provide any response data
	ResponseData map[string]any `json:"responseData,omitempty"`

	// Echoed of `vendorName`
	VendorName string `json:"vendorName,omitempty"`
}

/*
Call a request registered to a vendor.

A vendor is a unique name registered by a third-party plugin or script, which allows for custom requests and events to be added to obs-websocket.
If a plugin or script implements vendor requests or events, documentation is expected to be provided with them.
*/
func (c *Client) CallVendorRequest(params *CallVendorRequestParams) (*CallVendorRequestResponse, error) {
	data := &CallVendorRequestResponse{}
	return data, c.client.SendRequest(params, data)
}
