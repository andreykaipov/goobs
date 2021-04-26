// This file has been automatically generated. Don't edit it.

package sources

type GetSourcesListParams struct{}
type GetSourcesListResponse struct {
	Sources []struct {
		// Unique source name
		Name string `json:"name"`
		// Source type. Value is one of the following: "input", "filter", "transition", "scene" or "unknown"
		Type string `json:"type"`
		// Non-unique source internal type (a.k.a type id)
		TypeId string `json:"typeId"`
	} `json:"sources"`
}
type GetSourcesTypesListParams struct{}
type GetSourcesTypesListResponse struct {
	Ids []struct {
		Caps struct {
			// True if interaction with this sources of this type is possible
			CanInteract bool `json:"canInteract"`
			// True if sources of this type should not be fully duplicated
			DoNotDuplicate bool `json:"doNotDuplicate"`
			// True if sources of this type may cause a feedback loop if it's audio is monitored and shouldn't be
			DoNotSelfMonitor bool `json:"doNotSelfMonitor"`
			// True if sources of this type provide audio
			HasAudio bool `json:"hasAudio"`
			// True if sources of this type provide video
			HasVideo bool `json:"hasVideo"`
			// True if source of this type provide frames asynchronously
			IsAsync bool `json:"isAsync"`
			// True if sources of this type composite one or more sub-sources
			IsComposite bool `json:"isComposite"`
		} `json:"caps"`
		// Default settings of this source type
		DefaultSettings map[string]interface{} `json:"defaultSettings"`
		// Display name of the source type
		DisplayName string `json:"displayName"`
		// Type. Value is one of the following: "input", "filter", "transition" or "other"
		Type string `json:"type"`
		// Non-unique internal source type ID
		TypeId string `json:"typeId"`
	} `json:"ids"`
}
type GetVolumeParams struct {
	// Source name.
	Source string `json:"source"`
}
type GetVolumeResponse struct {
	// Indicates whether the source is muted.
	Muted bool `json:"muted"`
	// Source name.
	Name string `json:"name"`
	// Volume of the source. Between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}
type SetVolumeParams struct {
	// Source name.
	Source string `json:"source"`
	// Desired volume. Must be between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}
type SetVolumeResponse struct{}
type GetMuteParams struct {
	// Source name.
	Source string `json:"source"`
}
type GetMuteResponse struct {
	// Mute status of the source.
	Muted bool `json:"muted"`
	// Source name.
	Name string `json:"name"`
}
type SetMuteParams struct {
	// Desired mute status.
	Mute bool `json:"mute"`
	// Source name.
	Source string `json:"source"`
}
type SetMuteResponse struct{}
type ToggleMuteParams struct {
	// Source name.
	Source string `json:"source"`
}
type ToggleMuteResponse struct{}
type SetSyncOffsetParams struct {
	// The desired audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
	// Source name.
	Source string `json:"source"`
}
type SetSyncOffsetResponse struct{}
type GetSyncOffsetParams struct {
	// Source name.
	Source string `json:"source"`
}
type GetSyncOffsetResponse struct {
	// Source name.
	Name string `json:"name"`
	// The audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
}
type GetSourceSettingsParams struct {
	// Source name.
	SourceName string `json:"sourceName"`
	// Type of the specified source. Useful for type-checking if you expect a specific settings schema.
	SourceType string `json:"sourceType"`
}
type GetSourceSettingsResponse struct {
	// Source name
	SourceName string `json:"sourceName"`
	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	// Type of the specified source
	SourceType string `json:"sourceType"`
}
type SetSourceSettingsParams struct {
	// Source name.
	SourceName string `json:"sourceName"`
	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	// Type of the specified source. Useful for type-checking to avoid settings a set of settings incompatible with the actual source's type.
	SourceType string `json:"sourceType"`
}
type SetSourceSettingsResponse struct {
	// Source name
	SourceName string `json:"sourceName"`
	// Updated source settings
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	// Type of the specified source
	SourceType string `json:"sourceType"`
}
type GetTextGDIPlusPropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}
type GetTextGDIPlusPropertiesResponse struct {
	// Text Alignment ("left", "center", "right").
	Align string `json:"align"`
	// Background color.
	BkColor int `json:"bk-color"`
	// Background opacity (0-100).
	BkOpacity int `json:"bk-opacity"`
	// Chat log.
	Chatlog bool `json:"chatlog"`
	// Chat log lines.
	ChatlogLines int `json:"chatlog_lines"`
	// Text color.
	Color int `json:"color"`
	// Extents wrap.
	Extents bool `json:"extents"`
	// Extents cx.
	ExtentsCx int `json:"extents_cx"`
	// Extents cy.
	ExtentsCy int `json:"extents_cy"`
	// File path name.
	File string `json:"file"`
	Font struct {
		// Font face.
		Face string `json:"face"`
		// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3, Underline=5, Strikeout=8`
		Flags int `json:"flags"`
		// Font text size.
		Size int `json:"size"`
		// Font Style (unknown function).
		Style string `json:"style"`
	} `json:"font"`
	// Gradient enabled.
	Gradient bool `json:"gradient"`
	// Gradient color.
	GradientColor int `json:"gradient_color"`
	// Gradient direction.
	GradientDir float64 `json:"gradient_dir"`
	// Gradient opacity (0-100).
	GradientOpacity int `json:"gradient_opacity"`
	// Outline.
	Outline bool `json:"outline"`
	// Outline color.
	OutlineColor int `json:"outline_color"`
	// Outline opacity (0-100).
	OutlineOpacity int `json:"outline_opacity"`
	// Outline size.
	OutlineSize int `json:"outline_size"`
	// Read text from the specified file.
	ReadFromFile bool `json:"read_from_file"`
	// Source name.
	Source string `json:"source"`
	// Text content to be displayed.
	Text string `json:"text"`
	// Text vertical alignment ("top", "center", "bottom").
	Valign string `json:"valign"`
	// Vertical text enabled.
	Vertical bool `json:"vertical"`
}
type SetTextGDIPlusPropertiesParams struct {
	// Text Alignment ("left", "center", "right").
	Align string `json:"align"`
	// Background color.
	BkColor int `json:"bk-color"`
	// Background opacity (0-100).
	BkOpacity int `json:"bk-opacity"`
	// Chat log.
	Chatlog bool `json:"chatlog"`
	// Chat log lines.
	ChatlogLines int `json:"chatlog_lines"`
	// Text color.
	Color int `json:"color"`
	// Extents wrap.
	Extents bool `json:"extents"`
	// Extents cx.
	ExtentsCx int `json:"extents_cx"`
	// Extents cy.
	ExtentsCy int `json:"extents_cy"`
	// File path name.
	File string `json:"file"`
	Font struct {
		// Font face.
		Face string `json:"face"`
		// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3, Underline=5, Strikeout=8`
		Flags int `json:"flags"`
		// Font text size.
		Size int `json:"size"`
		// Font Style (unknown function).
		Style string `json:"style"`
	} `json:"font"`
	// Gradient enabled.
	Gradient bool `json:"gradient"`
	// Gradient color.
	GradientColor int `json:"gradient_color"`
	// Gradient direction.
	GradientDir float64 `json:"gradient_dir"`
	// Gradient opacity (0-100).
	GradientOpacity int `json:"gradient_opacity"`
	// Outline.
	Outline bool `json:"outline"`
	// Outline color.
	OutlineColor int `json:"outline_color"`
	// Outline opacity (0-100).
	OutlineOpacity int `json:"outline_opacity"`
	// Outline size.
	OutlineSize int `json:"outline_size"`
	// Read text from the specified file.
	ReadFromFile bool `json:"read_from_file"`
	// Visibility of the scene item.
	Render bool `json:"render"`
	// Name of the source.
	Source string `json:"source"`
	// Text content to be displayed.
	Text string `json:"text"`
	// Text vertical alignment ("top", "center", "bottom").
	Valign string `json:"valign"`
	// Vertical text enabled.
	Vertical bool `json:"vertical"`
}
type SetTextGDIPlusPropertiesResponse struct{}
type GetTextFreetype2PropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}
type GetTextFreetype2PropertiesResponse struct {
	// Gradient top color.
	Color1 int `json:"color1"`
	// Gradient bottom color.
	Color2 int `json:"color2"`
	// Custom width (0 to disable).
	CustomWidth int `json:"custom_width"`
	// Drop shadow.
	DropShadow bool `json:"drop_shadow"`
	Font       struct {
		// Font face.
		Face string `json:"face"`
		// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3, Underline=5, Strikeout=8`
		Flags int `json:"flags"`
		// Font text size.
		Size int `json:"size"`
		// Font Style (unknown function).
		Style string `json:"style"`
	} `json:"font"`
	// Read text from the specified file.
	FromFile bool `json:"from_file"`
	// Chat log.
	LogMode bool `json:"log_mode"`
	// Outline.
	Outline bool `json:"outline"`
	// Source name
	Source string `json:"source"`
	// Text content to be displayed.
	Text string `json:"text"`
	// File path.
	TextFile string `json:"text_file"`
	// Word wrap.
	WordWrap bool `json:"word_wrap"`
}
type SetTextFreetype2PropertiesParams struct {
	// Gradient top color.
	Color1 int `json:"color1"`
	// Gradient bottom color.
	Color2 int `json:"color2"`
	// Custom width (0 to disable).
	CustomWidth int `json:"custom_width"`
	// Drop shadow.
	DropShadow bool `json:"drop_shadow"`
	Font       struct {
		// Font face.
		Face string `json:"face"`
		// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3, Underline=5, Strikeout=8`
		Flags int `json:"flags"`
		// Font text size.
		Size int `json:"size"`
		// Font Style (unknown function).
		Style string `json:"style"`
	} `json:"font"`
	// Read text from the specified file.
	FromFile bool `json:"from_file"`
	// Chat log.
	LogMode bool `json:"log_mode"`
	// Outline.
	Outline bool `json:"outline"`
	// Source name.
	Source string `json:"source"`
	// Text content to be displayed.
	Text string `json:"text"`
	// File path.
	TextFile string `json:"text_file"`
	// Word wrap.
	WordWrap bool `json:"word_wrap"`
}
type SetTextFreetype2PropertiesResponse struct{}
type GetBrowserSourcePropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}
type GetBrowserSourcePropertiesResponse struct {
	// CSS to inject.
	Css string `json:"css"`
	// Framerate.
	Fps int `json:"fps"`
	// Height.
	Height int `json:"height"`
	// Indicates that a local file is in use.
	IsLocalFile bool `json:"is_local_file"`
	// file path.
	LocalFile string `json:"local_file"`
	// Indicates whether the source should be shutdown when not visible.
	Shutdown bool `json:"shutdown"`
	// Source name.
	Source string `json:"source"`
	// Url.
	Url string `json:"url"`
	// Width.
	Width int `json:"width"`
}
type SetBrowserSourcePropertiesParams struct {
	// CSS to inject.
	Css string `json:"css"`
	// Framerate.
	Fps int `json:"fps"`
	// Height.
	Height int `json:"height"`
	// Indicates that a local file is in use.
	IsLocalFile bool `json:"is_local_file"`
	// file path.
	LocalFile string `json:"local_file"`
	// Visibility of the scene item.
	Render bool `json:"render"`
	// Indicates whether the source should be shutdown when not visible.
	Shutdown bool `json:"shutdown"`
	// Name of the source.
	Source string `json:"source"`
	// Url.
	Url string `json:"url"`
	// Width.
	Width int `json:"width"`
}
type SetBrowserSourcePropertiesResponse struct{}
type GetSpecialSourcesParams struct{}
type GetSpecialSourcesResponse struct {
	// Name of the first Desktop Audio capture source.
	Desktop1 string `json:"desktop-1"`
	// Name of the second Desktop Audio capture source.
	Desktop2 string `json:"desktop-2"`
	// Name of the first Mic/Aux input source.
	Mic1 string `json:"mic-1"`
	// Name of the second Mic/Aux input source.
	Mic2 string `json:"mic-2"`
	// NAme of the third Mic/Aux input source.
	Mic3 string `json:"mic-3"`
}
type GetSourceFiltersParams struct {
	// Source name
	SourceName string `json:"sourceName"`
}
type GetSourceFiltersResponse struct {
	Filters []struct {
		// Filter name
		Name string `json:"name"`
		// Filter settings
		Settings map[string]interface{} `json:"settings"`
		// Filter type
		Type string `json:"type"`
	} `json:"filters"`
}
type AddFilterToSourceParams struct {
	// Name of the new filter
	FilterName string `json:"filterName"`
	// Filter settings
	FilterSettings map[string]interface{} `json:"filterSettings"`
	// Filter type
	FilterType string `json:"filterType"`
	// Name of the source on which the filter is added
	SourceName string `json:"sourceName"`
}
type AddFilterToSourceResponse struct{}
type RemoveFilterFromSourceParams struct {
	// Name of the filter to remove
	FilterName string `json:"filterName"`
	// Name of the source from which the specified filter is removed
	SourceName string `json:"sourceName"`
}
type RemoveFilterFromSourceResponse struct{}
type ReorderSourceFilterParams struct {
	// Name of the filter to reorder
	FilterName string `json:"filterName"`
	// Desired position of the filter in the chain
	NewIndex int `json:"newIndex"`
	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}
type ReorderSourceFilterResponse struct{}
type MoveSourceFilterParams struct {
	// Name of the filter to reorder
	FilterName string `json:"filterName"`
	// How to move the filter around in the source's filter chain. Either "up", "down", "top" or "bottom".
	MovementType string `json:"movementType"`
	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}
type MoveSourceFilterResponse struct{}
type SetSourceFilterSettingsParams struct {
	// Name of the filter to reconfigure
	FilterName string `json:"filterName"`
	// New settings. These will be merged to the current filter settings.
	FilterSettings map[string]interface{} `json:"filterSettings"`
	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}
type SetSourceFilterSettingsResponse struct{}
