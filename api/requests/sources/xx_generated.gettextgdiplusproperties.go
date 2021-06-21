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
	Source string `json:"source,omitempty"`
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
	Align string `json:"align,omitempty"`

	// Background color.
	BkColor int `json:"bk_color,omitempty"`

	// Background opacity (0-100).
	BkOpacity int `json:"bk_opacity,omitempty"`

	// Chat log.
	Chatlog bool `json:"chatlog"`

	// Chat log lines.
	ChatlogLines int `json:"chatlog_lines,omitempty"`

	// Text color.
	Color int `json:"color,omitempty"`

	// Extents wrap.
	Extents bool `json:"extents"`

	// Extents cx.
	ExtentsCx int `json:"extents_cx,omitempty"`

	// Extents cy.
	ExtentsCy int `json:"extents_cy,omitempty"`

	// File path name.
	File string `json:"file,omitempty"`

	// The font specification for this object.
	Font *typedefs.Font `json:"font,omitempty"`

	// Gradient enabled.
	Gradient bool `json:"gradient"`

	// Gradient color.
	GradientColor int `json:"gradient_color,omitempty"`

	// Gradient direction.
	GradientDir float64 `json:"gradient_dir,omitempty"`

	// Gradient opacity (0-100).
	GradientOpacity int `json:"gradient_opacity,omitempty"`

	// Outline.
	Outline bool `json:"outline"`

	// Outline color.
	OutlineColor int `json:"outline_color,omitempty"`

	// Outline opacity (0-100).
	OutlineOpacity int `json:"outline_opacity,omitempty"`

	// Outline size.
	OutlineSize int `json:"outline_size,omitempty"`

	// Read text from the specified file.
	ReadFromFile bool `json:"read_from_file"`

	// Source name.
	Source string `json:"source,omitempty"`

	// Text content to be displayed.
	Text string `json:"text,omitempty"`

	// Text vertical alignment ("top", "center", "bottom").
	Valign string `json:"valign,omitempty"`

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
