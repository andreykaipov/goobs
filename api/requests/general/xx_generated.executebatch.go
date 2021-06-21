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

	Requests []*Request `json:"requests,omitempty"`
}

type Request struct {
	// ID of the individual request. Can be any string and not required to be unique. Defaults to empty string if not
	// specified.
	MessageId string `json:"message-id,omitempty"`

	// Request type. Eg. `GetVersion`.
	RequestType string `json:"request-type,omitempty"`
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

	Results []*Result `json:"results,omitempty"`
}

type Result struct {
	// Error message accompanying an `error` status.
	Error string `json:"error,omitempty"`

	// ID of the individual request which was originally provided by the client.
	MessageId string `json:"message-id,omitempty"`

	// Status response as string. Either `ok` or `error`.
	Status string `json:"status,omitempty"`
}

// ExecuteBatch sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ExecuteBatch(params *ExecuteBatchParams) (*ExecuteBatchResponse, error) {
	data := &ExecuteBatchResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
