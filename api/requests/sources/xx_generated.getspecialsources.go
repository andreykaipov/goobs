// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSpecialSourcesParams represents the params body for the "GetSpecialSources" request.
Get configured special sources like Desktop Audio and Mic/Aux sources.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetSpecialSources.
*/
type GetSpecialSourcesParams struct {
	requests.ParamsBasic
}

// Name just returns "GetSpecialSources".
func (o *GetSpecialSourcesParams) Name() string {
	return "GetSpecialSources"
}

/*
GetSpecialSourcesResponse represents the response body for the "GetSpecialSources" request.
Get configured special sources like Desktop Audio and Mic/Aux sources.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetSpecialSources.
*/
type GetSpecialSourcesResponse struct {
	requests.ResponseBasic

	// Name of the first Desktop Audio capture source.
	Desktop1 string `json:"desktop-1"`

	// Name of the second Desktop Audio capture source.
	Desktop2 string `json:"desktop-2"`

	// Name of the first Mic/Aux input source.
	Mic1 string `json:"mic-1"`

	// Name of the second Mic/Aux input source.
	Mic2 string `json:"mic-2"`

	// NAme of the third Mic/Aux input source.
	Mic3 string `json:"mic-3"`
}

// GetSpecialSources sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetSpecialSources(
	paramss ...*GetSpecialSourcesParams,
) (*GetSpecialSourcesResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSpecialSourcesParams{{}}
	}
	params := paramss[0]
	data := &GetSpecialSourcesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
