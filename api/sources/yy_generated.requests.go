// This file has been automatically generated. Don't edit it.

package sources

// GetSourcesListParams contains the request body for the [GetSourcesList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList) request.
type GetSourcesListParams struct{}

// GetSourcesListResponse contains the request body for the [GetSourcesList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList) request.
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

// GetSourcesTypesListParams contains the request body for the [GetSourcesTypesList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList) request.
type GetSourcesTypesListParams struct{}

// GetSourcesTypesListResponse contains the request body for the [GetSourcesTypesList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList) request.
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

// GetVolumeParams contains the request body for the [GetVolume](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume) request.
type GetVolumeParams struct {
	// Source name.
	Source string `json:"source"`
}

// GetVolumeResponse contains the request body for the [GetVolume](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume) request.
type GetVolumeResponse struct {
	// Indicates whether the source is muted.
	Muted bool `json:"muted"`
	// Source name.
	Name string `json:"name"`
	// Volume of the source. Between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

// SetVolumeParams contains the request body for the [SetVolume](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume) request.
type SetVolumeParams struct {
	// Source name.
	Source string `json:"source"`
	// Desired volume. Must be between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

// SetVolumeResponse contains the request body for the [SetVolume](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume) request.
type SetVolumeResponse struct{}

// GetMuteParams contains the request body for the [GetMute](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute) request.
type GetMuteParams struct {
	// Source name.
	Source string `json:"source"`
}

// GetMuteResponse contains the request body for the [GetMute](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute) request.
type GetMuteResponse struct {
	// Mute status of the source.
	Muted bool `json:"muted"`
	// Source name.
	Name string `json:"name"`
}

// SetMuteParams contains the request body for the [SetMute](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute) request.
type SetMuteParams struct {
	// Desired mute status.
	Mute bool `json:"mute"`
	// Source name.
	Source string `json:"source"`
}

// SetMuteResponse contains the request body for the [SetMute](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute) request.
type SetMuteResponse struct{}

// ToggleMuteParams contains the request body for the [ToggleMute](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute) request.
type ToggleMuteParams struct {
	// Source name.
	Source string `json:"source"`
}

// ToggleMuteResponse contains the request body for the [ToggleMute](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute) request.
type ToggleMuteResponse struct{}

// SetSyncOffsetParams contains the request body for the [SetSyncOffset](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset) request.
type SetSyncOffsetParams struct {
	// The desired audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
	// Source name.
	Source string `json:"source"`
}

// SetSyncOffsetResponse contains the request body for the [SetSyncOffset](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset) request.
type SetSyncOffsetResponse struct{}

// GetSyncOffsetParams contains the request body for the [GetSyncOffset](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset) request.
type GetSyncOffsetParams struct {
	// Source name.
	Source string `json:"source"`
}

// GetSyncOffsetResponse contains the request body for the [GetSyncOffset](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset) request.
type GetSyncOffsetResponse struct {
	// Source name.
	Name string `json:"name"`
	// The audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
}

// GetSourceSettingsParams contains the request body for the [GetSourceSettings](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings) request.
type GetSourceSettingsParams struct {
	// Source name.
	SourceName string `json:"sourceName"`
	// Type of the specified source. Useful for type-checking if you expect a specific settings schema.
	SourceType string `json:"sourceType"`
}

// GetSourceSettingsResponse contains the request body for the [GetSourceSettings](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings) request.
type GetSourceSettingsResponse struct {
	// Source name
	SourceName string `json:"sourceName"`
	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	// Type of the specified source
	SourceType string `json:"sourceType"`
}

// SetSourceSettingsParams contains the request body for the [SetSourceSettings](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings) request.
type SetSourceSettingsParams struct {
	// Source name.
	SourceName string `json:"sourceName"`
	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	// Type of the specified source. Useful for type-checking to avoid settings a set of settings incompatible with the actual source's type.
	SourceType string `json:"sourceType"`
}

// SetSourceSettingsResponse contains the request body for the [SetSourceSettings](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings) request.
type SetSourceSettingsResponse struct {
	// Source name
	SourceName string `json:"sourceName"`
	// Updated source settings
	SourceSettings map[string]interface{} `json:"sourceSettings"`
	// Type of the specified source
	SourceType string `json:"sourceType"`
}

// GetTextGDIPlusPropertiesResponse contains the request body for the [GetTextGDIPlusProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextGDIPlusProperties) request.
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

// GetTextGDIPlusPropertiesParams contains the request body for the [GetTextGDIPlusProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextGDIPlusProperties) request.
type GetTextGDIPlusPropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}

// SetTextGDIPlusPropertiesParams contains the request body for the [SetTextGDIPlusProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties) request.
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

// SetTextGDIPlusPropertiesResponse contains the request body for the [SetTextGDIPlusProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties) request.
type SetTextGDIPlusPropertiesResponse struct{}

// GetTextFreetype2PropertiesParams contains the request body for the [GetTextFreetype2Properties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties) request.
type GetTextFreetype2PropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}

