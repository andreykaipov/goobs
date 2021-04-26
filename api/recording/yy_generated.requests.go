// This file has been automatically generated. Don't edit it.

package recording

type StartStopRecordingParams struct{}
type StartStopRecordingResponse struct{}
type StartRecordingParams struct{}
type StartRecordingResponse struct{}
type StopRecordingParams struct{}
type StopRecordingResponse struct{}
type SetRecordingFolderParams struct {
	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}
type SetRecordingFolderResponse struct{}
type GetRecordingFolderParams struct{}
type GetRecordingFolderResponse struct {
	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}
