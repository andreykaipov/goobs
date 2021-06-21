// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RefreshBrowserSourceParams represents the params body for the "RefreshBrowserSource" request.
Refreshes the specified browser source.
Since 4.9.0.
*/
type RefreshBrowserSourceParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "RefreshBrowserSource".
func (o *RefreshBrowserSourceParams) GetSelfName() string {
	return "RefreshBrowserSource"
}

/*
RefreshBrowserSourceResponse represents the response body for the "RefreshBrowserSource" request.
Refreshes the specified browser source.
Since v4.9.0.
*/
type RefreshBrowserSourceResponse struct {
	requests.ResponseBasic
}

// RefreshBrowserSource sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RefreshBrowserSource(params *RefreshBrowserSourceParams) (*RefreshBrowserSourceResponse, error) {
	data := &RefreshBrowserSourceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
