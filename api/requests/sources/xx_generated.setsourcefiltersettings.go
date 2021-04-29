// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterSettingsParams represents the params body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsParams struct {
	requests.Params

	// Name of the filter to reconfigure
	FilterName string `json:"filterName"`

	// New settings. These will be merged to the current filter settings.
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// GetRequestType returns the RequestType of SetSourceFilterSettingsParams
func (o *SetSourceFilterSettingsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetSourceFilterSettingsParams
func (o *SetSourceFilterSettingsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetSourceFilterSettingsParams
func (o *SetSourceFilterSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSourceFilterSettingsResponse represents the response body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetSourceFilterSettingsResponse
func (o *SetSourceFilterSettingsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetSourceFilterSettingsResponse
func (o *SetSourceFilterSettingsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetSourceFilterSettingsResponse
func (o *SetSourceFilterSettingsResponse) GetError() string {
	return o.Error
}

// SetSourceFilterSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterSettings(
	params *SetSourceFilterSettingsParams,
) (*SetSourceFilterSettingsResponse, error) {
	params.RequestType = "SetSourceFilterSettings"
	data := &SetSourceFilterSettingsResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
