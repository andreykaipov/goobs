// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourcesTypesListParams represents the params body for the "GetSourcesTypesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetSourcesTypesListParams
func (o *GetSourcesTypesListParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSourcesTypesListParams
func (o *GetSourcesTypesListParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSourcesTypesListParams
func (o *GetSourcesTypesListParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourcesTypesListResponse represents the response body for the "GetSourcesTypesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListResponse struct {
	requests.Response

	Ids []struct {
		Caps struct {
			// True if interaction with this sources of this type is possible
			CanInteract bool `json:"canInteract"`

			// True if sources of this type should not be fully duplicated
			DoNotDuplicate bool `json:"doNotDuplicate"`

			// True if sources of this type may cause a feedback loop if it's audio is monitored and
			// shouldn't be
			DoNotSelfMonitor bool `json:"doNotSelfMonitor"`

			// True if sources of this type provide audio
			HasAudio bool `json:"hasAudio"`

			// True if sources of this type provide video
			HasVideo bool `json:"hasVideo"`

			// True if source of this type provide frames asynchronously
			IsAsync bool `json:"isAsync"`

			// True if sources of this type composite one or more sub-sources
			IsComposite bool `json:"isComposite"`
		} `json:"caps"`

		// Default settings of this source type
		DefaultSettings map[string]interface{} `json:"defaultSettings"`

		// Display name of the source type
		DisplayName string `json:"displayName"`

		// Type. Value is one of the following: "input", "filter", "transition" or "other"
		Type string `json:"type"`

		// Non-unique internal source type ID
		TypeId string `json:"typeId"`
	} `json:"ids"`
}

// GetMessageID returns the MessageID of GetSourcesTypesListResponse
func (o *GetSourcesTypesListResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSourcesTypesListResponse
func (o *GetSourcesTypesListResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSourcesTypesListResponse
func (o *GetSourcesTypesListResponse) GetError() string {
	return o.Error
}

// GetSourcesTypesList sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetSourcesTypesList(
	paramss ...*GetSourcesTypesListParams,
) (*GetSourcesTypesListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSourcesTypesListParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetSourcesTypesList"
	data := &GetSourcesTypesListResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
