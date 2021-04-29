// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceSettingsParams represents the params body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsParams struct {
	requests.Params

	// Source name.
	SourceName string `json:"sourceName"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source. Useful for type-checking to avoid settings a set of settings
	// incompatible with the actual source's type.
	SourceType string `json:"sourceType"`
}

// GetRequestType returns the RequestType of SetSourceSettingsParams
func (o *SetSourceSettingsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetSourceSettingsParams
func (o *SetSourceSettingsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetSourceSettingsParams
func (o *SetSourceSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSourceSettingsResponse represents the response body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsResponse struct {
	requests.Response

	// Source name
	SourceName string `json:"sourceName"`

	// Updated source settings
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source
	SourceType string `json:"sourceType"`
}

// GetMessageID returns the MessageID of SetSourceSettingsResponse
func (o *SetSourceSettingsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetSourceSettingsResponse
func (o *SetSourceSettingsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetSourceSettingsResponse
func (o *SetSourceSettingsResponse) GetError() string {
	return o.Error
}

// SetSourceSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceSettings(
	params *SetSourceSettingsParams,
) (*SetSourceSettingsResponse, error) {
	params.RequestType = "SetSourceSettings"
	data := &SetSourceSettingsResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
