// This file has been automatically generated. Don't edit it.

package general

/*
TriggerHotkeyByNameParams represents the params body for the "TriggerHotkeyByName" request.
Triggers a hotkey using its name. See `GetHotkeyList`
*/
type TriggerHotkeyByNameParams struct {
	// Name of the hotkey to trigger
	HotkeyName string `json:"hotkeyName,omitempty"`
}

/*
TriggerHotkeyByNameResponse represents the response body for the "TriggerHotkeyByName" request.
Triggers a hotkey using its name. See `GetHotkeyList`
*/
type TriggerHotkeyByNameResponse struct{}

// TriggerHotkeyByName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TriggerHotkeyByName(params *TriggerHotkeyByNameParams) (*TriggerHotkeyByNameResponse, error) {
	resp, err := c.SendRequest("TriggerHotkeyByName", params)
	if err != nil {
		return nil, err
	}
	return resp.(*TriggerHotkeyByNameResponse), nil
}
