// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TriggerHotkeyByKeySequenceParams represents the params body for the "TriggerHotkeyByKeySequence" request.
Triggers a hotkey using a sequence of keys.
*/
type TriggerHotkeyByKeySequenceParams struct {
	requests.ParamsBasic

	// The OBS key ID to use. See https://github.com/obsproject/obs-studio/blob/master/libobs/obs-hotkeys.h
	KeyId string `json:"keyId,omitempty"`

	KeyModifiers *KeyModifiers `json:"keyModifiers,omitempty"`
}

type KeyModifiers struct {
	// Press ALT
	Alt bool `json:"alt,omitempty"`

	// Press CMD (Mac)
	Command bool `json:"command,omitempty"`

	// Press CTRL
	Control bool `json:"control,omitempty"`

	// Press Shift
	Shift bool `json:"shift,omitempty"`
}

// GetSelfName just returns "TriggerHotkeyByKeySequence".
func (o *TriggerHotkeyByKeySequenceParams) GetSelfName() string {
	return "TriggerHotkeyByKeySequence"
}

/*
TriggerHotkeyByKeySequenceResponse represents the response body for the "TriggerHotkeyByKeySequence" request.
Triggers a hotkey using a sequence of keys.
*/
type TriggerHotkeyByKeySequenceResponse struct {
	requests.ResponseBasic
}

// TriggerHotkeyByKeySequence sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TriggerHotkeyByKeySequence(
	params *TriggerHotkeyByKeySequenceParams,
) (*TriggerHotkeyByKeySequenceResponse, error) {
	data := &TriggerHotkeyByKeySequenceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
