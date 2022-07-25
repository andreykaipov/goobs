// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetHotkeyListParams represents the params body for the "GetHotkeyList" request.
Gets an array of all hotkey names in OBS
*/
type GetHotkeyListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetHotkeyList".
func (o *GetHotkeyListParams) GetSelfName() string {
	return "GetHotkeyList"
}

/*
GetHotkeyListResponse represents the response body for the "GetHotkeyList" request.
Gets an array of all hotkey names in OBS
*/
type GetHotkeyListResponse struct {
	requests.ResponseBasic

	// Array of hotkey names
	Hotkeys []string `json:"hotkeys,omitempty"`
}

// GetHotkeyList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetHotkeyList(paramss ...*GetHotkeyListParams) (*GetHotkeyListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetHotkeyListParams{{}}
	}
	params := paramss[0]
	data := &GetHotkeyListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
