// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourcesTypesListParams represents the params body for the "GetSourcesTypesList" request.
Get a list of all available sources types

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListParams struct {
	requests.ParamsBasic
}

// Name just returns "GetSourcesTypesList".
func (o *GetSourcesTypesListParams) Name() string {
	return "GetSourcesTypesList"
}

/*
GetSourcesTypesListResponse represents the response body for the "GetSourcesTypesList" request.
Get a list of all available sources types

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListResponse struct {
	requests.ResponseBasic

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

// GetSourcesTypesList sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetSourcesTypesList(
	paramss ...*GetSourcesTypesListParams,
) (*GetSourcesTypesListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSourcesTypesListParams{{}}
	}
	params := paramss[0]
	data := &GetSourcesTypesListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
