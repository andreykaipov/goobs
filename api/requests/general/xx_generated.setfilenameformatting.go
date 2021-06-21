// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetFilenameFormattingParams represents the params body for the "SetFilenameFormatting" request.
Set the filename formatting string
Since 4.3.0.
*/
type SetFilenameFormattingParams struct {
	requests.ParamsBasic

	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting,omitempty"`
}

// GetSelfName just returns "SetFilenameFormatting".
func (o *SetFilenameFormattingParams) GetSelfName() string {
	return "SetFilenameFormatting"
}

/*
SetFilenameFormattingResponse represents the response body for the "SetFilenameFormatting" request.
Set the filename formatting string
Since v4.3.0.
*/
type SetFilenameFormattingResponse struct {
	requests.ResponseBasic
}

// SetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetFilenameFormatting(params *SetFilenameFormattingParams) (*SetFilenameFormattingResponse, error) {
	data := &SetFilenameFormattingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
