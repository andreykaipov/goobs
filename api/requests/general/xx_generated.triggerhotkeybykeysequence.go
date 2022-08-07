// This file has been automatically generated. Don't edit it.

package general

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
TriggerHotkeyByKeySequenceParams represents the params body for the "TriggerHotkeyByKeySequence" request.
Triggers a hotkey using a sequence of keys.
*/
type TriggerHotkeyByKeySequenceParams struct {
	// The OBS key ID to use. See https://github.com/obsproject/obs-studio/blob/master/libobs/obs-hotkeys.h
	KeyId string `json:"keyId,omitempty"`

	// Key modifiers to apply
	KeyModifiers *typedefs.KeyModifiers `json:"keyModifiers,omitempty"`
}

/*
TriggerHotkeyByKeySequenceResponse represents the response body for the "TriggerHotkeyByKeySequence" request.
Triggers a hotkey using a sequence of keys.
*/
type TriggerHotkeyByKeySequenceResponse struct{}

// TriggerHotkeyByKeySequence sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) TriggerHotkeyByKeySequence(
	paramss ...*TriggerHotkeyByKeySequenceParams,
) (*TriggerHotkeyByKeySequenceResponse, error) {
	if len(paramss) == 0 {
		paramss = []*TriggerHotkeyByKeySequenceParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("TriggerHotkeyByKeySequence", params)
	if err != nil {
		return nil, err
	}
	return resp.(*TriggerHotkeyByKeySequenceResponse), nil
}
