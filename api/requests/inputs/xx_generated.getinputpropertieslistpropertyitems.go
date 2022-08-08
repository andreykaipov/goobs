// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetInputPropertiesListPropertyItems request.
type GetInputPropertiesListPropertyItemsParams struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Name of the list property to get the items of
	PropertyName string `json:"propertyName,omitempty"`
}

// Returns the associated request.
func (o *GetInputPropertiesListPropertyItemsParams) GetRequestName() string {
	return "GetInputPropertiesListPropertyItems"
}

// Represents the response body for the GetInputPropertiesListPropertyItems request.
type GetInputPropertiesListPropertyItemsResponse struct {
	PropertyItems []*typedefs.PropertyItem `json:"propertyItems,omitempty"`
}

/*
Gets the items of a list property from an input's properties.

Note: Use this in cases where an input provides a dynamic, selectable list of items. For example, display capture, where it provides a list of available displays.
*/
func (c *Client) GetInputPropertiesListPropertyItems(
	params *GetInputPropertiesListPropertyItemsParams,
) (*GetInputPropertiesListPropertyItemsResponse, error) {
	data := &GetInputPropertiesListPropertyItemsResponse{}
	return data, c.SendRequest(params, data)
}
