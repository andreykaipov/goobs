// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputDeinterlaceMode request.
type GetInputDeinterlaceModeParams struct {
	// Name of the input
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewGetInputDeinterlaceModeParams() *GetInputDeinterlaceModeParams {
	return &GetInputDeinterlaceModeParams{}
}
func (o *GetInputDeinterlaceModeParams) WithInputName(x string) *GetInputDeinterlaceModeParams {
	o.InputName = &x
	return o
}
func (o *GetInputDeinterlaceModeParams) WithInputUuid(x string) *GetInputDeinterlaceModeParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *GetInputDeinterlaceModeParams) GetRequestName() string {
	return "GetInputDeinterlaceMode"
}

// Represents the response body for the GetInputDeinterlaceMode request.
type GetInputDeinterlaceModeResponse struct {
	_response

	// Deinterlace mode of the input
	InputDeinterlaceMode string `json:"inputDeinterlaceMode,omitempty"`
}

/*
Gets the deinterlace mode of an input.

Deinterlace Modes:

- `OBS_DEINTERLACE_MODE_DISABLE`
- `OBS_DEINTERLACE_MODE_DISCARD`
- `OBS_DEINTERLACE_MODE_RETRO`
- `OBS_DEINTERLACE_MODE_BLEND`
- `OBS_DEINTERLACE_MODE_BLEND_2X`
- `OBS_DEINTERLACE_MODE_LINEAR`
- `OBS_DEINTERLACE_MODE_LINEAR_2X`
- `OBS_DEINTERLACE_MODE_YADIF`
- `OBS_DEINTERLACE_MODE_YADIF_2X`

Note: Deinterlacing functionality is restricted to async inputs only.
*/
func (c *Client) GetInputDeinterlaceMode(
	paramss ...*GetInputDeinterlaceModeParams,
) (*GetInputDeinterlaceModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputDeinterlaceModeParams{{}}
	}
	params := paramss[0]
	data := &GetInputDeinterlaceModeResponse{}
	return data, c.client.SendRequest(params, data)
}
