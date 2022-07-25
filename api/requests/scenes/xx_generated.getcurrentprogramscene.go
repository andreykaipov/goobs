// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentProgramSceneParams represents the params body for the "GetCurrentProgramScene" request.
Gets the current program scene.
*/
type GetCurrentProgramSceneParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetCurrentProgramScene".
func (o *GetCurrentProgramSceneParams) GetSelfName() string {
	return "GetCurrentProgramScene"
}

/*
GetCurrentProgramSceneResponse represents the response body for the "GetCurrentProgramScene" request.
Gets the current program scene.
*/
type GetCurrentProgramSceneResponse struct {
	requests.ResponseBasic

	// Current program scene
	CurrentProgramSceneName string `json:"currentProgramSceneName,omitempty"`
}

// GetCurrentProgramScene sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentProgramScene(
	paramss ...*GetCurrentProgramSceneParams,
) (*GetCurrentProgramSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentProgramSceneParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentProgramSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
