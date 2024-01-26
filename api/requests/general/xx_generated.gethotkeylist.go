// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the GetHotkeyList request.
type GetHotkeyListParams struct{}

// Returns the associated request.
func (o *GetHotkeyListParams) GetRequestName() string {
	return "GetHotkeyList"
}

// Represents the response body for the GetHotkeyList request.
type GetHotkeyListResponse struct {
	_response

	// Array of hotkey names
	Hotkeys []string `json:"hotkeys,omitempty"`
}

/*
Gets an array of all hotkey names in OBS.

Note: Hotkey functionality in obs-websocket comes as-is, and we do not guarantee support if things are broken. In 9/10 usages of hotkey requests, there exists a better, more reliable method via other requests.
*/
func (c *Client) GetHotkeyList(paramss ...*GetHotkeyListParams) (*GetHotkeyListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetHotkeyListParams{{}}
	}
	params := paramss[0]
	data := &GetHotkeyListResponse{}
	return data, c.client.SendRequest(params, data)
}
