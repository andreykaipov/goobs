# Request Mapping

Assuming we've defined a `goobs` client as follows:

```go
client, err := goobs.New("localhost:4455", goobs.WithPassword("whatever"))
```

The following tables show how to make the appropriate `goobs` request for any given any [obs-websocket request](https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requests):

## config
| obs-websocket | goobs |
| --- | --- |
| CreateProfile | `client.Config.CreateProfile(...)` |
| CreateSceneCollection | `client.Config.CreateSceneCollection(...)` |
| GetPersistentData | `client.Config.GetPersistentData(...)` |
| GetProfileList | `client.Config.GetProfileList(...)` |
| GetProfileParameter | `client.Config.GetProfileParameter(...)` |
| GetRecordDirectory | `client.Config.GetRecordDirectory(...)` |
| GetSceneCollectionList | `client.Config.GetSceneCollectionList(...)` |
| GetStreamServiceSettings | `client.Config.GetStreamServiceSettings(...)` |
| GetVideoSettings | `client.Config.GetVideoSettings(...)` |
| RemoveProfile | `client.Config.RemoveProfile(...)` |
| SetCurrentProfile | `client.Config.SetCurrentProfile(...)` |
| SetCurrentSceneCollection | `client.Config.SetCurrentSceneCollection(...)` |
| SetPersistentData | `client.Config.SetPersistentData(...)` |
| SetProfileParameter | `client.Config.SetProfileParameter(...)` |
| SetRecordDirectory | `client.Config.SetRecordDirectory(...)` |
| SetStreamServiceSettings | `client.Config.SetStreamServiceSettings(...)` |
| SetVideoSettings | `client.Config.SetVideoSettings(...)` |

## filters
| obs-websocket | goobs |
| --- | --- |
| CreateSourceFilter | `client.Filters.CreateSourceFilter(...)` |
| GetSourceFilter | `client.Filters.GetSourceFilter(...)` |
| GetSourceFilterDefaultSettings | `client.Filters.GetSourceFilterDefaultSettings(...)` |
| GetSourceFilterKindList | `client.Filters.GetSourceFilterKindList(...)` |
| GetSourceFilterList | `client.Filters.GetSourceFilterList(...)` |
| RemoveSourceFilter | `client.Filters.RemoveSourceFilter(...)` |
| SetSourceFilterEnabled | `client.Filters.SetSourceFilterEnabled(...)` |
| SetSourceFilterIndex | `client.Filters.SetSourceFilterIndex(...)` |
| SetSourceFilterName | `client.Filters.SetSourceFilterName(...)` |
| SetSourceFilterSettings | `client.Filters.SetSourceFilterSettings(...)` |

## general
| obs-websocket | goobs |
| --- | --- |
| BroadcastCustomEvent | `client.General.BroadcastCustomEvent(...)` |
| CallVendorRequest | `client.General.CallVendorRequest(...)` |
| GetHotkeyList | `client.General.GetHotkeyList(...)` |
| GetStats | `client.General.GetStats(...)` |
| GetVersion | `client.General.GetVersion(...)` |
| Sleep | `client.General.Sleep(...)` |
| TriggerHotkeyByKeySequence | `client.General.TriggerHotkeyByKeySequence(...)` |
| TriggerHotkeyByName | `client.General.TriggerHotkeyByName(...)` |

## inputs
| obs-websocket | goobs |
| --- | --- |
| CreateInput | `client.Inputs.CreateInput(...)` |
| GetInputAudioBalance | `client.Inputs.GetInputAudioBalance(...)` |
| GetInputAudioMonitorType | `client.Inputs.GetInputAudioMonitorType(...)` |
| GetInputAudioSyncOffset | `client.Inputs.GetInputAudioSyncOffset(...)` |
| GetInputAudioTracks | `client.Inputs.GetInputAudioTracks(...)` |
| GetInputDefaultSettings | `client.Inputs.GetInputDefaultSettings(...)` |
| GetInputKindList | `client.Inputs.GetInputKindList(...)` |
| GetInputList | `client.Inputs.GetInputList(...)` |
| GetInputMute | `client.Inputs.GetInputMute(...)` |
| GetInputPropertiesListPropertyItems | `client.Inputs.GetInputPropertiesListPropertyItems(...)` |
| GetInputSettings | `client.Inputs.GetInputSettings(...)` |
| GetInputVolume | `client.Inputs.GetInputVolume(...)` |
| GetSpecialInputs | `client.Inputs.GetSpecialInputs(...)` |
| PressInputPropertiesButton | `client.Inputs.PressInputPropertiesButton(...)` |
| RemoveInput | `client.Inputs.RemoveInput(...)` |
| SetInputAudioBalance | `client.Inputs.SetInputAudioBalance(...)` |
| SetInputAudioMonitorType | `client.Inputs.SetInputAudioMonitorType(...)` |
| SetInputAudioSyncOffset | `client.Inputs.SetInputAudioSyncOffset(...)` |
| SetInputAudioTracks | `client.Inputs.SetInputAudioTracks(...)` |
| SetInputMute | `client.Inputs.SetInputMute(...)` |
| SetInputName | `client.Inputs.SetInputName(...)` |
| SetInputSettings | `client.Inputs.SetInputSettings(...)` |
| SetInputVolume | `client.Inputs.SetInputVolume(...)` |
| ToggleInputMute | `client.Inputs.ToggleInputMute(...)` |

