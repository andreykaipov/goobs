// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the TriggerHotkeyByName request.
type TriggerHotkeyByNameParams struct {
	// Name of the hotkey to trigger
	HotkeyName string `json:"hotkeyName,omitempty"`
}

func NewTriggerHotkeyByNameParams() *TriggerHotkeyByNameParams {
	return &TriggerHotkeyByNameParams{}
}
func (o *TriggerHotkeyByNameParams) WithHotkeyName(x string) *TriggerHotkeyByNameParams {
	o.HotkeyName = x
	return o
}

// Returns the associated request.
func (o *TriggerHotkeyByNameParams) GetRequestName() string {
	return "TriggerHotkeyByName"
}

// Represents the response body for the TriggerHotkeyByName request.
type TriggerHotkeyByNameResponse struct{}

// Triggers a hotkey using its name. See `GetHotkeyList`
func (c *Client) TriggerHotkeyByName(params *TriggerHotkeyByNameParams) (*TriggerHotkeyByNameResponse, error) {
	data := &TriggerHotkeyByNameResponse{}
	return data, c.SendRequest(params, data)
}
