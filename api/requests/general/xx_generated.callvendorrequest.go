// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CallVendorRequestParams represents the params body for the "CallVendorRequest" request.
Call a request registered to a vendor.

A vendor is a unique name registered by a third-party plugin or script, which allows for custom requests and events to be added to obs-websocket.
If a plugin or script implements vendor requests or events, documentation is expected to be provided with them.
*/
type CallVendorRequestParams struct {
	requests.ParamsBasic

	// Object containing appropriate request data
	RequestData interface{} `json:"requestData,omitempty"`

	// The request type to call
	RequestType string `json:"requestType,omitempty"`

	// Name of the vendor to use
	VendorName string `json:"vendorName,omitempty"`
}

// GetSelfName just returns "CallVendorRequest".
func (o *CallVendorRequestParams) GetSelfName() string {
	return "CallVendorRequest"
}

/*
CallVendorRequestResponse represents the response body for the "CallVendorRequest" request.
Call a request registered to a vendor.

A vendor is a unique name registered by a third-party plugin or script, which allows for custom requests and events to be added to obs-websocket.
If a plugin or script implements vendor requests or events, documentation is expected to be provided with them.
*/
type CallVendorRequestResponse struct {
	requests.ResponseBasic

	// Object containing appropriate response data. {} if request does not provide any response data
	ResponseData interface{} `json:"responseData,omitempty"`
}

// CallVendorRequest sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CallVendorRequest(params *CallVendorRequestParams) (*CallVendorRequestResponse, error) {
	data := &CallVendorRequestResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
