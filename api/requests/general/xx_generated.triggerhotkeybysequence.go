// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TriggerHotkeyBySequenceParams represents the params body for the "TriggerHotkeyBySequence" request.
Executes hotkey routine, identified by bound combination of keys. A single key combination might trigger multiple hotkey routines depending on user settings
Since 4.9.0.
*/
type TriggerHotkeyBySequenceParams struct {
	requests.ParamsBasic

	// Main key identifier (e.g. `OBS_KEY_A` for key "A"). Available identifiers
	// [here](https://github.com/obsproject/obs-studio/blob/master/libobs/obs-hotkeys.h)
	KeyId string `json:"keyId,omitempty"`

	KeyModifiers *KeyModifiers `json:"keyModifiers,omitempty"`
}

type KeyModifiers struct {
	// Trigger Alt Key
	Alt bool `json:"alt"`

	// Trigger Command Key (Mac)
	Command bool `json:"command"`

	// Trigger Control (Ctrl) Key
	Control bool `json:"control"`

	// Trigger Shift Key
	Shift bool `json:"shift"`
}

// GetSelfName just returns "TriggerHotkeyBySequence".
func (o *TriggerHotkeyBySequenceParams) GetSelfName() string {
	return "TriggerHotkeyBySequence"
}

/*
TriggerHotkeyBySequenceResponse represents the response body for the "TriggerHotkeyBySequence" request.
Executes hotkey routine, identified by bound combination of keys. A single key combination might trigger multiple hotkey routines depending on user settings
Since v4.9.0.
*/
type TriggerHotkeyBySequenceResponse struct {
	requests.ResponseBasic
}

// TriggerHotkeyBySequence sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TriggerHotkeyBySequence(
	params *TriggerHotkeyBySequenceParams,
) (*TriggerHotkeyBySequenceResponse, error) {
	data := &TriggerHotkeyBySequenceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