// GetTextFreetype2PropertiesResponse contains the request body for the [GetTextFreetype2Properties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties) request.
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

// SetTextFreetype2PropertiesParams contains the request body for the [SetTextFreetype2Properties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties) request.
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

// SetTextFreetype2PropertiesResponse contains the request body for the [SetTextFreetype2Properties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties) request.
type SetTextFreetype2PropertiesResponse struct{}

// GetBrowserSourcePropertiesParams contains the request body for the [GetBrowserSourceProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties) request.
type GetBrowserSourcePropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}

// GetBrowserSourcePropertiesResponse contains the request body for the [GetBrowserSourceProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties) request.
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

// SetBrowserSourcePropertiesResponse contains the request body for the [SetBrowserSourceProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties) request.
type SetBrowserSourcePropertiesResponse struct{}

// SetBrowserSourcePropertiesParams contains the request body for the [SetBrowserSourceProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties) request.
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

// GetSpecialSourcesParams contains the request body for the [GetSpecialSources](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources) request.
type GetSpecialSourcesParams struct{}

// GetSpecialSourcesResponse contains the request body for the [GetSpecialSources](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources) request.
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

// GetSourceFiltersParams contains the request body for the [GetSourceFilters](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters) request.
type GetSourceFiltersParams struct {
	// Source name
	SourceName string `json:"sourceName"`
}

// GetSourceFiltersResponse contains the request body for the [GetSourceFilters](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters) request.
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

// AddFilterToSourceParams contains the request body for the [AddFilterToSource](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource) request.
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

// AddFilterToSourceResponse contains the request body for the [AddFilterToSource](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource) request.
type AddFilterToSourceResponse struct{}

// RemoveFilterFromSourceParams contains the request body for the [RemoveFilterFromSource](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource) request.
type RemoveFilterFromSourceParams struct {
	// Name of the filter to remove
	FilterName string `json:"filterName"`
	// Name of the source from which the specified filter is removed
	SourceName string `json:"sourceName"`
}

// RemoveFilterFromSourceResponse contains the request body for the [RemoveFilterFromSource](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource) request.
type RemoveFilterFromSourceResponse struct{}

// ReorderSourceFilterParams contains the request body for the [ReorderSourceFilter](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter) request.
type ReorderSourceFilterParams struct {
	// Name of the filter to reorder
	FilterName string `json:"filterName"`
	// Desired position of the filter in the chain
	NewIndex int `json:"newIndex"`
	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// ReorderSourceFilterResponse contains the request body for the [ReorderSourceFilter](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter) request.
type ReorderSourceFilterResponse struct{}

// MoveSourceFilterParams contains the request body for the [MoveSourceFilter](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter) request.
type MoveSourceFilterParams struct {
	// Name of the filter to reorder
	FilterName string `json:"filterName"`
	// How to move the filter around in the source's filter chain. Either "up", "down", "top" or "bottom".
	MovementType string `json:"movementType"`
	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// MoveSourceFilterResponse contains the request body for the [MoveSourceFilter](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter) request.
type MoveSourceFilterResponse struct{}

// SetSourceFilterSettingsParams contains the request body for the [SetSourceFilterSettings](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings) request.
type SetSourceFilterSettingsParams struct {
	// Name of the filter to reconfigure
	FilterName string `json:"filterName"`
	// New settings. These will be merged to the current filter settings.
	FilterSettings map[string]interface{} `json:"filterSettings"`
	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// SetSourceFilterSettingsResponse contains the request body for the [SetSourceFilterSettings](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings) request.
type SetSourceFilterSettingsResponse struct{}
