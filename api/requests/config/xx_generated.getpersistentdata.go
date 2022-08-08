// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the GetPersistentData request.
type GetPersistentDataParams struct {
	// The data realm to select. `OBS_WEBSOCKET_DATA_REALM_GLOBAL` or `OBS_WEBSOCKET_DATA_REALM_PROFILE`
	Realm string `json:"realm,omitempty"`

	// The name of the slot to retrieve data from
	SlotName string `json:"slotName,omitempty"`
}

// Returns the associated request.
func (o *GetPersistentDataParams) GetRequestName() string {
	return "GetPersistentData"
}

// Represents the response body for the GetPersistentData request.
type GetPersistentDataResponse struct {
	// Value associated with the slot. `null` if not set
	SlotValue interface{} `json:"slotValue,omitempty"`
}

// Gets the value of a "slot" from the selected persistent data realm.
func (c *Client) GetPersistentData(params *GetPersistentDataParams) (*GetPersistentDataResponse, error) {
	data := &GetPersistentDataResponse{}
	return data, c.SendRequest(params, data)
}
