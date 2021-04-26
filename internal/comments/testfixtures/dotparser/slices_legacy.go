package testfixtures

type ReorderSceneItemsRequestLegacy struct {
	Items []struct {
		Id int `json:"Id"`

		Name string `json:"Name"`
	} `json:"Items"`
}
