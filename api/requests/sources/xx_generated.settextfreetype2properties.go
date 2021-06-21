// This file has been automatically generated. Don't edit it.

package sources

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
SetTextFreetype2PropertiesParams represents the params body for the "SetTextFreetype2Properties" request.
Set the current properties of a Text Freetype 2 source.
Since 4.5.0.
*/
type SetTextFreetype2PropertiesParams struct {
	requests.ParamsBasic

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

	// Source name.
	Source string `json:"source,omitempty"`

	// Text content to be displayed.
	Text string `json:"text,omitempty"`

	// File path.
	TextFile string `json:"text_file,omitempty"`

	// Word wrap.
	WordWrap bool `json:"word_wrap"`
}

// GetSelfName just returns "SetTextFreetype2Properties".
func (o *SetTextFreetype2PropertiesParams) GetSelfName() string {
	return "SetTextFreetype2Properties"
}

/*
SetTextFreetype2PropertiesResponse represents the response body for the "SetTextFreetype2Properties" request.
Set the current properties of a Text Freetype 2 source.
Since v4.5.0.
*/
type SetTextFreetype2PropertiesResponse struct {
	requests.ResponseBasic
}

// SetTextFreetype2Properties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTextFreetype2Properties(
	params *SetTextFreetype2PropertiesParams,
) (*SetTextFreetype2PropertiesResponse, error) {
	data := &SetTextFreetype2PropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
