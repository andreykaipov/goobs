// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetBrowserSourcePropertiesParams represents the params body for the "SetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
type SetBrowserSourcePropertiesParams struct {
	requests.Params

	// CSS to inject.
	Css string `json:"css"`

	// Framerate.
	Fps int `json:"fps"`

	// Height.
	Height int `json:"height"`

	// Indicates that a local file is in use.
	IsLocalFile bool `json:"is_local_file"`

	// file path.
	LocalFile string `json:"local_file"`

	// Visibility of the scene item.
	Render bool `json:"render"`

	// Indicates whether the source should be shutdown when not visible.
	Shutdown bool `json:"shutdown"`

	// Name of the source.
	Source string `json:"source"`

	// Url.
	Url string `json:"url"`

	// Width.
	Width int `json:"width"`
}

// GetRequestType returns the RequestType of SetBrowserSourcePropertiesParams
func (o *SetBrowserSourcePropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetBrowserSourcePropertiesParams
func (o *SetBrowserSourcePropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetBrowserSourcePropertiesParams
func (o *SetBrowserSourcePropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetBrowserSourcePropertiesResponse represents the response body for the "SetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
type SetBrowserSourcePropertiesResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetBrowserSourcePropertiesResponse
func (o *SetBrowserSourcePropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetBrowserSourcePropertiesResponse
func (o *SetBrowserSourcePropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetBrowserSourcePropertiesResponse
func (o *SetBrowserSourcePropertiesResponse) GetError() string {
	return o.Error
}

// SetBrowserSourceProperties sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) SetBrowserSourceProperties(
	params *SetBrowserSourcePropertiesParams,
) (*SetBrowserSourcePropertiesResponse, error) {
	params.RequestType = "SetBrowserSourceProperties"
	data := &SetBrowserSourcePropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
