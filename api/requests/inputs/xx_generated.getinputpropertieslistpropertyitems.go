// This file has been automatically generated. Don't edit it.

package inputs

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
GetInputPropertiesListPropertyItemsParams represents the params body for the "GetInputPropertiesListPropertyItems" request.
Gets the items of a list property from an input's properties.

Note: Use this in cases where an input provides a dynamic, selectable list of items. For example, display capture, where it provides a list of available displays.
*/
type GetInputPropertiesListPropertyItemsParams struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Name of the list property to get the items of
	PropertyName string `json:"propertyName,omitempty"`
}

/*
GetInputPropertiesListPropertyItemsResponse represents the response body for the "GetInputPropertiesListPropertyItems" request.
Gets the items of a list property from an input's properties.

Note: Use this in cases where an input provides a dynamic, selectable list of items. For example, display capture, where it provides a list of available displays.
*/
type GetInputPropertiesListPropertyItemsResponse struct {
	PropertyItems []*typedefs.PropertyItem `json:"propertyItems,omitempty"`
}

// GetInputPropertiesListPropertyItems sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputPropertiesListPropertyItems(
	params *GetInputPropertiesListPropertyItemsParams,
) (*GetInputPropertiesListPropertyItemsResponse, error) {
	resp, err := c.SendRequest("GetInputPropertiesListPropertyItems", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputPropertiesListPropertyItemsResponse), nil
}
