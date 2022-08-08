// This file has been automatically generated. Don't edit it.

package general

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the TriggerHotkeyByKeySequence request.
type TriggerHotkeyByKeySequenceParams struct {
	// The OBS key ID to use. See https://github.com/obsproject/obs-studio/blob/master/libobs/obs-hotkeys.h
	KeyId string `json:"keyId,omitempty"`

	// Key modifiers to apply
	KeyModifiers *typedefs.KeyModifiers `json:"keyModifiers,omitempty"`
}

// Returns the associated request.
func (o *TriggerHotkeyByKeySequenceParams) GetRequestName() string {
	return "TriggerHotkeyByKeySequence"
}

// Represents the response body for the TriggerHotkeyByKeySequence request.
type TriggerHotkeyByKeySequenceResponse struct{}

// Triggers a hotkey using a sequence of keys.
func (c *Client) TriggerHotkeyByKeySequence(
	paramss ...*TriggerHotkeyByKeySequenceParams,
) (*TriggerHotkeyByKeySequenceResponse, error) {
	if len(paramss) == 0 {
		paramss = []*TriggerHotkeyByKeySequenceParams{{}}
	}
	params := paramss[0]
	data := &TriggerHotkeyByKeySequenceResponse{}
	return data, c.SendRequest(params, data)
}
