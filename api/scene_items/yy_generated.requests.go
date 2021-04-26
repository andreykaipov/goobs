// This file has been automatically generated. Don't edit it.

package sceneitems

type GetSceneItemPropertiesParams struct {
	Item      string `json:"item"`
	SceneName string `json:"scene-name"`
}
type GetSceneItemPropertiesResponse struct {
	Bounds struct {
		Alignment int     `json:"alignment"`
		Type      string  `json:"type"`
		X         float64 `json:"x"`
		Y         float64 `json:"y"`
	} `json:"bounds"`
	Crop struct {
		Bottom int `json:"bottom"`
		Left   int `json:"left"`
		Right  int `json:"right"`
		Top    int `json:"top"`
	} `json:"crop"`
	Name     string `json:"name"`
	Position struct {
		Alignment int `json:"alignment"`
		X         int `json:"x"`
		Y         int `json:"y"`
	} `json:"position"`
	Rotation float64 `json:"rotation"`
	Scale    struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"scale"`
	Visible bool `json:"visible"`
}
type SetSceneItemPropertiesParams struct {
	Bounds struct {
		Alignment int     `json:"alignment"`
		Type      string  `json:"type"`
		X         float64 `json:"x"`
		Y         float64 `json:"y"`
	} `json:"bounds"`
	Crop struct {
		Bottom int `json:"bottom"`
		Left   int `json:"left"`
		Right  int `json:"right"`
		Top    int `json:"top"`
	} `json:"crop"`
	Item     string `json:"item"`
	Position struct {
		Alignment int `json:"alignment"`
		X         int `json:"x"`
		Y         int `json:"y"`
	} `json:"position"`
	Rotation float64 `json:"rotation"`
	Scale    struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"scale"`
	SceneName string `json:"scene-name"`
	Visible   bool   `json:"visible"`
}
type SetSceneItemPropertiesResponse struct{}
type ResetSceneItemParams struct {
	Item      string `json:"item"`
	SceneName string `json:"scene-name"`
}
type ResetSceneItemResponse struct{}
type DeleteSceneItemParams struct {
	Item struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"item"`
	Scene string `json:"scene"`
}
type DeleteSceneItemResponse struct{}
type DuplicateSceneItemParams struct {
	FromScene string `json:"fromScene"`
	Item      struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"item"`
	ToScene string `json:"toScene"`
}
type DuplicateSceneItemResponse struct {
	Item struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"item"`
	Scene string `json:"scene"`
}
