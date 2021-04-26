// This file has been automatically generated. Don't edit it.

package sources

/*
GetSourcesListParams represents the params body for the "GetSourcesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListParams struct{}

/*
GetSourcesListResponse represents the response body for the "GetSourcesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListResponse struct {
	Sources []struct {
		// Unique source name
		Name string `json:"name"`

		// Source type. Value is one of the following: "input", "filter", "transition", "scene" or
		// "unknown"
		Type string `json:"type"`

		// Non-unique source internal type (a.k.a type id)
		TypeId string `json:"typeId"`
	} `json:"sources"`
}

/*
GetSourcesTypesListParams represents the params body for the "GetSourcesTypesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListParams struct{}

/*
GetSourcesTypesListResponse represents the response body for the "GetSourcesTypesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListResponse struct {
	Ids []struct {
		Caps struct {
			// True if interaction with this sources of this type is possible
			CanInteract bool `json:"canInteract"`

			// True if sources of this type should not be fully duplicated
			DoNotDuplicate bool `json:"doNotDuplicate"`

			// True if sources of this type may cause a feedback loop if it's audio is monitored and
			// shouldn't be
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

/*
GetVolumeParams represents the params body for the "GetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeParams struct {
	// Source name.
	Source string `json:"source"`
}

/*
GetVolumeResponse represents the response body for the "GetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeResponse struct {
	// Indicates whether the source is muted.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`

	// Volume of the source. Between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

/*
SetVolumeParams represents the params body for the "SetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume.
*/
type SetVolumeParams struct {
	// Source name.
	Source string `json:"source"`

	// Desired volume. Must be between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

/*
SetVolumeResponse represents the response body for the "SetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume.
*/
type SetVolumeResponse struct{}

/*
GetMuteParams represents the params body for the "GetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteParams struct {
	// Source name.
	Source string `json:"source"`
}

/*
GetMuteResponse represents the response body for the "GetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteResponse struct {
	// Mute status of the source.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`
}

/*
SetMuteParams represents the params body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteParams struct {
	// Desired mute status.
	Mute bool `json:"mute"`

	// Source name.
	Source string `json:"source"`
}

/*
SetMuteResponse represents the response body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteResponse struct{}

/*
ToggleMuteParams represents the params body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteParams struct {
	// Source name.
	Source string `json:"source"`
}

/*
ToggleMuteResponse represents the response body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteResponse struct{}

/*
SetSyncOffsetParams represents the params body for the "SetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetParams struct {
	// The desired audio sync offset (in nanoseconds).
	Offset int `json:"offset"`

	// Source name.
	Source string `json:"source"`
}

/*
SetSyncOffsetResponse represents the response body for the "SetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetResponse struct{}

/*
GetSyncOffsetParams represents the params body for the "GetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetParams struct {
	// Source name.
	Source string `json:"source"`
}

/*
GetSyncOffsetResponse represents the response body for the "GetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetResponse struct {
	// Source name.
	Name string `json:"name"`

	// The audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
}

/*
GetSourceSettingsParams represents the params body for the "GetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings.
*/
type GetSourceSettingsParams struct {
	// Source name.
	SourceName string `json:"sourceName"`

	// Type of the specified source. Useful for type-checking if you expect a specific settings
	// schema.
	SourceType string `json:"sourceType"`
}

/*
GetSourceSettingsResponse represents the response body for the "GetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings.
*/
type GetSourceSettingsResponse struct {
	// Source name
	SourceName string `json:"sourceName"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source
	SourceType string `json:"sourceType"`
}

/*
SetSourceSettingsParams represents the params body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsParams struct {
	// Source name.
	SourceName string `json:"sourceName"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source. Useful for type-checking to avoid settings a set of settings
	// incompatible with the actual source's type.
	SourceType string `json:"sourceType"`
}

/*
SetSourceSettingsResponse represents the response body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsResponse struct {
	// Source name
	SourceName string `json:"sourceName"`

	// Updated source settings
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source
	SourceType string `json:"sourceType"`
}

/*
GetTextGDIPlusPropertiesParams represents the params body for the "GetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextGDIPlusProperties.
*/
type GetTextGDIPlusPropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}

