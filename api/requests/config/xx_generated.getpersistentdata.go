// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetPersistentDataParams represents the params body for the "GetPersistentData" request.
Gets the value of a "slot" from the selected persistent data realm.
*/
type GetPersistentDataParams struct {
	requests.ParamsBasic

	// The data realm to select. `OBS_WEBSOCKET_DATA_REALM_GLOBAL` or `OBS_WEBSOCKET_DATA_REALM_PROFILE`
	Realm string `json:"realm,omitempty"`

	// The name of the slot to retrieve data from
	SlotName string `json:"slotName,omitempty"`
}

// GetSelfName just returns "GetPersistentData".
func (o *GetPersistentDataParams) GetSelfName() string {
	return "GetPersistentData"
}

/*
GetPersistentDataResponse represents the response body for the "GetPersistentData" request.
Gets the value of a "slot" from the selected persistent data realm.
*/
type GetPersistentDataResponse struct {
	requests.ResponseBasic

	// Value associated with the slot. `null` if not set
	SlotValue interface{} `json:"slotValue,omitempty"`
}

// GetPersistentData sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetPersistentData(params *GetPersistentDataParams) (*GetPersistentDataResponse, error) {
	data := &GetPersistentDataResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
