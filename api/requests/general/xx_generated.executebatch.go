// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ExecuteBatchParams represents the params body for the "ExecuteBatch" request.
Executes a list of requests sequentially (one-by-one on the same thread).
Since 4.9.0.
*/
type ExecuteBatchParams struct {
	requests.ParamsBasic

	// Stop processing batch requests if one returns a failure.
	AbortOnFail bool `json:"abortOnFail"`

	Requests []struct {
		// ID of the individual request. Can be any string and not required to be unique. Defaults
		// to empty string if not specified.
		MessageId string `json:"message-id"`

		// Request type. Eg. `GetVersion`.
		RequestType string `json:"request-type"`
	} `json:"requests"`
}

// GetSelfName just returns "ExecuteBatch".
func (o *ExecuteBatchParams) GetSelfName() string {
	return "ExecuteBatch"
}

/*
ExecuteBatchResponse represents the response body for the "ExecuteBatch" request.
Executes a list of requests sequentially (one-by-one on the same thread).
Since v4.9.0.
*/
type ExecuteBatchResponse struct {
	requests.ResponseBasic

	Results []struct {
		// Error message accompanying an `error` status.
		Error string `json:"error"`

		// ID of the individual request which was originally provided by the client.
		MessageId string `json:"message-id"`

		// Status response as string. Either `ok` or `error`.
		Status string `json:"status"`
	} `json:"results"`
}

// ExecuteBatch sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ExecuteBatch(params *ExecuteBatchParams) (*ExecuteBatchResponse, error) {
	data := &ExecuteBatchResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
