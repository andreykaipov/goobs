// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TriggerHotkeyByNameParams represents the params body for the "TriggerHotkeyByName" request.
Executes hotkey routine, identified by hotkey unique name
Since 4.9.0.
*/
type TriggerHotkeyByNameParams struct {
	requests.ParamsBasic

	// Unique name of the hotkey, as defined when registering the hotkey (e.g. "ReplayBuffer.Save")
	HotkeyName string `json:"hotkeyName,omitempty"`
}

// GetSelfName just returns "TriggerHotkeyByName".
func (o *TriggerHotkeyByNameParams) GetSelfName() string {
	return "TriggerHotkeyByName"
}

/*
TriggerHotkeyByNameResponse represents the response body for the "TriggerHotkeyByName" request.
Executes hotkey routine, identified by hotkey unique name
Since v4.9.0.
*/
type TriggerHotkeyByNameResponse struct {
	requests.ResponseBasic
}

// TriggerHotkeyByName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TriggerHotkeyByName(params *TriggerHotkeyByNameParams) (*TriggerHotkeyByNameResponse, error) {
	data := &TriggerHotkeyByNameResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
