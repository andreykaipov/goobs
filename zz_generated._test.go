// This file has been automatically generated. Don't edit it.

package goobs_test

import (
	"net"
	"net/http"
	"os"
	"testing"

	goobs "github.com/andreykaipov/goobs"
	config "github.com/andreykaipov/goobs/api/requests/config"
	filters "github.com/andreykaipov/goobs/api/requests/filters"
	general "github.com/andreykaipov/goobs/api/requests/general"
	inputs "github.com/andreykaipov/goobs/api/requests/inputs"
	mediainputs "github.com/andreykaipov/goobs/api/requests/mediainputs"
	outputs "github.com/andreykaipov/goobs/api/requests/outputs"
	record "github.com/andreykaipov/goobs/api/requests/record"
	sceneitems "github.com/andreykaipov/goobs/api/requests/sceneitems"
	scenes "github.com/andreykaipov/goobs/api/requests/scenes"
	sources "github.com/andreykaipov/goobs/api/requests/sources"
	stream "github.com/andreykaipov/goobs/api/requests/stream"
	transitions "github.com/andreykaipov/goobs/api/requests/transitions"
	ui "github.com/andreykaipov/goobs/api/requests/ui"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
	websocket "github.com/gorilla/websocket"
	assert "github.com/stretchr/testify/assert"
)

