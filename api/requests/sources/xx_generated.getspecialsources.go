// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSpecialSourcesParams represents the params body for the "GetSpecialSources" request.
Get configured special sources like Desktop Audio and Mic/Aux sources.
Since 4.1.0.
*/
type GetSpecialSourcesParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetSpecialSources".
func (o *GetSpecialSourcesParams) GetSelfName() string {
	return "GetSpecialSources"
}

/*
GetSpecialSourcesResponse represents the response body for the "GetSpecialSources" request.
Get configured special sources like Desktop Audio and Mic/Aux sources.
Since v4.1.0.
*/
type GetSpecialSourcesResponse struct {
	requests.ResponseBasic

	// Name of the first Desktop Audio capture source.
	Desktop1 string `json:"desktop-1,omitempty"`

	// Name of the second Desktop Audio capture source.
	Desktop2 string `json:"desktop-2,omitempty"`

	// Name of the first Mic/Aux input source.
	Mic1 string `json:"mic-1,omitempty"`

	// Name of the second Mic/Aux input source.
	Mic2 string `json:"mic-2,omitempty"`

	// NAme of the third Mic/Aux input source.
	Mic3 string `json:"mic-3,omitempty"`
}

// GetSpecialSources sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetSpecialSources(paramss ...*GetSpecialSourcesParams) (*GetSpecialSourcesResponse, error) {
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
