// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveFilterFromSourceParams represents the params body for the "RemoveFilterFromSource" request.
Remove a filter from a source
Since 4.5.0.
*/
type RemoveFilterFromSourceParams struct {
	requests.ParamsBasic

	// Name of the filter to remove
	FilterName string `json:"filterName,omitempty"`

	// Name of the source from which the specified filter is removed
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "RemoveFilterFromSource".
func (o *RemoveFilterFromSourceParams) GetSelfName() string {
	return "RemoveFilterFromSource"
}

/*
RemoveFilterFromSourceResponse represents the response body for the "RemoveFilterFromSource" request.
Remove a filter from a source
Since v4.5.0.
*/
type RemoveFilterFromSourceResponse struct {
	requests.ResponseBasic
}

// RemoveFilterFromSource sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveFilterFromSource(params *RemoveFilterFromSourceParams) (*RemoveFilterFromSourceResponse, error) {
	data := &RemoveFilterFromSourceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
