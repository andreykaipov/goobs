package testfixtures

type ReorderSceneItemsRequestLegacy struct {
	Items []*Item `json:"Items,omitempty"`
}

type Item struct {
	Id int `json:"Id"`

	Name string `json:"Name"`
}
