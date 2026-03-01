// This file has been automatically generated. Don't edit it.

package goobs_test

import (
	"net/http"
	"os"
	"testing"

	goobs "github.com/andreykaipov/goobs"
	canvas "github.com/andreykaipov/goobs/api/requests/canvas"
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
	assert "github.com/stretchr/testify/assert"
)

func Test_canvas(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Canvas.GetCanvasList(&canvas.GetCanvasListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
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

	_, err = client.Config.CreateProfile(&config.CreateProfileParams{ProfileName: &[]string{"docker"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Config.CreateSceneCollection(
		&config.CreateSceneCollectionParams{SceneCollectionName: &[]string{"test"}[0]},
	)
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Config.GetPersistentData(&config.GetPersistentDataParams{
		Realm:    &[]string{"OBS_WEBSOCKET_DATA_REALM_GLOBAL"}[0],
		SlotName: &[]string{"test"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetProfileList(&config.GetProfileListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetProfileParameter(&config.GetProfileParameterParams{
		ParameterCategory: &[]string{"test"}[0],
		ParameterName:     &[]string{"test"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetRecordDirectory(&config.GetRecordDirectoryParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetSceneCollectionList(&config.GetSceneCollectionListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetStreamServiceSettings(&config.GetStreamServiceSettingsParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.GetVideoSettings(&config.GetVideoSettingsParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetCurrentProfile(&config.SetCurrentProfileParams{ProfileName: &[]string{"docker"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetCurrentSceneCollection(
		&config.SetCurrentSceneCollectionParams{SceneCollectionName: &[]string{"test"}[0]},
	)
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetPersistentData(&config.SetPersistentDataParams{
		Realm:     &[]string{"OBS_WEBSOCKET_DATA_REALM_GLOBAL"}[0],
		SlotName:  &[]string{"test"}[0],
		SlotValue: "",
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetProfileParameter(&config.SetProfileParameterParams{
		ParameterCategory: &[]string{"test"}[0],
		ParameterName:     &[]string{"test"}[0],
		ParameterValue:    &[]string{"test"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetRecordDirectory(&config.SetRecordDirectoryParams{RecordDirectory: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetStreamServiceSettings(&config.SetStreamServiceSettingsParams{
		StreamServiceSettings: &typedefs.StreamServiceSettings{},
		StreamServiceType:     &[]string{"rtmp_custom"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.SetVideoSettings(&config.SetVideoSettingsParams{
		BaseHeight:     &[]float64{10.0}[0],
		BaseWidth:      &[]float64{10.0}[0],
		FpsDenominator: &[]float64{10.0}[0],
		FpsNumerator:   &[]float64{10.0}[0],
		OutputHeight:   &[]float64{10.0}[0],
		OutputWidth:    &[]float64{10.0}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Config.RemoveProfile(&config.RemoveProfileParams{ProfileName: &[]string{"docker"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
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
		CanvasName:     &[]string{"test"}[0],
		CanvasUuid:     nil,
		FilterKind:     &[]string{"scroll_filter"}[0],
		FilterName:     &[]string{"test"}[0],
		FilterSettings: map[string]interface{}{"test": "test"},
		SourceName:     &[]string{"test"}[0],
		SourceUuid:     nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilterDefaultSettings(
		&filters.GetSourceFilterDefaultSettingsParams{FilterKind: &[]string{"scroll_filter"}[0]},
	)
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilterKindList(&filters.GetSourceFilterKindListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilterList(&filters.GetSourceFilterListParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		SourceName: &[]string{"test"}[0],
		SourceUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilter(&filters.GetSourceFilterParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		FilterName: &[]string{"test"}[0],
		SourceName: &[]string{"test"}[0],
		SourceUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterEnabled(&filters.SetSourceFilterEnabledParams{
		CanvasName:    &[]string{"test"}[0],
		CanvasUuid:    nil,
		FilterEnabled: &[]bool{true}[0],
		FilterName:    &[]string{"test"}[0],
		SourceName:    &[]string{"test"}[0],
		SourceUuid:    nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterIndex(&filters.SetSourceFilterIndexParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		FilterIndex: &[]int{1}[0],
		FilterName:  &[]string{"test"}[0],
		SourceName:  &[]string{"test"}[0],
		SourceUuid:  nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterName(&filters.SetSourceFilterNameParams{
		CanvasName:    &[]string{"test"}[0],
		CanvasUuid:    nil,
		FilterName:    &[]string{"test"}[0],
		NewFilterName: &[]string{"test"}[0],
		SourceName:    &[]string{"test"}[0],
		SourceUuid:    nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Filters.SetSourceFilterSettings(&filters.SetSourceFilterSettingsParams{
		CanvasName:     &[]string{"test"}[0],
		CanvasUuid:     nil,
		FilterName:     &[]string{"test"}[0],
		FilterSettings: map[string]interface{}{"test": "test"},
		Overlay:        &[]bool{true}[0],
		SourceName:     &[]string{"test"}[0],
		SourceUuid:     nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Filters.RemoveSourceFilter(&filters.RemoveSourceFilterParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		FilterName: &[]string{"test"}[0],
		SourceName: &[]string{"test"}[0],
		SourceUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
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
		&general.BroadcastCustomEventParams{EventData: map[string]interface{}{"test": "test"}},
	)
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.General.CallVendorRequest(&general.CallVendorRequestParams{
		RequestData: map[string]interface{}{"test": "test"},
		RequestType: &[]string{"test"}[0],
		VendorName:  &[]string{"test"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.General.GetHotkeyList(&general.GetHotkeyListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.General.GetStats(&general.GetStatsParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.General.GetVersion(&general.GetVersionParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.General.Sleep(&general.SleepParams{
		SleepFrames: &[]float64{1.0}[0],
		SleepMillis: &[]float64{1.0}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.General.TriggerHotkeyByKeySequence(&general.TriggerHotkeyByKeySequenceParams{
		KeyId:        &[]string{"OBS_KEY_SHIFT"}[0],
		KeyModifiers: &typedefs.KeyModifiers{},
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.General.TriggerHotkeyByName(&general.TriggerHotkeyByNameParams{
		ContextName: &[]string{"test"}[0],
		HotkeyName:  &[]string{"OBSBasic.ShowContextBar"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
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
		CanvasName:       &[]string{"test2"}[0],
		CanvasUuid:       nil,
		InputKind:        &[]string{"ffmpeg_source"}[0],
		InputName:        &[]string{"test2"}[0],
		InputSettings:    map[string]interface{}{"test": "test"},
		SceneItemEnabled: &[]bool{true}[0],
		SceneName:        &[]string{"Scene"}[0],
		SceneUuid:        nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioBalance(&inputs.GetInputAudioBalanceParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioMonitorType(&inputs.GetInputAudioMonitorTypeParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioSyncOffset(&inputs.GetInputAudioSyncOffsetParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioTracks(&inputs.GetInputAudioTracksParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputDefaultSettings(
		&inputs.GetInputDefaultSettingsParams{InputKind: &[]string{"ffmpeg_source"}[0]},
	)
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputDeinterlaceFieldOrder(&inputs.GetInputDeinterlaceFieldOrderParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputDeinterlaceMode(&inputs.GetInputDeinterlaceModeParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputKindList(&inputs.GetInputKindListParams{Unversioned: &[]bool{true}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputList(&inputs.GetInputListParams{InputKind: &[]string{"ffmpeg_source"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputMute(&inputs.GetInputMuteParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputPropertiesListPropertyItems(&inputs.GetInputPropertiesListPropertyItemsParams{
		InputName:    &[]string{"test2"}[0],
		InputUuid:    nil,
		PropertyName: &[]string{"test2"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.GetInputSettings(&inputs.GetInputSettingsParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputVolume(&inputs.GetInputVolumeParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.GetSpecialInputs(&inputs.GetSpecialInputsParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.PressInputPropertiesButton(&inputs.PressInputPropertiesButtonParams{
		InputName:    &[]string{"test2"}[0],
		InputUuid:    nil,
		PropertyName: &[]string{"test2"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputAudioBalance(&inputs.SetInputAudioBalanceParams{
		InputAudioBalance: &[]float64{1.0}[0],
		InputName:         &[]string{"test2"}[0],
		InputUuid:         nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputAudioMonitorType(&inputs.SetInputAudioMonitorTypeParams{
		InputName:   &[]string{"test2"}[0],
		InputUuid:   nil,
		MonitorType: &[]string{"test2"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputAudioSyncOffset(&inputs.SetInputAudioSyncOffsetParams{
		InputAudioSyncOffset: &[]float64{1.0}[0],
		InputName:            &[]string{"test2"}[0],
		InputUuid:            nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputAudioTracks(&inputs.SetInputAudioTracksParams{
		InputAudioTracks: &typedefs.InputAudioTracks{"test": true},
		InputName:        &[]string{"test2"}[0],
		InputUuid:        nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputDeinterlaceFieldOrder(&inputs.SetInputDeinterlaceFieldOrderParams{
		InputDeinterlaceFieldOrder: &[]string{"test2"}[0],
		InputName:                  &[]string{"test2"}[0],
		InputUuid:                  nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputDeinterlaceMode(&inputs.SetInputDeinterlaceModeParams{
		InputDeinterlaceMode: &[]string{"test2"}[0],
		InputName:            &[]string{"test2"}[0],
		InputUuid:            nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputMute(&inputs.SetInputMuteParams{
		InputMuted: &[]bool{true}[0],
		InputName:  &[]string{"test2"}[0],
		InputUuid:  nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputName(&inputs.SetInputNameParams{
		InputName:    &[]string{"test2"}[0],
		InputUuid:    nil,
		NewInputName: &[]string{"test2"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.SetInputSettings(&inputs.SetInputSettingsParams{
		InputName:     &[]string{"test2"}[0],
		InputSettings: map[string]interface{}{"test": "test"},
		InputUuid:     nil,
		Overlay:       &[]bool{true}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputVolume(&inputs.SetInputVolumeParams{
		InputName:      &[]string{"test2"}[0],
		InputUuid:      nil,
		InputVolumeDb:  &[]float64{1.0}[0],
		InputVolumeMul: &[]float64{1.0}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Inputs.ToggleInputMute(&inputs.ToggleInputMuteParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Inputs.RemoveInput(&inputs.RemoveInputParams{
		InputName: &[]string{"test2"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
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

	_, err = client.MediaInputs.GetMediaInputStatus(&mediainputs.GetMediaInputStatusParams{
		InputName: &[]string{"test"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.MediaInputs.OffsetMediaInputCursor(&mediainputs.OffsetMediaInputCursorParams{
		InputName:         &[]string{"test"}[0],
		InputUuid:         nil,
		MediaCursorOffset: &[]float64{1.0}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.MediaInputs.SetMediaInputCursor(&mediainputs.SetMediaInputCursorParams{
		InputName:   &[]string{"test"}[0],
		InputUuid:   nil,
		MediaCursor: &[]float64{1.0}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.MediaInputs.TriggerMediaInputAction(&mediainputs.TriggerMediaInputActionParams{
		InputName:   &[]string{"test"}[0],
		InputUuid:   nil,
		MediaAction: &[]string{"OBS_WEBSOCKET_MEDIA_INPUT_ACTION_PAUSE"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
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
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetOutputList(&outputs.GetOutputListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Outputs.GetOutputSettings(&outputs.GetOutputSettingsParams{OutputName: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetOutputStatus(&outputs.GetOutputStatusParams{OutputName: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetReplayBufferStatus(&outputs.GetReplayBufferStatusParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.GetVirtualCamStatus(&outputs.GetVirtualCamStatusParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.SaveReplayBuffer(&outputs.SaveReplayBufferParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.SetOutputSettings(&outputs.SetOutputSettingsParams{
		OutputName:     &[]string{"test"}[0],
		OutputSettings: map[string]interface{}{"test": "test"},
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StartOutput(&outputs.StartOutputParams{OutputName: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StartReplayBuffer(&outputs.StartReplayBufferParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StartVirtualCam(&outputs.StartVirtualCamParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StopOutput(&outputs.StopOutputParams{OutputName: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StopReplayBuffer(&outputs.StopReplayBufferParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.StopVirtualCam(&outputs.StopVirtualCamParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.ToggleOutput(&outputs.ToggleOutputParams{OutputName: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.ToggleReplayBuffer(&outputs.ToggleReplayBufferParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Outputs.ToggleVirtualCam(&outputs.ToggleVirtualCamParams{})
	if err != nil {
		t.Logf("%s", err)
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

	_, err = client.Record.CreateRecordChapter(&record.CreateRecordChapterParams{ChapterName: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Record.GetRecordStatus(&record.GetRecordStatusParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.PauseRecord(&record.PauseRecordParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.ResumeRecord(&record.ResumeRecordParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Record.SplitRecordFile(&record.SplitRecordFileParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Record.StartRecord(&record.StartRecordParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.StopRecord(&record.StopRecordParams{})
	t.Logf("skipped: %s", err)
	_, err = client.Record.ToggleRecord(&record.ToggleRecordParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Record.ToggleRecordPause(&record.ToggleRecordPauseParams{})
	if err != nil {
		t.Logf("%s", err)
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
		CanvasName:       &[]string{"test"}[0],
		CanvasUuid:       nil,
		SceneItemEnabled: &[]bool{true}[0],
		SceneName:        &[]string{"Scene"}[0],
		SceneUuid:        nil,
		SourceName:       &[]string{"test"}[0],
		SourceUuid:       nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.DuplicateSceneItem(&sceneitems.DuplicateSceneItemParams{
		CanvasName:           &[]string{"test"}[0],
		CanvasUuid:           nil,
		DestinationSceneName: &[]string{"test"}[0],
		DestinationSceneUuid: nil,
		SceneItemId:          &[]int{1}[0],
		SceneName:            &[]string{"Scene"}[0],
		SceneUuid:            nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.SceneItems.GetGroupSceneItemList(&sceneitems.GetGroupSceneItemListParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		SceneName:  &[]string{"Scene"}[0],
		SceneUuid:  nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.SceneItems.GetSceneItemBlendMode(&sceneitems.GetSceneItemBlendModeParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneName:   &[]string{"Scene"}[0],
		SceneUuid:   nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemEnabled(&sceneitems.GetSceneItemEnabledParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneName:   &[]string{"Scene"}[0],
		SceneUuid:   nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemId(&sceneitems.GetSceneItemIdParams{
		CanvasName:   &[]string{"test"}[0],
		CanvasUuid:   nil,
		SceneName:    &[]string{"Scene"}[0],
		SceneUuid:    nil,
		SearchOffset: &[]float64{1.0}[0],
		SourceName:   &[]string{"test"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemIndex(&sceneitems.GetSceneItemIndexParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneName:   &[]string{"Scene"}[0],
		SceneUuid:   nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		SceneName:  &[]string{"Scene"}[0],
		SceneUuid:  nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemLocked(&sceneitems.GetSceneItemLockedParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneName:   &[]string{"Scene"}[0],
		SceneUuid:   nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemSource(&sceneitems.GetSceneItemSourceParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneName:   &[]string{"Scene"}[0],
		SceneUuid:   nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemTransform(&sceneitems.GetSceneItemTransformParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneName:   &[]string{"Scene"}[0],
		SceneUuid:   nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemBlendMode(&sceneitems.SetSceneItemBlendModeParams{
		CanvasName:         &[]string{"test"}[0],
		CanvasUuid:         nil,
		SceneItemBlendMode: &[]string{"OBS_BLEND_NORMAL"}[0],
		SceneItemId:        &[]int{1}[0],
		SceneName:          &[]string{"Scene"}[0],
		SceneUuid:          nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemEnabled(&sceneitems.SetSceneItemEnabledParams{
		CanvasName:       &[]string{"test"}[0],
		CanvasUuid:       nil,
		SceneItemEnabled: &[]bool{true}[0],
		SceneItemId:      &[]int{1}[0],
		SceneName:        &[]string{"Scene"}[0],
		SceneUuid:        nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemIndex(&sceneitems.SetSceneItemIndexParams{
		CanvasName:     &[]string{"test"}[0],
		CanvasUuid:     nil,
		SceneItemId:    &[]int{1}[0],
		SceneItemIndex: &[]int{1}[0],
		SceneName:      &[]string{"Scene"}[0],
		SceneUuid:      nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemLocked(&sceneitems.SetSceneItemLockedParams{
		CanvasName:      &[]string{"test"}[0],
		CanvasUuid:      nil,
		SceneItemId:     &[]int{1}[0],
		SceneItemLocked: &[]bool{true}[0],
		SceneName:       &[]string{"Scene"}[0],
		SceneUuid:       nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemTransform(&sceneitems.SetSceneItemTransformParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneItemTransform: &typedefs.SceneItemTransform{
			BoundsHeight: 1.0,
			BoundsType:   "OBS_BOUNDS_NONE",
			BoundsWidth:  1.0,
		},
		SceneName: &[]string{"Scene"}[0],
		SceneUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.SceneItems.RemoveSceneItem(&sceneitems.RemoveSceneItemParams{
		CanvasName:  &[]string{"test"}[0],
		CanvasUuid:  nil,
		SceneItemId: &[]int{1}[0],
		SceneName:   &[]string{"nonexistent"}[0],
		SceneUuid:   nil,
	})
	if err != nil {
		t.Logf("%s", err)
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

	_, err = client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: &[]string{"Scene"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Scenes.GetCurrentPreviewScene(&scenes.GetCurrentPreviewSceneParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetCurrentProgramScene(&scenes.GetCurrentProgramSceneParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetGroupList(&scenes.GetGroupListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetSceneList(&scenes.GetSceneListParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.GetSceneSceneTransitionOverride(&scenes.GetSceneSceneTransitionOverrideParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		SceneName:  &[]string{"Scene"}[0],
		SceneUuid:  nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.SetCurrentPreviewScene(&scenes.SetCurrentPreviewSceneParams{
		SceneName: &[]string{"Scene"}[0],
		SceneUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
		SceneName: &[]string{"Scene"}[0],
		SceneUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.SetSceneName(&scenes.SetSceneNameParams{
		CanvasName:   &[]string{"test"}[0],
		CanvasUuid:   nil,
		NewSceneName: &[]string{"Scene"}[0],
		SceneName:    &[]string{"Scene"}[0],
		SceneUuid:    nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Scenes.SetSceneSceneTransitionOverride(&scenes.SetSceneSceneTransitionOverrideParams{
		CanvasName:         &[]string{"test"}[0],
		CanvasUuid:         nil,
		SceneName:          &[]string{"Scene"}[0],
		SceneUuid:          nil,
		TransitionDuration: &[]float64{50.0}[0],
		TransitionName:     &[]string{"Cut"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Scenes.RemoveScene(&scenes.RemoveSceneParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		SceneName:  &[]string{"nonexistent"}[0],
		SceneUuid:  nil,
	})
	if err != nil {
		t.Logf("%s", err)
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

	_, err = client.Sources.GetSourceActive(&sources.GetSourceActiveParams{
		CanvasName: &[]string{"test"}[0],
		CanvasUuid: nil,
		SourceName: &[]string{"test"}[0],
		SourceUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Sources.GetSourceScreenshot(&sources.GetSourceScreenshotParams{
		CanvasName:              &[]string{"test"}[0],
		CanvasUuid:              nil,
		ImageCompressionQuality: &[]float64{8.0}[0],
		ImageFormat:             &[]string{"png"}[0],
		ImageHeight:             &[]float64{8.0}[0],
		ImageWidth:              &[]float64{8.0}[0],
		SourceName:              &[]string{"test"}[0],
		SourceUuid:              nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Sources.SaveSourceScreenshot(&sources.SaveSourceScreenshotParams{
		CanvasName:              &[]string{"test"}[0],
		CanvasUuid:              nil,
		ImageCompressionQuality: &[]float64{8.0}[0],
		ImageFilePath:           &[]string{"test"}[0],
		ImageFormat:             &[]string{"png"}[0],
		ImageHeight:             &[]float64{8.0}[0],
		ImageWidth:              &[]float64{8.0}[0],
		SourceName:              &[]string{"test"}[0],
		SourceUuid:              nil,
	})
	if err != nil {
		t.Logf("%s", err)
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
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Stream.SendStreamCaption(&stream.SendStreamCaptionParams{CaptionText: &[]string{"test"}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Stream.StartStream(&stream.StartStreamParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Stream.StopStream(&stream.StopStreamParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Stream.ToggleStream(&stream.ToggleStreamParams{})
	if err != nil {
		t.Logf("%s", err)
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
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.GetCurrentSceneTransition(&transitions.GetCurrentSceneTransitionParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.GetSceneTransitionList(&transitions.GetSceneTransitionListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.GetTransitionKindList(&transitions.GetTransitionKindListParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransitionDuration(
		&transitions.SetCurrentSceneTransitionDurationParams{TransitionDuration: &[]float64{50.0}[0]},
	)
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransition(
		&transitions.SetCurrentSceneTransitionParams{TransitionName: &[]string{"Cut"}[0]},
	)
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransitionSettings(&transitions.SetCurrentSceneTransitionSettingsParams{
		Overlay:            &[]bool{true}[0],
		TransitionSettings: map[string]interface{}{"test": "test"},
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Transitions.SetTBarPosition(&transitions.SetTBarPositionParams{
		Position: &[]float64{1.0}[0],
		Release:  &[]bool{true}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Transitions.TriggerStudioModeTransition(&transitions.TriggerStudioModeTransitionParams{})
	if err != nil {
		t.Logf("%s", err)
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
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.GetStudioModeEnabled(&ui.GetStudioModeEnabledParams{})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenInputFiltersDialog(&ui.OpenInputFiltersDialogParams{
		InputName: &[]string{"test"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenInputInteractDialog(&ui.OpenInputInteractDialogParams{
		InputName: &[]string{"test"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.Error(t, err)
	_, err = client.Ui.OpenInputPropertiesDialog(&ui.OpenInputPropertiesDialogParams{
		InputName: &[]string{"test"}[0],
		InputUuid: nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenSourceProjector(&ui.OpenSourceProjectorParams{
		CanvasName:        &[]string{"test"}[0],
		CanvasUuid:        nil,
		MonitorIndex:      &[]int{1}[0],
		ProjectorGeometry: nil,
		SourceName:        &[]string{"test"}[0],
		SourceUuid:        nil,
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.OpenVideoMixProjector(&ui.OpenVideoMixProjectorParams{
		MonitorIndex:      &[]int{1}[0],
		ProjectorGeometry: nil,
		VideoMixType:      &[]string{"OBS_WEBSOCKET_VIDEO_MIX_TYPE_PREVIEW"}[0],
	})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
	_, err = client.Ui.SetStudioModeEnabled(&ui.SetStudioModeEnabledParams{StudioModeEnabled: &[]bool{true}[0]})
	if err != nil {
		t.Logf("%s", err)
	}
	assert.NoError(t, err)
}
