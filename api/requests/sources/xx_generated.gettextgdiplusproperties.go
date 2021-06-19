// This file has been automatically generated. Don't edit it.

package sources

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetTextGDIPlusPropertiesParams represents the params body for the "GetTextGDIPlusProperties" request.
Get the current properties of a Text GDI Plus source.
Since 4.1.0.
*/
type GetTextGDIPlusPropertiesParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`
}

// GetSelfName just returns "GetTextGDIPlusProperties".
func (o *GetTextGDIPlusPropertiesParams) GetSelfName() string {
	return "GetTextGDIPlusProperties"
}

/*
GetTextGDIPlusPropertiesResponse represents the response body for the "GetTextGDIPlusProperties" request.
Get the current properties of a Text GDI Plus source.
Since v4.1.0.
*/
type GetTextGDIPlusPropertiesResponse struct {
	requests.ResponseBasic

	// Text Alignment ("left", "center", "right").
	Align string `json:"align"`

	// Background color.
	BkColor int `json:"bk_color"`

	// Background opacity (0-100).
	BkOpacity int `json:"bk_opacity"`

	// Chat log.
	Chatlog bool `json:"chatlog"`

	// Chat log lines.
	ChatlogLines int `json:"chatlog_lines"`

	// Text color.
	Color int `json:"color"`

	// Extents wrap.
	Extents bool `json:"extents"`

	// Extents cx.
	ExtentsCx int `json:"extents_cx"`

	// Extents cy.
	ExtentsCy int `json:"extents_cy"`

	// File path name.
	File string `json:"file"`

	// The font specification for this object.
	Font typedefs.Font `json:"font"`

	// Gradient enabled.
	Gradient bool `json:"gradient"`

	// Gradient color.
	GradientColor int `json:"gradient_color"`

	// Gradient direction.
	GradientDir float64 `json:"gradient_dir"`

	// Gradient opacity (0-100).
	GradientOpacity int `json:"gradient_opacity"`

	// Outline.
	Outline bool `json:"outline"`

	// Outline color.
	OutlineColor int `json:"outline_color"`

	// Outline opacity (0-100).
	OutlineOpacity int `json:"outline_opacity"`

	// Outline size.
	OutlineSize int `json:"outline_size"`

	// Read text from the specified file.
	ReadFromFile bool `json:"read_from_file"`

	// Source name.
	Source string `json:"source"`

	// Text content to be displayed.
	Text string `json:"text"`

	// Text vertical alignment ("top", "center", "bottom").
	Valign string `json:"valign"`

	// Vertical text enabled.
	Vertical bool `json:"vertical"`
}

// GetTextGDIPlusProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetTextGDIPlusProperties(
	params *GetTextGDIPlusPropertiesParams,
) (*GetTextGDIPlusPropertiesResponse, error) {
	data := &GetTextGDIPlusPropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
