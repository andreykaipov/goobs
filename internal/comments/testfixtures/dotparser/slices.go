package testfixtures

type GetSourcesListRequest struct {
	Sources []struct {
		Name string `json:"Name"`

		Type string `json:"Type"`

		TypeId string `json:"TypeId"`
	} `json:"Sources"`
}