## media inputs
| obs-websocket | goobs |
| --- | --- |
| GetMediaInputStatus | `client.MediaInputs.GetMediaInputStatus(...)` |
| OffsetMediaInputCursor | `client.MediaInputs.OffsetMediaInputCursor(...)` |
| SetMediaInputCursor | `client.MediaInputs.SetMediaInputCursor(...)` |
| TriggerMediaInputAction | `client.MediaInputs.TriggerMediaInputAction(...)` |

## outputs
| obs-websocket | goobs |
| --- | --- |
| GetLastReplayBufferReplay | `client.Outputs.GetLastReplayBufferReplay(...)` |
| GetOutputList | `client.Outputs.GetOutputList(...)` |
| GetOutputSettings | `client.Outputs.GetOutputSettings(...)` |
| GetOutputStatus | `client.Outputs.GetOutputStatus(...)` |
| GetReplayBufferStatus | `client.Outputs.GetReplayBufferStatus(...)` |
| GetVirtualCamStatus | `client.Outputs.GetVirtualCamStatus(...)` |
| SaveReplayBuffer | `client.Outputs.SaveReplayBuffer(...)` |
| SetOutputSettings | `client.Outputs.SetOutputSettings(...)` |
| StartOutput | `client.Outputs.StartOutput(...)` |
| StartReplayBuffer | `client.Outputs.StartReplayBuffer(...)` |
| StartVirtualCam | `client.Outputs.StartVirtualCam(...)` |
| StopOutput | `client.Outputs.StopOutput(...)` |
| StopReplayBuffer | `client.Outputs.StopReplayBuffer(...)` |
| StopVirtualCam | `client.Outputs.StopVirtualCam(...)` |
| ToggleOutput | `client.Outputs.ToggleOutput(...)` |
| ToggleReplayBuffer | `client.Outputs.ToggleReplayBuffer(...)` |
| ToggleVirtualCam | `client.Outputs.ToggleVirtualCam(...)` |

## record
| obs-websocket | goobs |
| --- | --- |
| CreateRecordChapter | `client.Record.CreateRecordChapter(...)` |
| GetRecordStatus | `client.Record.GetRecordStatus(...)` |
| PauseRecord | `client.Record.PauseRecord(...)` |
| ResumeRecord | `client.Record.ResumeRecord(...)` |
| SplitRecordFile | `client.Record.SplitRecordFile(...)` |
| StartRecord | `client.Record.StartRecord(...)` |
| StopRecord | `client.Record.StopRecord(...)` |
| ToggleRecord | `client.Record.ToggleRecord(...)` |
| ToggleRecordPause | `client.Record.ToggleRecordPause(...)` |

