// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetFilenameFormattingParams represents the params body for the "GetFilenameFormatting" request.
Get the filename formatting string

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingParams struct {
	requests.ParamsBasic
}

// Name just returns "GetFilenameFormatting".
func (o *GetFilenameFormattingParams) Name() string {
	return "GetFilenameFormatting"
}

/*
GetFilenameFormattingResponse represents the response body for the "GetFilenameFormatting" request.
Get the filename formatting string

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingResponse struct {
	requests.ResponseBasic

	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting"`
}

// GetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetFilenameFormatting(
	paramss ...*GetFilenameFormattingParams,
) (*GetFilenameFormattingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetFilenameFormattingParams{{}}
	}
	params := paramss[0]
	data := &GetFilenameFormattingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
