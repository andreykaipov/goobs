package testfixtures

type GetSourcesListRequest struct {
	Sources []*Source `json:"Sources,omitempty"`
}

type Source struct {
	Name string `json:"Name"`

	Type string `json:"Type"`

	TypeId string `json:"TypeId"`
}
