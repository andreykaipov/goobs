// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
PressInputPropertiesButtonParams represents the params body for the "PressInputPropertiesButton" request.
Presses a button in the properties of an input.

Note: Use this in cases where there is a button in the properties of an input that cannot be accessed in any other way. For example, browser sources, where there is a refresh button.
*/
type PressInputPropertiesButtonParams struct {
	requests.ParamsBasic

	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Name of the button property to press
	PropertyName string `json:"propertyName,omitempty"`
}

// GetSelfName just returns "PressInputPropertiesButton".
func (o *PressInputPropertiesButtonParams) GetSelfName() string {
	return "PressInputPropertiesButton"
}

/*
PressInputPropertiesButtonResponse represents the response body for the "PressInputPropertiesButton" request.
Presses a button in the properties of an input.

Note: Use this in cases where there is a button in the properties of an input that cannot be accessed in any other way. For example, browser sources, where there is a refresh button.
*/
type PressInputPropertiesButtonResponse struct {
	requests.ResponseBasic
}

// PressInputPropertiesButton sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) PressInputPropertiesButton(
	params *PressInputPropertiesButtonParams,
) (*PressInputPropertiesButtonResponse, error) {
	data := &PressInputPropertiesButtonResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
