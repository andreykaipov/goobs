// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTextFreetype2PropertiesParams represents the params body for the "GetTextFreetype2Properties" request.
Get the current properties of a Text Freetype 2 source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties.
*/
type GetTextFreetype2PropertiesParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`
}

// Name just returns "GetTextFreetype2Properties".
func (o *GetTextFreetype2PropertiesParams) Name() string {
	return "GetTextFreetype2Properties"
}

/*
GetTextFreetype2PropertiesResponse represents the response body for the "GetTextFreetype2Properties" request.
Get the current properties of a Text Freetype 2 source.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties.
*/
type GetTextFreetype2PropertiesResponse struct {
	requests.ResponseBasic

	// Gradient top color.
	Color1 int `json:"color1"`

	// Gradient bottom color.
	Color2 int `json:"color2"`

	// Custom width (0 to disable).
	CustomWidth int `json:"custom_width"`

	// Drop shadow.
	DropShadow bool `json:"drop_shadow"`

	Font struct {
		// Font face.
		Face string `json:"face"`

		// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3, Underline=5, Strikeout=8`
		Flags int `json:"flags"`

		// Font text size.
		Size int `json:"size"`

		// Font Style (unknown function).
		Style string `json:"style"`
	} `json:"font"`

	// Read text from the specified file.
	FromFile bool `json:"from_file"`

	// Chat log.
	LogMode bool `json:"log_mode"`

	// Outline.
	Outline bool `json:"outline"`

	// Source name
	Source string `json:"source"`

	// Text content to be displayed.
	Text string `json:"text"`

	// File path.
	TextFile string `json:"text_file"`

	// Word wrap.
	WordWrap bool `json:"word_wrap"`
}

// GetTextFreetype2Properties sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) GetTextFreetype2Properties(
	params *GetTextFreetype2PropertiesParams,
) (*GetTextFreetype2PropertiesResponse, error) {
	data := &GetTextFreetype2PropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
