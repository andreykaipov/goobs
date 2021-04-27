// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceSettingsParams represents the params body for the "GetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings.
*/
type GetSourceSettingsParams struct {
	requests.Params

	// Source name.
	SourceName string `json:"sourceName"`

	// Type of the specified source. Useful for type-checking if you expect a specific settings
	// schema.
	SourceType string `json:"sourceType"`
}

// GetRequestType returns the RequestType of GetSourceSettingsParams
func (o *GetSourceSettingsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSourceSettingsParams
func (o *GetSourceSettingsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSourceSettingsParams
func (o *GetSourceSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourceSettingsResponse represents the response body for the "GetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings.
*/
type GetSourceSettingsResponse struct {
	requests.Response

	// Source name
	SourceName string `json:"sourceName"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source
	SourceType string `json:"sourceType"`
}

// GetMessageID returns the MessageID of GetSourceSettingsResponse
func (o *GetSourceSettingsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSourceSettingsResponse
func (o *GetSourceSettingsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSourceSettingsResponse
func (o *GetSourceSettingsResponse) GetError() string {
	return o.Error
}

// GetSourceSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceSettings(
	params *GetSourceSettingsParams,
) (*GetSourceSettingsResponse, error) {
	params.RequestType = "GetSourceSettings"
	data := &GetSourceSettingsResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
