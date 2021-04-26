package testfixtures

type GetSourcesTypesListResponse struct {
	Ids []struct {
		Caps struct {
			CanInteract bool `json:"CanInteract"`

			DoNotDuplicate bool `json:"DoNotDuplicate"`

			DoNotSelfMonitor bool `json:"DoNotSelfMonitor"`

			Extra []struct {
				Poopy bool `json:"Poopy"`

				Wow bool `json:"Wow"`
			} `json:"Extra"`

			HasAudio bool `json:"HasAudio"`

			HasVideo bool `json:"HasVideo"`

			IsAsync bool `json:"IsAsync"`

			IsComposite bool `json:"IsComposite"`
		} `json:"Caps"`

		DefaultSettings map[string]interface{} `json:"DefaultSettings"`

		DisplayName string `json:"DisplayName"`

		Type string `json:"Type"`

		TypeId string `json:"TypeId"`
	} `json:"Ids"`
}
