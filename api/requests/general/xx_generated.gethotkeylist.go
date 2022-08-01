// This file has been automatically generated. Don't edit it.

package general

/*
GetHotkeyListParams represents the params body for the "GetHotkeyList" request.
Gets an array of all hotkey names in OBS
*/
type GetHotkeyListParams struct{}

/*
GetHotkeyListResponse represents the response body for the "GetHotkeyList" request.
Gets an array of all hotkey names in OBS
*/
type GetHotkeyListResponse struct {
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
	resp, err := c.SendRequest("GetHotkeyList", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetHotkeyListResponse), nil
}
