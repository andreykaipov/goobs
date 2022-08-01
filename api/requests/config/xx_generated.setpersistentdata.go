// This file has been automatically generated. Don't edit it.

package config

/*
SetPersistentDataParams represents the params body for the "SetPersistentData" request.
Sets the value of a "slot" from the selected persistent data realm.
*/
type SetPersistentDataParams struct {
	// The data realm to select. `OBS_WEBSOCKET_DATA_REALM_GLOBAL` or `OBS_WEBSOCKET_DATA_REALM_PROFILE`
	Realm string `json:"realm,omitempty"`

	// The name of the slot to retrieve data from
	SlotName string `json:"slotName,omitempty"`

	// The value to apply to the slot
	SlotValue interface{} `json:"slotValue,omitempty"`
}

/*
SetPersistentDataResponse represents the response body for the "SetPersistentData" request.
Sets the value of a "slot" from the selected persistent data realm.
*/
type SetPersistentDataResponse struct{}

// SetPersistentData sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetPersistentData(params *SetPersistentDataParams) (*SetPersistentDataResponse, error) {
	resp, err := c.SendRequest("SetPersistentData", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetPersistentDataResponse), nil
}
