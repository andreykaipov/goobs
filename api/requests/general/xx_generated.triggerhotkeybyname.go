// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the TriggerHotkeyByName request.
type TriggerHotkeyByNameParams struct {
	// Name of context of the hotkey to trigger
	ContextName *string `json:"contextName,omitempty"`

	// Name of the hotkey to trigger
	HotkeyName *string `json:"hotkeyName,omitempty"`
}

func NewTriggerHotkeyByNameParams() *TriggerHotkeyByNameParams {
	return &TriggerHotkeyByNameParams{}
}
func (o *TriggerHotkeyByNameParams) WithContextName(x string) *TriggerHotkeyByNameParams {
	o.ContextName = &x
	return o
}
func (o *TriggerHotkeyByNameParams) WithHotkeyName(x string) *TriggerHotkeyByNameParams {
	o.HotkeyName = &x
	return o
}

// Returns the associated request.
func (o *TriggerHotkeyByNameParams) GetRequestName() string {
	return "TriggerHotkeyByName"
}

// Represents the response body for the TriggerHotkeyByName request.
type TriggerHotkeyByNameResponse struct {
	_response
}

/*
Triggers a hotkey using its name. See `GetHotkeyList`.

Note: Hotkey functionality in obs-websocket comes as-is, and we do not guarantee support if things are broken. In 9/10 usages of hotkey requests, there exists a better, more reliable method via other requests.
*/
func (c *Client) TriggerHotkeyByName(params *TriggerHotkeyByNameParams) (*TriggerHotkeyByNameResponse, error) {
	data := &TriggerHotkeyByNameResponse{}
	return data, c.client.SendRequest(params, data)
}
