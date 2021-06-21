// This file has been automatically generated. Don't edit it.

package sources

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetTextFreetype2PropertiesParams represents the params body for the "GetTextFreetype2Properties" request.
Get the current properties of a Text Freetype 2 source.
Since 4.5.0.
*/
type GetTextFreetype2PropertiesParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source,omitempty"`
}

// GetSelfName just returns "GetTextFreetype2Properties".
func (o *GetTextFreetype2PropertiesParams) GetSelfName() string {
	return "GetTextFreetype2Properties"
}

/*
GetTextFreetype2PropertiesResponse represents the response body for the "GetTextFreetype2Properties" request.
Get the current properties of a Text Freetype 2 source.
Since v4.5.0.
*/
type GetTextFreetype2PropertiesResponse struct {
	requests.ResponseBasic

	// Gradient top color.
	Color1 int `json:"color1,omitempty"`

	// Gradient bottom color.
	Color2 int `json:"color2,omitempty"`

	// Custom width (0 to disable).
	CustomWidth int `json:"custom_width,omitempty"`

	// Drop shadow.
	DropShadow bool `json:"drop_shadow"`

	// The font specification for this object.
	Font *typedefs.Font `json:"font,omitempty"`

	// Read text from the specified file.
	FromFile bool `json:"from_file"`

	// Chat log.
	LogMode bool `json:"log_mode"`

	// Outline.
	Outline bool `json:"outline"`

	// Source name
	Source string `json:"source,omitempty"`

	// Text content to be displayed.
	Text string `json:"text,omitempty"`

	// File path.
	TextFile string `json:"text_file,omitempty"`

	// Word wrap.
	WordWrap bool `json:"word_wrap"`
}

// GetTextFreetype2Properties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetTextFreetype2Properties(
	params *GetTextFreetype2PropertiesParams,
) (*GetTextFreetype2PropertiesResponse, error) {
	data := &GetTextFreetype2PropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