## scene items
| obs-websocket | goobs |
| --- | --- |
| CreateSceneItem | `client.SceneItems.CreateSceneItem(...)` |
| DuplicateSceneItem | `client.SceneItems.DuplicateSceneItem(...)` |
| GetGroupSceneItemList | `client.SceneItems.GetGroupSceneItemList(...)` |
| GetSceneItemBlendMode | `client.SceneItems.GetSceneItemBlendMode(...)` |
| GetSceneItemEnabled | `client.SceneItems.GetSceneItemEnabled(...)` |
| GetSceneItemId | `client.SceneItems.GetSceneItemId(...)` |
| GetSceneItemIndex | `client.SceneItems.GetSceneItemIndex(...)` |
| GetSceneItemList | `client.SceneItems.GetSceneItemList(...)` |
| GetSceneItemLocked | `client.SceneItems.GetSceneItemLocked(...)` |
| GetSceneItemSource | `client.SceneItems.GetSceneItemSource(...)` |
| GetSceneItemTransform | `client.SceneItems.GetSceneItemTransform(...)` |
| RemoveSceneItem | `client.SceneItems.RemoveSceneItem(...)` |
| SetSceneItemBlendMode | `client.SceneItems.SetSceneItemBlendMode(...)` |
| SetSceneItemEnabled | `client.SceneItems.SetSceneItemEnabled(...)` |
| SetSceneItemIndex | `client.SceneItems.SetSceneItemIndex(...)` |
| SetSceneItemLocked | `client.SceneItems.SetSceneItemLocked(...)` |
| SetSceneItemTransform | `client.SceneItems.SetSceneItemTransform(...)` |

## scenes
| obs-websocket | goobs |
| --- | --- |
| CreateScene | `client.Scenes.CreateScene(...)` |
| GetCurrentPreviewScene | `client.Scenes.GetCurrentPreviewScene(...)` |
| GetCurrentProgramScene | `client.Scenes.GetCurrentProgramScene(...)` |
| GetGroupList | `client.Scenes.GetGroupList(...)` |
| GetSceneList | `client.Scenes.GetSceneList(...)` |
| GetSceneSceneTransitionOverride | `client.Scenes.GetSceneSceneTransitionOverride(...)` |
| RemoveScene | `client.Scenes.RemoveScene(...)` |
| SetCurrentPreviewScene | `client.Scenes.SetCurrentPreviewScene(...)` |
| SetCurrentProgramScene | `client.Scenes.SetCurrentProgramScene(...)` |
| SetSceneName | `client.Scenes.SetSceneName(...)` |
| SetSceneSceneTransitionOverride | `client.Scenes.SetSceneSceneTransitionOverride(...)` |

## sources
| obs-websocket | goobs |
| --- | --- |
| GetSourceActive | `client.Sources.GetSourceActive(...)` |
| GetSourceScreenshot | `client.Sources.GetSourceScreenshot(...)` |
| SaveSourceScreenshot | `client.Sources.SaveSourceScreenshot(...)` |

## stream
| obs-websocket | goobs |
| --- | --- |
| GetStreamStatus | `client.Stream.GetStreamStatus(...)` |
| SendStreamCaption | `client.Stream.SendStreamCaption(...)` |
| StartStream | `client.Stream.StartStream(...)` |
| StopStream | `client.Stream.StopStream(...)` |
| ToggleStream | `client.Stream.ToggleStream(...)` |

## transitions
| obs-websocket | goobs |
| --- | --- |
| GetCurrentSceneTransition | `client.Transitions.GetCurrentSceneTransition(...)` |
| GetCurrentSceneTransitionCursor | `client.Transitions.GetCurrentSceneTransitionCursor(...)` |
| GetSceneTransitionList | `client.Transitions.GetSceneTransitionList(...)` |
| GetTransitionKindList | `client.Transitions.GetTransitionKindList(...)` |
| SetCurrentSceneTransition | `client.Transitions.SetCurrentSceneTransition(...)` |
| SetCurrentSceneTransitionDuration | `client.Transitions.SetCurrentSceneTransitionDuration(...)` |
| SetCurrentSceneTransitionSettings | `client.Transitions.SetCurrentSceneTransitionSettings(...)` |
| SetTBarPosition | `client.Transitions.SetTBarPosition(...)` |
| TriggerStudioModeTransition | `client.Transitions.TriggerStudioModeTransition(...)` |

## ui
| obs-websocket | goobs |
| --- | --- |
| GetMonitorList | `client.Ui.GetMonitorList(...)` |
| GetStudioModeEnabled | `client.Ui.GetStudioModeEnabled(...)` |
| OpenInputFiltersDialog | `client.Ui.OpenInputFiltersDialog(...)` |
| OpenInputInteractDialog | `client.Ui.OpenInputInteractDialog(...)` |
| OpenInputPropertiesDialog | `client.Ui.OpenInputPropertiesDialog(...)` |
| OpenSourceProjector | `client.Ui.OpenSourceProjector(...)` |
| OpenVideoMixProjector | `client.Ui.OpenVideoMixProjector(...)` |
| SetStudioModeEnabled | `client.Ui.SetStudioModeEnabled(...)` |