/*
GetTextGDIPlusPropertiesResponse represents the response body for the "GetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextGDIPlusProperties.
*/
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

/*
SetTextGDIPlusPropertiesParams represents the params body for the "SetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties.
*/
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

/*
SetTextGDIPlusPropertiesResponse represents the response body for the "SetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties.
*/
type SetTextGDIPlusPropertiesResponse struct{}

/*
GetTextFreetype2PropertiesParams represents the params body for the "GetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties.
*/
type GetTextFreetype2PropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}

/*
GetTextFreetype2PropertiesResponse represents the response body for the "GetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties.
*/
type GetTextFreetype2PropertiesResponse struct {
	// Gradient top color.
	Color1 int `json:"color1"`

	// Gradient bottom color.
	Color2 int `json:"color2"`

	// Custom width (0 to disable).
	CustomWidth int `json:"custom_width"`

	// Drop shadow.
	DropShadow bool `json:"drop_shadow"`

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

/*
SetTextFreetype2PropertiesParams represents the params body for the "SetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties.
*/
type SetTextFreetype2PropertiesParams struct {
	// Gradient top color.
	Color1 int `json:"color1"`

	// Gradient bottom color.
	Color2 int `json:"color2"`

	// Custom width (0 to disable).
	CustomWidth int `json:"custom_width"`

	// Drop shadow.
	DropShadow bool `json:"drop_shadow"`

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

/*
SetTextFreetype2PropertiesResponse represents the response body for the "SetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties.
*/
type SetTextFreetype2PropertiesResponse struct{}

/*
GetBrowserSourcePropertiesParams represents the params body for the "GetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
type GetBrowserSourcePropertiesParams struct {
	// Source name.
	Source string `json:"source"`
}

/*
GetBrowserSourcePropertiesResponse represents the response body for the "GetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
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

/*
SetBrowserSourcePropertiesParams represents the params body for the "SetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
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

/*
SetBrowserSourcePropertiesResponse represents the response body for the "SetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
type SetBrowserSourcePropertiesResponse struct{}

/*
GetSpecialSourcesParams represents the params body for the "GetSpecialSources" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources.
*/
type GetSpecialSourcesParams struct{}

/*
GetSpecialSourcesResponse represents the response body for the "GetSpecialSources" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources.
*/
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

/*
GetSourceFiltersParams represents the params body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
type GetSourceFiltersParams struct {
	// Source name
	SourceName string `json:"sourceName"`
}

/*
GetSourceFiltersResponse represents the response body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
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

/*
AddFilterToSourceParams represents the params body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
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

/*
AddFilterToSourceResponse represents the response body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
type AddFilterToSourceResponse struct{}

/*
RemoveFilterFromSourceParams represents the params body for the "RemoveFilterFromSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource.
*/
type RemoveFilterFromSourceParams struct {
	// Name of the filter to remove
	FilterName string `json:"filterName"`

	// Name of the source from which the specified filter is removed
	SourceName string `json:"sourceName"`
}

/*
RemoveFilterFromSourceResponse represents the response body for the "RemoveFilterFromSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource.
*/
type RemoveFilterFromSourceResponse struct{}

/*
ReorderSourceFilterParams represents the params body for the "ReorderSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter.
*/
type ReorderSourceFilterParams struct {
	// Name of the filter to reorder
	FilterName string `json:"filterName"`

	// Desired position of the filter in the chain
	NewIndex int `json:"newIndex"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

/*
ReorderSourceFilterResponse represents the response body for the "ReorderSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter.
*/
type ReorderSourceFilterResponse struct{}

/*
MoveSourceFilterParams represents the params body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterParams struct {
	// Name of the filter to reorder
	FilterName string `json:"filterName"`

	// How to move the filter around in the source's filter chain. Either "up", "down", "top" or
	// "bottom".
	MovementType string `json:"movementType"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

/*
MoveSourceFilterResponse represents the response body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterResponse struct{}

/*
SetSourceFilterSettingsParams represents the params body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsParams struct {
	// Name of the filter to reconfigure
	FilterName string `json:"filterName"`

	// New settings. These will be merged to the current filter settings.
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

/*
SetSourceFilterSettingsResponse represents the response body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsResponse struct{}
