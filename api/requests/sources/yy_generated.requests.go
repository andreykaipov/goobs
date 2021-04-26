// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourcesListParams represents the params body for the "GetSourcesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListParams struct {
	requests.Params
}

func (o *GetSourcesListParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetSourcesListParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourcesListResponse represents the response body for the "GetSourcesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListResponse struct {
	requests.Response

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

func (o *GetSourcesListResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetSourcesListResponse) GetStatus() string {
	return o.Status
}
func (o *GetSourcesListResponse) GetError() string {
	return o.Error
}

/*
GetSourcesTypesListParams represents the params body for the "GetSourcesTypesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListParams struct {
	requests.Params
}

func (o *GetSourcesTypesListParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetSourcesTypesListParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourcesTypesListResponse represents the response body for the "GetSourcesTypesList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourcesTypesList.
*/
type GetSourcesTypesListResponse struct {
	requests.Response

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

func (o *GetSourcesTypesListResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetSourcesTypesListResponse) GetStatus() string {
	return o.Status
}
func (o *GetSourcesTypesListResponse) GetError() string {
	return o.Error
}

/*
GetVolumeParams represents the params body for the "GetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

func (o *GetVolumeParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetVolumeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetVolumeResponse represents the response body for the "GetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeResponse struct {
	requests.Response

	// Indicates whether the source is muted.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`

	// Volume of the source. Between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

func (o *GetVolumeResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetVolumeResponse) GetStatus() string {
	return o.Status
}
func (o *GetVolumeResponse) GetError() string {
	return o.Error
}

/*
SetVolumeParams represents the params body for the "SetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume.
*/
type SetVolumeParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`

	// Desired volume. Must be between `0.0` and `1.0`.
	Volume float64 `json:"volume"`
}

func (o *SetVolumeParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetVolumeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetVolumeResponse represents the response body for the "SetVolume" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetVolume.
*/
type SetVolumeResponse struct {
	requests.Response
}

func (o *SetVolumeResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetVolumeResponse) GetStatus() string {
	return o.Status
}
func (o *SetVolumeResponse) GetError() string {
	return o.Error
}

/*
GetMuteParams represents the params body for the "GetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

func (o *GetMuteParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetMuteParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetMuteResponse represents the response body for the "GetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteResponse struct {
	requests.Response

	// Mute status of the source.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`
}

func (o *GetMuteResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetMuteResponse) GetStatus() string {
	return o.Status
}
func (o *GetMuteResponse) GetError() string {
	return o.Error
}

/*
SetMuteParams represents the params body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteParams struct {
	requests.Params

	// Desired mute status.
	Mute bool `json:"mute"`

	// Source name.
	Source string `json:"source"`
}

func (o *SetMuteParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetMuteParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetMuteResponse represents the response body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteResponse struct {
	requests.Response
}

func (o *SetMuteResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetMuteResponse) GetStatus() string {
	return o.Status
}
func (o *SetMuteResponse) GetError() string {
	return o.Error
}

/*
ToggleMuteParams represents the params body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

func (o *ToggleMuteParams) GetRequestType() string {
	return o.RequestType
}
func (o *ToggleMuteParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ToggleMuteResponse represents the response body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteResponse struct {
	requests.Response
}

func (o *ToggleMuteResponse) GetMessageID() string {
	return o.MessageID
}
func (o *ToggleMuteResponse) GetStatus() string {
	return o.Status
}
func (o *ToggleMuteResponse) GetError() string {
	return o.Error
}

/*
SetSyncOffsetParams represents the params body for the "SetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetParams struct {
	requests.Params

	// The desired audio sync offset (in nanoseconds).
	Offset int `json:"offset"`

	// Source name.
	Source string `json:"source"`
}

func (o *SetSyncOffsetParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetSyncOffsetParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSyncOffsetResponse represents the response body for the "SetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetResponse struct {
	requests.Response
}

func (o *SetSyncOffsetResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetSyncOffsetResponse) GetStatus() string {
	return o.Status
}
func (o *SetSyncOffsetResponse) GetError() string {
	return o.Error
}

/*
GetSyncOffsetParams represents the params body for the "GetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

func (o *GetSyncOffsetParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetSyncOffsetParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSyncOffsetResponse represents the response body for the "GetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetResponse struct {
	requests.Response

	// Source name.
	Name string `json:"name"`

	// The audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
}

func (o *GetSyncOffsetResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetSyncOffsetResponse) GetStatus() string {
	return o.Status
}
func (o *GetSyncOffsetResponse) GetError() string {
	return o.Error
}

/*
GetSourceSettingsParams represents the params body for the "GetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings.
*/
type GetSourceSettingsParams struct {
	requests.Params

	// Source name.
	SourceName string `json:"sourceName"`

	// Type of the specified source. Useful for type-checking if you expect a specific settings
	// schema.
	SourceType string `json:"sourceType"`
}

func (o *GetSourceSettingsParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetSourceSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourceSettingsResponse represents the response body for the "GetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceSettings.
*/
type GetSourceSettingsResponse struct {
	requests.Response

	// Source name
	SourceName string `json:"sourceName"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source
	SourceType string `json:"sourceType"`
}

func (o *GetSourceSettingsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetSourceSettingsResponse) GetStatus() string {
	return o.Status
}
func (o *GetSourceSettingsResponse) GetError() string {
	return o.Error
}

/*
SetSourceSettingsParams represents the params body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsParams struct {
	requests.Params

	// Source name.
	SourceName string `json:"sourceName"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source. Useful for type-checking to avoid settings a set of settings
	// incompatible with the actual source's type.
	SourceType string `json:"sourceType"`
}

func (o *SetSourceSettingsParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetSourceSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSourceSettingsResponse represents the response body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsResponse struct {
	requests.Response

	// Source name
	SourceName string `json:"sourceName"`

	// Updated source settings
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source
	SourceType string `json:"sourceType"`
}

func (o *SetSourceSettingsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetSourceSettingsResponse) GetStatus() string {
	return o.Status
}
func (o *SetSourceSettingsResponse) GetError() string {
	return o.Error
}

/*
GetTextGDIPlusPropertiesParams represents the params body for the "GetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextGDIPlusProperties.
*/
type GetTextGDIPlusPropertiesParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

func (o *GetTextGDIPlusPropertiesParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetTextGDIPlusPropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetTextGDIPlusPropertiesResponse represents the response body for the "GetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextGDIPlusProperties.
*/
type GetTextGDIPlusPropertiesResponse struct {
	requests.Response

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

func (o *GetTextGDIPlusPropertiesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetTextGDIPlusPropertiesResponse) GetStatus() string {
	return o.Status
}
func (o *GetTextGDIPlusPropertiesResponse) GetError() string {
	return o.Error
}

/*
SetTextGDIPlusPropertiesParams represents the params body for the "SetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties.
*/
type SetTextGDIPlusPropertiesParams struct {
	requests.Params

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

func (o *SetTextGDIPlusPropertiesParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetTextGDIPlusPropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetTextGDIPlusPropertiesResponse represents the response body for the "SetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties.
*/
type SetTextGDIPlusPropertiesResponse struct {
	requests.Response
}

func (o *SetTextGDIPlusPropertiesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetTextGDIPlusPropertiesResponse) GetStatus() string {
	return o.Status
}
func (o *SetTextGDIPlusPropertiesResponse) GetError() string {
	return o.Error
}

/*
GetTextFreetype2PropertiesParams represents the params body for the "GetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties.
*/
type GetTextFreetype2PropertiesParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

func (o *GetTextFreetype2PropertiesParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetTextFreetype2PropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetTextFreetype2PropertiesResponse represents the response body for the "GetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTextFreetype2Properties.
*/
type GetTextFreetype2PropertiesResponse struct {
	requests.Response

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

func (o *GetTextFreetype2PropertiesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetTextFreetype2PropertiesResponse) GetStatus() string {
	return o.Status
}
func (o *GetTextFreetype2PropertiesResponse) GetError() string {
	return o.Error
}

/*
SetTextFreetype2PropertiesParams represents the params body for the "SetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties.
*/
type SetTextFreetype2PropertiesParams struct {
	requests.Params

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

func (o *SetTextFreetype2PropertiesParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetTextFreetype2PropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetTextFreetype2PropertiesResponse represents the response body for the "SetTextFreetype2Properties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextFreetype2Properties.
*/
type SetTextFreetype2PropertiesResponse struct {
	requests.Response
}

func (o *SetTextFreetype2PropertiesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetTextFreetype2PropertiesResponse) GetStatus() string {
	return o.Status
}
func (o *SetTextFreetype2PropertiesResponse) GetError() string {
	return o.Error
}

/*
GetBrowserSourcePropertiesParams represents the params body for the "GetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
type GetBrowserSourcePropertiesParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

func (o *GetBrowserSourcePropertiesParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetBrowserSourcePropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetBrowserSourcePropertiesResponse represents the response body for the "GetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetBrowserSourceProperties.
*/
type GetBrowserSourcePropertiesResponse struct {
	requests.Response

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

func (o *GetBrowserSourcePropertiesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetBrowserSourcePropertiesResponse) GetStatus() string {
	return o.Status
}
func (o *GetBrowserSourcePropertiesResponse) GetError() string {
	return o.Error
}

/*
SetBrowserSourcePropertiesParams represents the params body for the "SetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
type SetBrowserSourcePropertiesParams struct {
	requests.Params

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

func (o *SetBrowserSourcePropertiesParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetBrowserSourcePropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetBrowserSourcePropertiesResponse represents the response body for the "SetBrowserSourceProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetBrowserSourceProperties.
*/
type SetBrowserSourcePropertiesResponse struct {
	requests.Response
}

func (o *SetBrowserSourcePropertiesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetBrowserSourcePropertiesResponse) GetStatus() string {
	return o.Status
}
func (o *SetBrowserSourcePropertiesResponse) GetError() string {
	return o.Error
}

/*
GetSpecialSourcesParams represents the params body for the "GetSpecialSources" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources.
*/
type GetSpecialSourcesParams struct {
	requests.Params
}

func (o *GetSpecialSourcesParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetSpecialSourcesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSpecialSourcesResponse represents the response body for the "GetSpecialSources" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources.
*/
type GetSpecialSourcesResponse struct {
	requests.Response

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

func (o *GetSpecialSourcesResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetSpecialSourcesResponse) GetStatus() string {
	return o.Status
}
func (o *GetSpecialSourcesResponse) GetError() string {
	return o.Error
}

/*
GetSourceFiltersParams represents the params body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
type GetSourceFiltersParams struct {
	requests.Params

	// Source name
	SourceName string `json:"sourceName"`
}

func (o *GetSourceFiltersParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetSourceFiltersParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSourceFiltersResponse represents the response body for the "GetSourceFilters" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSourceFilters.
*/
type GetSourceFiltersResponse struct {
	requests.Response

	Filters []struct {
		// Filter name
		Name string `json:"name"`

		// Filter settings
		Settings map[string]interface{} `json:"settings"`

		// Filter type
		Type string `json:"type"`
	} `json:"filters"`
}

func (o *GetSourceFiltersResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetSourceFiltersResponse) GetStatus() string {
	return o.Status
}
func (o *GetSourceFiltersResponse) GetError() string {
	return o.Error
}

/*
AddFilterToSourceParams represents the params body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
type AddFilterToSourceParams struct {
	requests.Params

	// Name of the new filter
	FilterName string `json:"filterName"`

	// Filter settings
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Filter type
	FilterType string `json:"filterType"`

	// Name of the source on which the filter is added
	SourceName string `json:"sourceName"`
}

func (o *AddFilterToSourceParams) GetRequestType() string {
	return o.RequestType
}
func (o *AddFilterToSourceParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
AddFilterToSourceResponse represents the response body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
type AddFilterToSourceResponse struct {
	requests.Response
}

func (o *AddFilterToSourceResponse) GetMessageID() string {
	return o.MessageID
}
func (o *AddFilterToSourceResponse) GetStatus() string {
	return o.Status
}
func (o *AddFilterToSourceResponse) GetError() string {
	return o.Error
}

/*
RemoveFilterFromSourceParams represents the params body for the "RemoveFilterFromSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource.
*/
type RemoveFilterFromSourceParams struct {
	requests.Params

	// Name of the filter to remove
	FilterName string `json:"filterName"`

	// Name of the source from which the specified filter is removed
	SourceName string `json:"sourceName"`
}

func (o *RemoveFilterFromSourceParams) GetRequestType() string {
	return o.RequestType
}
func (o *RemoveFilterFromSourceParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
RemoveFilterFromSourceResponse represents the response body for the "RemoveFilterFromSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#RemoveFilterFromSource.
*/
type RemoveFilterFromSourceResponse struct {
	requests.Response
}

func (o *RemoveFilterFromSourceResponse) GetMessageID() string {
	return o.MessageID
}
func (o *RemoveFilterFromSourceResponse) GetStatus() string {
	return o.Status
}
func (o *RemoveFilterFromSourceResponse) GetError() string {
	return o.Error
}

/*
ReorderSourceFilterParams represents the params body for the "ReorderSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter.
*/
type ReorderSourceFilterParams struct {
	requests.Params

	// Name of the filter to reorder
	FilterName string `json:"filterName"`

	// Desired position of the filter in the chain
	NewIndex int `json:"newIndex"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

func (o *ReorderSourceFilterParams) GetRequestType() string {
	return o.RequestType
}
func (o *ReorderSourceFilterParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ReorderSourceFilterResponse represents the response body for the "ReorderSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ReorderSourceFilter.
*/
type ReorderSourceFilterResponse struct {
	requests.Response
}

func (o *ReorderSourceFilterResponse) GetMessageID() string {
	return o.MessageID
}
func (o *ReorderSourceFilterResponse) GetStatus() string {
	return o.Status
}
func (o *ReorderSourceFilterResponse) GetError() string {
	return o.Error
}

/*
MoveSourceFilterParams represents the params body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterParams struct {
	requests.Params

	// Name of the filter to reorder
	FilterName string `json:"filterName"`

	// How to move the filter around in the source's filter chain. Either "up", "down", "top" or
	// "bottom".
	MovementType string `json:"movementType"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

func (o *MoveSourceFilterParams) GetRequestType() string {
	return o.RequestType
}
func (o *MoveSourceFilterParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
MoveSourceFilterResponse represents the response body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterResponse struct {
	requests.Response
}

func (o *MoveSourceFilterResponse) GetMessageID() string {
	return o.MessageID
}
func (o *MoveSourceFilterResponse) GetStatus() string {
	return o.Status
}
func (o *MoveSourceFilterResponse) GetError() string {
	return o.Error
}

/*
SetSourceFilterSettingsParams represents the params body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsParams struct {
	requests.Params

	// Name of the filter to reconfigure
	FilterName string `json:"filterName"`

	// New settings. These will be merged to the current filter settings.
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

func (o *SetSourceFilterSettingsParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetSourceFilterSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSourceFilterSettingsResponse represents the response body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsResponse struct {
	requests.Response
}

func (o *SetSourceFilterSettingsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetSourceFilterSettingsResponse) GetStatus() string {
	return o.Status
}
func (o *SetSourceFilterSettingsResponse) GetError() string {
	return o.Error
}
