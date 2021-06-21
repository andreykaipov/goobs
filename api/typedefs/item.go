package typedefs

// Item represents an item among several requests in the scene items category.
type Item struct {
	Id int `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}
