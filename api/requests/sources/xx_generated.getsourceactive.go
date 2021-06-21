// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceActiveParams represents the params body for the "GetSourceActive" request.
Get the source's active status of a specified source (if it is showing in the final mix).
Since 4.9.1.
*/
type GetSourceActiveParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetSourceActive".
func (o *GetSourceActiveParams) GetSelfName() string {
	return "GetSourceActive"
}

/*
GetSourceActiveResponse represents the response body for the "GetSourceActive" request.
Get the source's active status of a specified source (if it is showing in the final mix).
Since v4.9.1.
*/
type GetSourceActiveResponse struct {
	requests.ResponseBasic

	// Source active status of the source.
	SourceActive bool `json:"sourceActive"`
}

// GetSourceActive sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceActive(params *GetSourceActiveParams) (*GetSourceActiveResponse, error) {
	data := &GetSourceActiveResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
