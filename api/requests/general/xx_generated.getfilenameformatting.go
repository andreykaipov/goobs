// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetFilenameFormattingParams represents the params body for the "GetFilenameFormatting" request.
Get the filename formatting string
Since 4.3.0.
*/
type GetFilenameFormattingParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetFilenameFormatting".
func (o *GetFilenameFormattingParams) GetSelfName() string {
	return "GetFilenameFormatting"
}

/*
GetFilenameFormattingResponse represents the response body for the "GetFilenameFormatting" request.
Get the filename formatting string
Since v4.3.0.
*/
type GetFilenameFormattingResponse struct {
	requests.ResponseBasic

	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting,omitempty"`
}

// GetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
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