func Test_client(t *testing.T) {
	var err error
	_, err = goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("wrongpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.Error(t, err)
	assert.IsType(t, &websocket.CloseError{}, err)
	assert.Equal(t, err.(*websocket.CloseError).Code, 4009)
	_, err = goobs.New(
		"localhost:42069",
		goobs.WithPassword("wrongpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.Error(t, err)
	assert.IsType(t, &net.OpError{}, err)
}

func Test_config(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Config.CreateProfile(&config.CreateProfileParams{ProfileName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "CreateProfile", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.CreateSceneCollection(&config.CreateSceneCollectionParams{SceneCollectionName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "CreateSceneCollection", err)
	}
	assert.Error(t, err)
	_, err = client.Config.GetPersistentData(&config.GetPersistentDataParams{
		Realm:    "OBS_WEBSOCKET_DATA_REALM_GLOBAL",
		SlotName: "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetPersistentData", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetProfileList(&config.GetProfileListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetProfileList", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetProfileParameter(&config.GetProfileParameterParams{
		ParameterCategory: "test",
		ParameterName:     "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetProfileParameter", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetRecordDirectory(&config.GetRecordDirectoryParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetRecordDirectory", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetSceneCollectionList(&config.GetSceneCollectionListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneCollectionList", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetStreamServiceSettings(&config.GetStreamServiceSettingsParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetStreamServiceSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetVideoSettings(&config.GetVideoSettingsParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetVideoSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetCurrentProfile(&config.SetCurrentProfileParams{ProfileName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "SetCurrentProfile", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetCurrentSceneCollection(
		&config.SetCurrentSceneCollectionParams{SceneCollectionName: "test"},
	)
	if err != nil {
		t.Logf("%s\n%s\n", "SetCurrentSceneCollection", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetPersistentData(&config.SetPersistentDataParams{
		Realm:     "OBS_WEBSOCKET_DATA_REALM_GLOBAL",
		SlotName:  "test",
		SlotValue: "",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetPersistentData", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetProfileParameter(&config.SetProfileParameterParams{
		ParameterCategory: "test",
		ParameterName:     "test",
		ParameterValue:    "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetProfileParameter", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetStreamServiceSettings(&config.SetStreamServiceSettingsParams{
		StreamServiceSettings: &typedefs.StreamServiceSettings{},
		StreamServiceType:     "rtmp_custom",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetStreamServiceSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetVideoSettings(&config.SetVideoSettingsParams{
		BaseHeight:     10.0,
		BaseWidth:      10.0,
		FpsDenominator: 10.0,
		FpsNumerator:   10.0,
		OutputHeight:   10.0,
		OutputWidth:    10.0,
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetVideoSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.RemoveProfile(&config.RemoveProfileParams{ProfileName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "RemoveProfile", err)
	}
	assert.NoError(t, err)
}

func Test_filters(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Filters.CreateSourceFilter(&filters.CreateSourceFilterParams{
		FilterKind:     "scroll_filter",
		FilterName:     "test",
		FilterSettings: map[string]interface{}{"test": "test"},
		SourceName:     "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "CreateSourceFilter", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilterDefaultSettings(
		&filters.GetSourceFilterDefaultSettingsParams{FilterKind: "scroll_filter"},
	)
	if err != nil {
		t.Logf("%s\n%s\n", "GetSourceFilterDefaultSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilterList(&filters.GetSourceFilterListParams{SourceName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSourceFilterList", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilter(&filters.GetSourceFilterParams{
		FilterName: "test",
		SourceName: "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSourceFilter", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterEnabled(&filters.SetSourceFilterEnabledParams{
		FilterEnabled: &[]bool{true}[0],
		FilterName:    "test",
		SourceName:    "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSourceFilterEnabled", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterIndex(&filters.SetSourceFilterIndexParams{
		FilterIndex: 1.0,
		FilterName:  "test",
		SourceName:  "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSourceFilterIndex", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterName(&filters.SetSourceFilterNameParams{
		FilterName:    "test",
		NewFilterName: "test",
		SourceName:    "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSourceFilterName", err)
	}
	assert.Error(t, err)
	_, err = client.Filters.SetSourceFilterSettings(&filters.SetSourceFilterSettingsParams{
		FilterName:     "test",
		FilterSettings: map[string]interface{}{"test": "test"},
		Overlay:        &[]bool{true}[0],
		SourceName:     "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSourceFilterSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.RemoveSourceFilter(&filters.RemoveSourceFilterParams{
		FilterName: "test",
		SourceName: "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "RemoveSourceFilter", err)
	}
	assert.NoError(t, err)
}

func Test_general(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.General.BroadcastCustomEvent(
		&general.BroadcastCustomEventParams{EventData: map[string]bool{"test": true}},
	)
	if err != nil {
		t.Logf("%s\n%s\n", "BroadcastCustomEvent", err)
	}
	assert.NoError(t, err)
	_, err = client.General.CallVendorRequest(&general.CallVendorRequestParams{
		RequestData: "",
		RequestType: "test",
		VendorName:  "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "CallVendorRequest", err)
	}
	assert.Error(t, err)
	_, err = client.General.GetHotkeyList(&general.GetHotkeyListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetHotkeyList", err)
	}
	assert.NoError(t, err)
	_, err = client.General.GetStats(&general.GetStatsParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetStats", err)
	}
	assert.NoError(t, err)
	_, err = client.General.GetVersion(&general.GetVersionParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetVersion", err)
	}
	assert.NoError(t, err)
	_, err = client.General.Sleep(&general.SleepParams{
		SleepFrames: 1.0,
		SleepMillis: 1.0,
	})
	if err != nil {
		t.Logf("%s\n%s\n", "Sleep", err)
	}
	assert.Error(t, err)
	_, err = client.General.TriggerHotkeyByKeySequence(&general.TriggerHotkeyByKeySequenceParams{
		KeyId:        "OBS_KEY_SHIFT",
		KeyModifiers: &typedefs.KeyModifiers{},
	})
	if err != nil {
		t.Logf("%s\n%s\n", "TriggerHotkeyByKeySequence", err)
	}
	assert.NoError(t, err)
	_, err = client.General.TriggerHotkeyByName(
		&general.TriggerHotkeyByNameParams{HotkeyName: "OBSBasic.ShowContextBar"},
	)
	if err != nil {
		t.Logf("%s\n%s\n", "TriggerHotkeyByName", err)
	}
	assert.NoError(t, err)
}

func Test_inputs(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Inputs.CreateInput(&inputs.CreateInputParams{
		InputKind:        "ffmpeg_source",
		InputName:        "test2",
		InputSettings:    map[string]interface{}{"test": "test"},
		SceneItemEnabled: &[]bool{true}[0],
		SceneName:        "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "CreateInput", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioBalance(&inputs.GetInputAudioBalanceParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputAudioBalance", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioMonitorType(&inputs.GetInputAudioMonitorTypeParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputAudioMonitorType", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioSyncOffset(&inputs.GetInputAudioSyncOffsetParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputAudioSyncOffset", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioTracks(&inputs.GetInputAudioTracksParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputAudioTracks", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputDefaultSettings(&inputs.GetInputDefaultSettingsParams{InputKind: "ffmpeg_source"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputDefaultSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputKindList(&inputs.GetInputKindListParams{Unversioned: &[]bool{true}[0]})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputKindList", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputList(&inputs.GetInputListParams{InputKind: "ffmpeg_source"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputList", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputMute(&inputs.GetInputMuteParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputMute", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputPropertiesListPropertyItems(&inputs.GetInputPropertiesListPropertyItemsParams{
		InputName:    "test2",
		PropertyName: "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputPropertiesListPropertyItems", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.GetInputSettings(&inputs.GetInputSettingsParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputVolume(&inputs.GetInputVolumeParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetInputVolume", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetSpecialInputs(&inputs.GetSpecialInputsParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSpecialInputs", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.PressInputPropertiesButton(&inputs.PressInputPropertiesButtonParams{
		InputName:    "test2",
		PropertyName: "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "PressInputPropertiesButton", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputAudioBalance(&inputs.SetInputAudioBalanceParams{
		InputAudioBalance: 1.0,
		InputName:         "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputAudioBalance", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputAudioMonitorType(&inputs.SetInputAudioMonitorTypeParams{
		InputName:   "test2",
		MonitorType: "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputAudioMonitorType", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputAudioSyncOffset(&inputs.SetInputAudioSyncOffsetParams{
		InputAudioSyncOffset: 1.0,
		InputName:            "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputAudioSyncOffset", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputAudioTracks(&inputs.SetInputAudioTracksParams{
		InputAudioTracks: &typedefs.InputAudioTracks{"test": true},
		InputName:        "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputAudioTracks", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputMute(&inputs.SetInputMuteParams{
		InputMuted: &[]bool{true}[0],
		InputName:  "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputMute", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputName(&inputs.SetInputNameParams{
		InputName:    "test2",
		NewInputName: "test2",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputName", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputSettings(&inputs.SetInputSettingsParams{
		InputName:     "test2",
		InputSettings: map[string]interface{}{"test": "test"},
		Overlay:       &[]bool{true}[0],
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputSettings", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputVolume(&inputs.SetInputVolumeParams{
		InputName:      "test2",
		InputVolumeDb:  1.0,
		InputVolumeMul: 1.0,
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetInputVolume", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.ToggleInputMute(&inputs.ToggleInputMuteParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "ToggleInputMute", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.RemoveInput(&inputs.RemoveInputParams{InputName: "test2"})
	if err != nil {
		t.Logf("%s\n%s\n", "RemoveInput", err)
	}
	assert.NoError(t, err)
}

func Test_mediainputs(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.MediaInputs.GetMediaInputStatus(&mediainputs.GetMediaInputStatusParams{InputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetMediaInputStatus", err)
	}
	assert.NoError(t, err)
	_, err = client.MediaInputs.OffsetMediaInputCursor(&mediainputs.OffsetMediaInputCursorParams{
		InputName:         "test",
		MediaCursorOffset: 1.0,
	})
	if err != nil {
		t.Logf("%s\n%s\n", "OffsetMediaInputCursor", err)
	}
	assert.NoError(t, err)
	_, err = client.MediaInputs.SetMediaInputCursor(&mediainputs.SetMediaInputCursorParams{
		InputName:   "test",
		MediaCursor: 1.0,
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetMediaInputCursor", err)
	}
	assert.NoError(t, err)
	_, err = client.MediaInputs.TriggerMediaInputAction(&mediainputs.TriggerMediaInputActionParams{
		InputName:   "test",
		MediaAction: "OBS_WEBSOCKET_MEDIA_INPUT_ACTION_PAUSE",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "TriggerMediaInputAction", err)
	}
	assert.NoError(t, err)
}

func Test_outputs(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Outputs.GetLastReplayBufferReplay(&outputs.GetLastReplayBufferReplayParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetLastReplayBufferReplay", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetOutputList(&outputs.GetOutputListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetOutputList", err)
	}
	assert.NoError(t, err)
	_, err = client.Outputs.GetOutputSettings(&outputs.GetOutputSettingsParams{OutputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetOutputSettings", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetOutputStatus(&outputs.GetOutputStatusParams{OutputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetOutputStatus", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetReplayBufferStatus(&outputs.GetReplayBufferStatusParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetReplayBufferStatus", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetVirtualCamStatus(&outputs.GetVirtualCamStatusParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetVirtualCamStatus", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.SaveReplayBuffer(&outputs.SaveReplayBufferParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "SaveReplayBuffer", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.SetOutputSettings(&outputs.SetOutputSettingsParams{
		OutputName:     "test",
		OutputSettings: "",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetOutputSettings", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StartOutput(&outputs.StartOutputParams{OutputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "StartOutput", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StartReplayBuffer(&outputs.StartReplayBufferParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StartReplayBuffer", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StartVirtualCam(&outputs.StartVirtualCamParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StartVirtualCam", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StopOutput(&outputs.StopOutputParams{OutputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "StopOutput", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StopReplayBuffer(&outputs.StopReplayBufferParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StopReplayBuffer", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StopVirtualCam(&outputs.StopVirtualCamParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StopVirtualCam", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.ToggleOutput(&outputs.ToggleOutputParams{OutputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "ToggleOutput", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.ToggleReplayBuffer(&outputs.ToggleReplayBufferParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "ToggleReplayBuffer", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.ToggleVirtualCam(&outputs.ToggleVirtualCamParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "ToggleVirtualCam", err)
	}
	assert.Error(t, err)
}

func Test_record(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Record.GetRecordStatus(&record.GetRecordStatusParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetRecordStatus", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.PauseRecord(&record.PauseRecordParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "PauseRecord", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.ResumeRecord(&record.ResumeRecordParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "ResumeRecord", err)
	}
	assert.Error(t, err)
	_, err = client.Record.StartRecord(&record.StartRecordParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StartRecord", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.StopRecord(&record.StopRecordParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StopRecord", err)
	}
	assert.Error(t, err)
	_, err = client.Record.ToggleRecord(&record.ToggleRecordParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "ToggleRecord", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.ToggleRecordPause(&record.ToggleRecordPauseParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "ToggleRecordPause", err)
	}
	assert.NoError(t, err)
}

func Test_sceneitems(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.SceneItems.CreateSceneItem(&sceneitems.CreateSceneItemParams{
		SceneItemEnabled: &[]bool{true}[0],
		SceneName:        "Scene",
		SourceName:       "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "CreateSceneItem", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.DuplicateSceneItem(&sceneitems.DuplicateSceneItemParams{
		DestinationSceneName: "test",
		SceneItemId:          1.0,
		SceneName:            "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "DuplicateSceneItem", err)
	}
	assert.Error(t, err)
	_, err = client.SceneItems.GetGroupSceneItemList(&sceneitems.GetGroupSceneItemListParams{SceneName: "Scene"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetGroupSceneItemList", err)
	}
	assert.Error(t, err)
	_, err = client.SceneItems.GetSceneItemBlendMode(&sceneitems.GetSceneItemBlendModeParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneItemBlendMode", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemEnabled(&sceneitems.GetSceneItemEnabledParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneItemEnabled", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemId(&sceneitems.GetSceneItemIdParams{
		SceneName:    "Scene",
		SearchOffset: 1.0,
		SourceName:   "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneItemId", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemIndex(&sceneitems.GetSceneItemIndexParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneItemIndex", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{SceneName: "Scene"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneItemList", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemLocked(&sceneitems.GetSceneItemLockedParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneItemLocked", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemTransform(&sceneitems.GetSceneItemTransformParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneItemTransform", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemBlendMode(&sceneitems.SetSceneItemBlendModeParams{
		SceneItemBlendMode: "OBS_BLEND_NORMAL",
		SceneItemId:        1.0,
		SceneName:          "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSceneItemBlendMode", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemEnabled(&sceneitems.SetSceneItemEnabledParams{
		SceneItemEnabled: &[]bool{true}[0],
		SceneItemId:      1.0,
		SceneName:        "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSceneItemEnabled", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemIndex(&sceneitems.SetSceneItemIndexParams{
		SceneItemId:    1.0,
		SceneItemIndex: 1.0,
		SceneName:      "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSceneItemIndex", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemLocked(&sceneitems.SetSceneItemLockedParams{
		SceneItemId:     1.0,
		SceneItemLocked: &[]bool{true}[0],
		SceneName:       "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSceneItemLocked", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemTransform(&sceneitems.SetSceneItemTransformParams{
		SceneItemId: 1.0,
		SceneItemTransform: &typedefs.SceneItemTransform{
			BoundsHeight: 1.0,
			BoundsType:   "OBS_BOUNDS_NONE",
			BoundsWidth:  1.0,
		},
		SceneName: "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSceneItemTransform", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.RemoveSceneItem(&sceneitems.RemoveSceneItemParams{
		SceneItemId: 1.0,
		SceneName:   "nonexistent",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "RemoveSceneItem", err)
	}
	assert.Error(t, err)
}

func Test_scenes(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: "Scene"})
	if err != nil {
		t.Logf("%s\n%s\n", "CreateScene", err)
	}
	assert.Error(t, err)
	_, err = client.Scenes.GetCurrentPreviewScene(&scenes.GetCurrentPreviewSceneParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetCurrentPreviewScene", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetCurrentProgramScene(&scenes.GetCurrentProgramSceneParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetCurrentProgramScene", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetGroupList(&scenes.GetGroupListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetGroupList", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetSceneList(&scenes.GetSceneListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneList", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetSceneSceneTransitionOverride(
		&scenes.GetSceneSceneTransitionOverrideParams{SceneName: "Scene"},
	)
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneSceneTransitionOverride", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.SetCurrentPreviewScene(&scenes.SetCurrentPreviewSceneParams{SceneName: "Scene"})
	if err != nil {
		t.Logf("%s\n%s\n", "SetCurrentPreviewScene", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{SceneName: "Scene"})
	if err != nil {
		t.Logf("%s\n%s\n", "SetCurrentProgramScene", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.SetSceneName(&scenes.SetSceneNameParams{
		NewSceneName: "Scene",
		SceneName:    "Scene",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSceneName", err)
	}
	assert.Error(t, err)
	_, err = client.Scenes.SetSceneSceneTransitionOverride(&scenes.SetSceneSceneTransitionOverrideParams{
		SceneName:          "Scene",
		TransitionDuration: 50.0,
		TransitionName:     "Cut",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetSceneSceneTransitionOverride", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.RemoveScene(&scenes.RemoveSceneParams{SceneName: "nonexistent"})
	if err != nil {
		t.Logf("%s\n%s\n", "RemoveScene", err)
	}
	assert.Error(t, err)
}

func Test_sources(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Sources.GetSourceActive(&sources.GetSourceActiveParams{SourceName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSourceActive", err)
	}
	assert.NoError(t, err)
	_, err = client.Sources.GetSourceScreenshot(&sources.GetSourceScreenshotParams{
		ImageCompressionQuality: 8.0,
		ImageFormat:             "png",
		ImageHeight:             8.0,
		ImageWidth:              8.0,
		SourceName:              "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSourceScreenshot", err)
	}
	assert.NoError(t, err)
	_, err = client.Sources.SaveSourceScreenshot(&sources.SaveSourceScreenshotParams{
		ImageCompressionQuality: 8.0,
		ImageFilePath:           "test",
		ImageFormat:             "png",
		ImageHeight:             8.0,
		ImageWidth:              8.0,
		SourceName:              "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SaveSourceScreenshot", err)
	}
	assert.NoError(t, err)
}

func Test_stream(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Stream.GetStreamStatus(&stream.GetStreamStatusParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetStreamStatus", err)
	}
	assert.NoError(t, err)
	_, err = client.Stream.SendStreamCaption(&stream.SendStreamCaptionParams{CaptionText: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "SendStreamCaption", err)
	}
	assert.Error(t, err)
	_, err = client.Stream.StartStream(&stream.StartStreamParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StartStream", err)
	}
	assert.NoError(t, err)
	_, err = client.Stream.StopStream(&stream.StopStreamParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "StopStream", err)
	}
	assert.Error(t, err)
	_, err = client.Stream.ToggleStream(&stream.ToggleStreamParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "ToggleStream", err)
	}
	assert.NoError(t, err)
}

func Test_transitions(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Transitions.GetCurrentSceneTransitionCursor(&transitions.GetCurrentSceneTransitionCursorParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetCurrentSceneTransitionCursor", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.GetCurrentSceneTransition(&transitions.GetCurrentSceneTransitionParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetCurrentSceneTransition", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.GetSceneTransitionList(&transitions.GetSceneTransitionListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetSceneTransitionList", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.GetTransitionKindList(&transitions.GetTransitionKindListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetTransitionKindList", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransitionDuration(
		&transitions.SetCurrentSceneTransitionDurationParams{TransitionDuration: 50.0},
	)
	if err != nil {
		t.Logf("%s\n%s\n", "SetCurrentSceneTransitionDuration", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransition(
		&transitions.SetCurrentSceneTransitionParams{TransitionName: "Cut"},
	)
	if err != nil {
		t.Logf("%s\n%s\n", "SetCurrentSceneTransition", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransitionSettings(&transitions.SetCurrentSceneTransitionSettingsParams{
		Overlay:            &[]bool{true}[0],
		TransitionSettings: map[string]interface{}{"test": "test"},
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetCurrentSceneTransitionSettings", err)
	}
	assert.Error(t, err)
	_, err = client.Transitions.SetTBarPosition(&transitions.SetTBarPositionParams{
		Position: 1.0,
		Release:  &[]bool{true}[0],
	})
	if err != nil {
		t.Logf("%s\n%s\n", "SetTBarPosition", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.TriggerStudioModeTransition(&transitions.TriggerStudioModeTransitionParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "TriggerStudioModeTransition", err)
	}
	assert.NoError(t, err)
}

func Test_ui(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Ui.GetMonitorList(&ui.GetMonitorListParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetMonitorList", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.GetStudioModeEnabled(&ui.GetStudioModeEnabledParams{})
	if err != nil {
		t.Logf("%s\n%s\n", "GetStudioModeEnabled", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenInputFiltersDialog(&ui.OpenInputFiltersDialogParams{InputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "OpenInputFiltersDialog", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenInputInteractDialog(&ui.OpenInputInteractDialogParams{InputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "OpenInputInteractDialog", err)
	}
	assert.Error(t, err)
	_, err = client.Ui.OpenInputPropertiesDialog(&ui.OpenInputPropertiesDialogParams{InputName: "test"})
	if err != nil {
		t.Logf("%s\n%s\n", "OpenInputPropertiesDialog", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenSourceProjector(&ui.OpenSourceProjectorParams{
		MonitorIndex:      1.0,
		ProjectorGeometry: "",
		SourceName:        "test",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "OpenSourceProjector", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenVideoMixProjector(&ui.OpenVideoMixProjectorParams{
		MonitorIndex:      1.0,
		ProjectorGeometry: "",
		VideoMixType:      "OBS_WEBSOCKET_VIDEO_MIX_TYPE_PREVIEW",
	})
	if err != nil {
		t.Logf("%s\n%s\n", "OpenVideoMixProjector", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.SetStudioModeEnabled(&ui.SetStudioModeEnabledParams{StudioModeEnabled: &[]bool{true}[0]})
	if err != nil {
		t.Logf("%s\n%s\n", "SetStudioModeEnabled", err)
	}
	assert.NoError(t, err)
}
