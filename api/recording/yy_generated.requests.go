// This file has been automatically generated. Don't edit it.

package recording

type StartStopRecordingParams struct{}
type StartStopRecordingResponse struct{}
type StartRecordingParams struct{}
type StartRecordingResponse struct{}
type StopRecordingParams struct{}
type StopRecordingResponse struct{}
type SetRecordingFolderParams struct {
	RecFolder string `json:"rec-folder"`
}
type SetRecordingFolderResponse struct{}
type GetRecordingFolderParams struct{}
type GetRecordingFolderResponse struct {
	RecFolder string `json:"rec-folder"`
}
