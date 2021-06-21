// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionDurationParams represents the params body for the "GetTransitionDuration" request.
Get the duration of the currently selected transition if supported.
Since 4.1.0.
*/
type GetTransitionDurationParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetTransitionDuration".
func (o *GetTransitionDurationParams) GetSelfName() string {
	return "GetTransitionDuration"
}

/*
GetTransitionDurationResponse represents the response body for the "GetTransitionDuration" request.
Get the duration of the currently selected transition if supported.
Since v4.1.0.
*/
type GetTransitionDurationResponse struct {
	requests.ResponseBasic

	// Duration of the current transition (in milliseconds).
	TransitionDuration int `json:"transition-duration,omitempty"`
}

// GetTransitionDuration sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionDuration(
	paramss ...*GetTransitionDurationParams,
) (*GetTransitionDurationResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionDurationParams{{}}
	}
	params := paramss[0]
	data := &GetTransitionDurationResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
