// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetPersistentData request.
type SetPersistentDataParams struct {
	// The data realm to select. `OBS_WEBSOCKET_DATA_REALM_GLOBAL` or `OBS_WEBSOCKET_DATA_REALM_PROFILE`
	Realm string `json:"realm,omitempty"`

	// The name of the slot to retrieve data from
	SlotName string `json:"slotName,omitempty"`

	// The value to apply to the slot
	SlotValue interface{} `json:"slotValue,omitempty"`
}

// Returns the associated request.
func (o *SetPersistentDataParams) GetRequestName() string {
	return "SetPersistentData"
}

// Represents the response body for the SetPersistentData request.
type SetPersistentDataResponse struct{}

// Sets the value of a "slot" from the selected persistent data realm.
func (c *Client) SetPersistentData(params *SetPersistentDataParams) (*SetPersistentDataResponse, error) {
	data := &SetPersistentDataResponse{}
	return data, c.SendRequest(params, data)
}
