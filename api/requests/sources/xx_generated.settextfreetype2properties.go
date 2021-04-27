// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTextFreetype2PropertiesParams represents the params body for the "SetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties.
*/
type SetTextFreetype2PropertiesParams struct {
	requests.Params

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

	// Source name.
	Source string `json:"source"`

	// Text content to be displayed.
	Text string `json:"text"`

	// File path.
	TextFile string `json:"text_file"`

	// Word wrap.
	WordWrap bool `json:"word_wrap"`
}

// GetRequestType returns the RequestType of SetTextFreetype2PropertiesParams
func (o *SetTextFreetype2PropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetTextFreetype2PropertiesParams
func (o *SetTextFreetype2PropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetTextFreetype2PropertiesParams
func (o *SetTextFreetype2PropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetTextFreetype2PropertiesResponse represents the response body for the "SetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties.
*/
type SetTextFreetype2PropertiesResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetTextFreetype2PropertiesResponse
func (o *SetTextFreetype2PropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetTextFreetype2PropertiesResponse
func (o *SetTextFreetype2PropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetTextFreetype2PropertiesResponse
func (o *SetTextFreetype2PropertiesResponse) GetError() string {
	return o.Error
}

// SetTextFreetype2Properties sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) SetTextFreetype2Properties(
	params *SetTextFreetype2PropertiesParams,
) (*SetTextFreetype2PropertiesResponse, error) {
	params.RequestType = "SetTextFreetype2Properties"
	data := &SetTextFreetype2PropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
