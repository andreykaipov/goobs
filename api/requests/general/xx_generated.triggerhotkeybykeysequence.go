// This file has been automatically generated. Don't edit it.

package general

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the TriggerHotkeyByKeySequence request.
type TriggerHotkeyByKeySequenceParams struct {
	// The OBS key ID to use. See https://github.com/obsproject/obs-studio/blob/master/libobs/obs-hotkeys.h
	KeyId *string `json:"keyId,omitempty"`

	// Object containing key modifiers to apply
	KeyModifiers *typedefs.KeyModifiers `json:"keyModifiers,omitempty"`
}

func NewTriggerHotkeyByKeySequenceParams() *TriggerHotkeyByKeySequenceParams {
	return &TriggerHotkeyByKeySequenceParams{}
}
func (o *TriggerHotkeyByKeySequenceParams) WithKeyId(x string) *TriggerHotkeyByKeySequenceParams {
	o.KeyId = &x
	return o
}

func (o *TriggerHotkeyByKeySequenceParams) WithKeyModifiers(
	x *typedefs.KeyModifiers,
) *TriggerHotkeyByKeySequenceParams {
	o.KeyModifiers = x
	return o
}

// Returns the associated request.
func (o *TriggerHotkeyByKeySequenceParams) GetRequestName() string {
	return "TriggerHotkeyByKeySequence"
}

// Represents the response body for the TriggerHotkeyByKeySequence request.
type TriggerHotkeyByKeySequenceResponse struct {
	_response
}

/*
Triggers a hotkey using a sequence of keys.

Note: Hotkey functionality in obs-websocket comes as-is, and we do not guarantee support if things are broken. In 9/10 usages of hotkey requests, there exists a better, more reliable method via other requests.
*/
func (c *Client) TriggerHotkeyByKeySequence(
	paramss ...*TriggerHotkeyByKeySequenceParams,
) (*TriggerHotkeyByKeySequenceResponse, error) {
	if len(paramss) == 0 {
		paramss = []*TriggerHotkeyByKeySequenceParams{{}}
	}
	params := paramss[0]
	data := &TriggerHotkeyByKeySequenceResponse{}
	return data, c.client.SendRequest(params, data)
}
