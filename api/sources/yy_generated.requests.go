// This file has been automatically generated. Don't edit it.

package sources

type GetSourcesListParams struct{}
type GetSourcesListResponse struct {
	Sources []struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		TypeId string `json:"typeId"`
	} `json:"sources"`
}
type GetSourcesTypesListParams struct{}
type GetSourcesTypesListResponse struct {
	Ids []struct {
		Caps struct {
			CanInteract      bool `json:"canInteract"`
			DoNotDuplicate   bool `json:"doNotDuplicate"`
			DoNotSelfMonitor bool `json:"doNotSelfMonitor"`
			HasAudio         bool `json:"hasAudio"`
			HasVideo         bool `json:"hasVideo"`
			IsAsync          bool `json:"isAsync"`
			IsComposite      bool `json:"isComposite"`
		} `json:"caps"`
		DefaultSettings map[string]interface{} `json:"defaultSettings"`
		DisplayName     string                 `json:"displayName"`
		Type            string                 `json:"type"`
		TypeId          string                 `json:"typeId"`
	} `json:"ids"`
}
type GetVolumeParams struct {
	Source string `json:"source"`
}
type GetVolumeResponse struct {
	Muted  bool    `json:"muted"`
	Name   string  `json:"name"`
	Volume float64 `json:"volume"`
}
type SetVolumeParams struct {
	Source string  `json:"source"`
	Volume float64 `json:"volume"`
}
type SetVolumeResponse struct{}
type GetMuteParams struct {
	Source string `json:"source"`
}
type GetMuteResponse struct {
	Muted bool   `json:"muted"`
	Name  string `json:"name"`
}
type SetMuteParams struct {
	Mute   bool   `json:"mute"`
	Source string `json:"source"`
}
type SetMuteResponse struct{}
type ToggleMuteParams struct {
	Source string `json:"source"`
}
type ToggleMuteResponse struct{}
type SetSyncOffsetParams struct {
	Offset int    `json:"offset"`
	Source string `json:"source"`
}
type SetSyncOffsetResponse struct{}
type GetSyncOffsetParams struct {
	Source string `json:"source"`
}
type GetSyncOffsetResponse struct {
	Name   string `json:"name"`
	Offset int    `json:"offset"`
}
type GetSourceSettingsParams struct {
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
}
type GetSourceSettingsResponse struct {
	SourceName     string                 `json:"sourceName"`
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	SourceType     string                 `json:"sourceType"`
}
type SetSourceSettingsParams struct {
	SourceName     string                 `json:"sourceName"`
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	SourceType     string                 `json:"sourceType"`
}
type SetSourceSettingsResponse struct {
	SourceName     string                 `json:"sourceName"`
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	SourceType     string                 `json:"sourceType"`
}
type GetTextGDIPlusPropertiesParams struct {
	Source string `json:"source"`
}
type GetTextGDIPlusPropertiesResponse struct {
	Align        string `json:"align"`
	BkColor      int    `json:"bk-color"`
	BkOpacity    int    `json:"bk-opacity"`
	Chatlog      bool   `json:"chatlog"`
	ChatlogLines int    `json:"chatlog_lines"`
	Color        int    `json:"color"`
	Extents      bool   `json:"extents"`
	ExtentsCx    int    `json:"extents_cx"`
	ExtentsCy    int    `json:"extents_cy"`
	File         string `json:"file"`
	Font         struct {
		Face  string `json:"face"`
		Flags int    `json:"flags"`
		Size  int    `json:"size"`
		Style string `json:"style"`
	} `json:"font"`
	Gradient        bool    `json:"gradient"`
	GradientColor   int     `json:"gradient_color"`
	GradientDir     float64 `json:"gradient_dir"`
	GradientOpacity int     `json:"gradient_opacity"`
	Outline         bool    `json:"outline"`
	OutlineColor    int     `json:"outline_color"`
	OutlineOpacity  int     `json:"outline_opacity"`
	OutlineSize     int     `json:"outline_size"`
	ReadFromFile    bool    `json:"read_from_file"`
	Source          string  `json:"source"`
	Text            string  `json:"text"`
	Valign          string  `json:"valign"`
	Vertical        bool    `json:"vertical"`
}
type SetTextGDIPlusPropertiesParams struct {
	Align        string `json:"align"`
	BkColor      int    `json:"bk-color"`
	BkOpacity    int    `json:"bk-opacity"`
	Chatlog      bool   `json:"chatlog"`
	ChatlogLines int    `json:"chatlog_lines"`
	Color        int    `json:"color"`
	Extents      bool   `json:"extents"`
	ExtentsCx    int    `json:"extents_cx"`
	ExtentsCy    int    `json:"extents_cy"`
	File         string `json:"file"`
	Font         struct {
		Face  string `json:"face"`
		Flags int    `json:"flags"`
		Size  int    `json:"size"`
		Style string `json:"style"`
	} `json:"font"`
	Gradient        bool    `json:"gradient"`
	GradientColor   int     `json:"gradient_color"`
	GradientDir     float64 `json:"gradient_dir"`
	GradientOpacity int     `json:"gradient_opacity"`
	Outline         bool    `json:"outline"`
	OutlineColor    int     `json:"outline_color"`
	OutlineOpacity  int     `json:"outline_opacity"`
	OutlineSize     int     `json:"outline_size"`
	ReadFromFile    bool    `json:"read_from_file"`
	Render          bool    `json:"render"`
	Source          string  `json:"source"`
	Text            string  `json:"text"`
	Valign          string  `json:"valign"`
	Vertical        bool    `json:"vertical"`
}
type SetTextGDIPlusPropertiesResponse struct{}
type GetTextFreetype2PropertiesParams struct {
	Source string `json:"source"`
}
type GetTextFreetype2PropertiesResponse struct {
	Color1      int  `json:"color1"`
	Color2      int  `json:"color2"`
	CustomWidth int  `json:"custom_width"`
	DropShadow  bool `json:"drop_shadow"`
	Font        struct {
		Face  string `json:"face"`
		Flags int    `json:"flags"`
		Size  int    `json:"size"`
		Style string `json:"style"`
	} `json:"font"`
	FromFile bool   `json:"from_file"`
	LogMode  bool   `json:"log_mode"`
	Outline  bool   `json:"outline"`
	Source   string `json:"source"`
	Text     string `json:"text"`
	TextFile string `json:"text_file"`
	WordWrap bool   `json:"word_wrap"`
}
type SetTextFreetype2PropertiesParams struct {
	Color1      int  `json:"color1"`
	Color2      int  `json:"color2"`
	CustomWidth int  `json:"custom_width"`
	DropShadow  bool `json:"drop_shadow"`
	Font        struct {
		Face  string `json:"face"`
		Flags int    `json:"flags"`
		Size  int    `json:"size"`
		Style string `json:"style"`
	} `json:"font"`
	FromFile bool   `json:"from_file"`
	LogMode  bool   `json:"log_mode"`
	Outline  bool   `json:"outline"`
	Source   string `json:"source"`
	Text     string `json:"text"`
	TextFile string `json:"text_file"`
	WordWrap bool   `json:"word_wrap"`
}
type SetTextFreetype2PropertiesResponse struct{}
type GetBrowserSourcePropertiesParams struct {
	Source string `json:"source"`
}
type GetBrowserSourcePropertiesResponse struct {
	Css         string `json:"css"`
	Fps         int    `json:"fps"`
	Height      int    `json:"height"`
	IsLocalFile bool   `json:"is_local_file"`
	LocalFile   string `json:"local_file"`
	Shutdown    bool   `json:"shutdown"`
	Source      string `json:"source"`
	Url         string `json:"url"`
	Width       int    `json:"width"`
}
type SetBrowserSourcePropertiesParams struct {
	Css         string `json:"css"`
	Fps         int    `json:"fps"`
	Height      int    `json:"height"`
	IsLocalFile bool   `json:"is_local_file"`
	LocalFile   string `json:"local_file"`
	Render      bool   `json:"render"`
	Shutdown    bool   `json:"shutdown"`
	Source      string `json:"source"`
	Url         string `json:"url"`
	Width       int    `json:"width"`
}
type SetBrowserSourcePropertiesResponse struct{}
type GetSpecialSourcesParams struct{}
type GetSpecialSourcesResponse struct {
	Desktop1 string `json:"desktop-1"`
	Desktop2 string `json:"desktop-2"`
	Mic1     string `json:"mic-1"`
	Mic2     string `json:"mic-2"`
	Mic3     string `json:"mic-3"`
}
type GetSourceFiltersParams struct {
	SourceName string `json:"sourceName"`
}
type GetSourceFiltersResponse struct {
	Filters []struct {
		Name     string                 `json:"name"`
		Settings map[string]interface{} `json:"settings"`
		Type     string                 `json:"type"`
	} `json:"filters"`
}
type AddFilterToSourceParams struct {
	FilterName     string                 `json:"filterName"`
	FilterSettings map[string]interface{} `json:"filterSettings"`
	FilterType     string                 `json:"filterType"`
	SourceName     string                 `json:"sourceName"`
}
type AddFilterToSourceResponse struct{}
type RemoveFilterFromSourceParams struct {
	FilterName string `json:"filterName"`
	SourceName string `json:"sourceName"`
}
type RemoveFilterFromSourceResponse struct{}
type ReorderSourceFilterParams struct {
	FilterName string `json:"filterName"`
	NewIndex   int    `json:"newIndex"`
	SourceName string `json:"sourceName"`
}
type ReorderSourceFilterResponse struct{}
type MoveSourceFilterParams struct {
	FilterName   string `json:"filterName"`
	MovementType string `json:"movementType"`
	SourceName   string `json:"sourceName"`
}
type MoveSourceFilterResponse struct{}
type SetSourceFilterSettingsParams struct {
	FilterName     string                 `json:"filterName"`
	FilterSettings map[string]interface{} `json:"filterSettings"`
	SourceName     string                 `json:"sourceName"`
}
type SetSourceFilterSettingsResponse struct{}
