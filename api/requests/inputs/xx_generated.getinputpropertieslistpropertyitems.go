// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputPropertiesListPropertyItemsParams represents the params body for the "GetInputPropertiesListPropertyItems" request.
Gets the items of a list property from an input's properties.

Note: Use this in cases where an input provides a dynamic, selectable list of items. For example, display capture, where it provides a list of available displays.
*/
type GetInputPropertiesListPropertyItemsParams struct {
	requests.ParamsBasic

	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Name of the list property to get the items of
	PropertyName string `json:"propertyName,omitempty"`
}

// GetSelfName just returns "GetInputPropertiesListPropertyItems".
func (o *GetInputPropertiesListPropertyItemsParams) GetSelfName() string {
	return "GetInputPropertiesListPropertyItems"
}

/*
GetInputPropertiesListPropertyItemsResponse represents the response body for the "GetInputPropertiesListPropertyItems" request.
Gets the items of a list property from an input's properties.

Note: Use this in cases where an input provides a dynamic, selectable list of items. For example, display capture, where it provides a list of available displays.
*/
type GetInputPropertiesListPropertyItemsResponse struct {
	requests.ResponseBasic

	// Array of items in the list property
	PropertyItems []interface{} `json:"propertyItems,omitempty"`
}

// GetInputPropertiesListPropertyItems sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputPropertiesListPropertyItems(
	params *GetInputPropertiesListPropertyItemsParams,
) (*GetInputPropertiesListPropertyItemsResponse, error) {
	data := &GetInputPropertiesListPropertyItemsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
