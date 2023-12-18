// This file has been automatically generated. Don't edit it.

package general

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the GetHotkeyList request.
type GetHotkeyListParams struct{}

// Returns the associated request.
func (o *GetHotkeyListParams) GetRequestName() string {
	return "GetHotkeyList"
}

// Represents the response body for the GetHotkeyList request.
type GetHotkeyListResponse struct {
	api.ResponseCommon

	// Array of hotkey names
	Hotkeys []string `json:"hotkeys,omitempty"`
}

// Gets an array of all hotkey names in OBS
func (c *Client) GetHotkeyList(paramss ...*GetHotkeyListParams) (*GetHotkeyListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetHotkeyListParams{{}}
	}
	params := paramss[0]
	data := &GetHotkeyListResponse{}
	return data, c.SendRequest(params, data)
}
