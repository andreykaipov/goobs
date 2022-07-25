// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetPersistentDataParams represents the params body for the "SetPersistentData" request.
Sets the value of a "slot" from the selected persistent data realm.
*/
type SetPersistentDataParams struct {
	requests.ParamsBasic

	// The data realm to select. `OBS_WEBSOCKET_DATA_REALM_GLOBAL` or `OBS_WEBSOCKET_DATA_REALM_PROFILE`
	Realm string `json:"realm,omitempty"`

	// The name of the slot to retrieve data from
	SlotName string `json:"slotName,omitempty"`

	// The value to apply to the slot
	SlotValue interface{} `json:"slotValue,omitempty"`
}

// GetSelfName just returns "SetPersistentData".
func (o *SetPersistentDataParams) GetSelfName() string {
	return "SetPersistentData"
}

/*
SetPersistentDataResponse represents the response body for the "SetPersistentData" request.
Sets the value of a "slot" from the selected persistent data realm.
*/
type SetPersistentDataResponse struct {
	requests.ResponseBasic
}

// SetPersistentData sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetPersistentData(params *SetPersistentDataParams) (*SetPersistentDataResponse, error) {
	data := &SetPersistentDataResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
