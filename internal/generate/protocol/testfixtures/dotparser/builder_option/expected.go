package testfixtures

type GetSourcesTypesListResponse struct {
	Ids []*Id `json:"Ids,omitempty"`

	Items []*Item `json:"Items,omitempty"`
}

type Id struct {
	DisplayName string `json:"DisplayName"`

	Type string `json:"Type"`

	TypeId string `json:"TypeId"`
}

type Item struct {
	Id int `json:"Id"`

	Name string `json:"Name"`
}
