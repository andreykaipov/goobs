package typedefs

type Input struct {
	InputName            string `json:"inputName"`
	InputKind            string `json:"inputKind"`
	UnversionedInputKind string `json:"unversionedInputKind"`
}

type Output struct {
	Name   string       `json:"outputName"`
	Kind   string       `json:"outputKind"`
	Width  int          `json:"outputWidth"`
	Height int          `json:"outputHeight"`
	Active bool         `json:"outputActive"`
	Flags  *OutputFlags `json:"outputFlags"`
}

type OutputFlags struct {
	Audio      bool `json:"OBS_OUTPUT_AUDIO"`
	Video      bool `json:"OBS_OUTPUT_VIDEO"`
	Encoded    bool `json:"OBS_OUTPUT_ENCODED"`
	MultiTrack bool `json:"OBS_OUTPUT_MULTI_TRACK"`
	Service    bool `json:"OBS_OUTPUT_SERVICE"`
}

type Scene struct {
	SceneIndex int    `json:"sceneIndex"`
	SceneName  string `json:"sceneName"`
}

type PropertyItem struct {
	ItemName    string `json:"itemName"`
	ItemEnabled bool   `json:"itemEnabled"`
	ItemValue   any    `json:"itemValue"`
}

type Filter struct {
	FilterEnabled  bool           `json:"filterEnabled"`
	FilterIndex    int            `json:"filterIndex"`
	FilterKind     string         `json:"filterKind"`
	FilterName     string         `json:"filterName"`
	FilterSettings map[string]any `json:"filterSettings,omitempty"`
}

type Transition struct {
	TransitionConfigurable bool   `json:"transitionConfigurable"`
	TransitionFixed        bool   `json:"transitionFixed"`
	TransitionKind         string `json:"transitionKind"`
	TransitionName         string `json:"transitionName"`
}

type SceneItemBasic struct {
	SceneItemID    int `json:"sceneItemId"`
	SceneItemIndex int `json:"sceneItemIndex"`
}

type SceneItem struct {
	InputKind          string             `json:"inputKind"`
	IsGroup            bool               `json:"isGroup"`
	SceneItemBlendMode string             `json:"sceneItemBlendMode"`
	SceneItemEnabled   bool               `json:"sceneItemEnabled"`
	SceneItemID        int                `json:"sceneItemId"`
	SceneItemIndex     int                `json:"sceneItemIndex"`
	SceneItemLocked    bool               `json:"sceneItemLocked"`
	SceneItemTransform SceneItemTransform `json:"sceneItemTransform"`
	SourceName         string             `json:"sourceName"`
	SourceType         string             `json:"sourceType"`
}

type InputAudioTracks map[string]bool

type KeyModifiers struct {
	Shift   string `json:"face"`
	Control int    `json:"flags"`
	Alt     int    `json:"size"`
	Command string `json:"style"`
}

type Monitor struct {
	MonitorHeight    int    `json:"monitorHeight"`
	MonitorIndex     int    `json:"monitorIndex"`
	MonitorName      string `json:"monitorName"`
	MonitorPositionX int    `json:"monitorPositionX"`
	MonitorPositionY int    `json:"monitorPositionY"`
	MonitorWidth     int    `json:"monitorWidth"`
}

type StreamServiceSettings struct {
	Bwtest   bool   `json:"bwtest"`
	Key      string `json:"key"`
	Password string `json:"password"`
	Server   string `json:"server"`
	UseAuth  bool   `json:"use_auth"`
	Username string `json:"username"`
}

type SceneItemTransform struct {
	Alignment       float64 `json:"alignment"`
	BoundsAlignment float64 `json:"boundsAlignment"`
	BoundsHeight    float64 `json:"boundsHeight"`
	BoundsType      string  `json:"boundsType"`
	BoundsWidth     float64 `json:"boundsWidth"`
	CropBottom      float64 `json:"cropBottom"`
	CropLeft        float64 `json:"cropLeft"`
	CropRight       float64 `json:"cropRight"`
	CropTop         float64 `json:"cropTop"`
	Height          float64 `json:"height"`
	PositionX       float64 `json:"positionX"`
	PositionY       float64 `json:"positionY"`
	Rotation        float64 `json:"rotation"`
	ScaleX          float64 `json:"scaleX"`
	ScaleY          float64 `json:"scaleY"`
	SourceHeight    float64 `json:"sourceHeight"`
	SourceWidth     float64 `json:"sourceWidth"`
	Width           float64 `json:"width"`
}

type InputVolumeMeter struct {
	Name   string       `json:"inputName"`
	Levels [][3]float64 `json:"inputLevelsMul"`
}
