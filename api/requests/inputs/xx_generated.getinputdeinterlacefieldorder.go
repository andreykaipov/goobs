// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputDeinterlaceFieldOrder request.
type GetInputDeinterlaceFieldOrderParams struct {
	// Name of the input
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewGetInputDeinterlaceFieldOrderParams() *GetInputDeinterlaceFieldOrderParams {
	return &GetInputDeinterlaceFieldOrderParams{}
}
func (o *GetInputDeinterlaceFieldOrderParams) WithInputName(x string) *GetInputDeinterlaceFieldOrderParams {
	o.InputName = &x
	return o
}
func (o *GetInputDeinterlaceFieldOrderParams) WithInputUuid(x string) *GetInputDeinterlaceFieldOrderParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *GetInputDeinterlaceFieldOrderParams) GetRequestName() string {
	return "GetInputDeinterlaceFieldOrder"
}

// Represents the response body for the GetInputDeinterlaceFieldOrder request.
type GetInputDeinterlaceFieldOrderResponse struct {
	_response

	// Deinterlace field order of the input
	InputDeinterlaceFieldOrder string `json:"inputDeinterlaceFieldOrder,omitempty"`
}

/*
Gets the deinterlace field order of an input.

Deinterlace Field Orders:

- `OBS_DEINTERLACE_FIELD_ORDER_TOP`
- `OBS_DEINTERLACE_FIELD_ORDER_BOTTOM`

Note: Deinterlacing functionality is restricted to async inputs only.
*/
func (c *Client) GetInputDeinterlaceFieldOrder(
	paramss ...*GetInputDeinterlaceFieldOrderParams,
) (*GetInputDeinterlaceFieldOrderResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputDeinterlaceFieldOrderParams{{}}
	}
	params := paramss[0]
	data := &GetInputDeinterlaceFieldOrderResponse{}
	return data, c.client.SendRequest(params, data)
}
