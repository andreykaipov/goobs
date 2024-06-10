// This file has been automatically generated. Don't edit it.

package record

// Represents the request body for the CreateRecordChapter request.
type CreateRecordChapterParams struct {
	// Name of the new chapter
	ChapterName *string `json:"chapterName,omitempty"`
}

func NewCreateRecordChapterParams() *CreateRecordChapterParams {
	return &CreateRecordChapterParams{}
}
func (o *CreateRecordChapterParams) WithChapterName(x string) *CreateRecordChapterParams {
	o.ChapterName = &x
	return o
}

// Returns the associated request.
func (o *CreateRecordChapterParams) GetRequestName() string {
	return "CreateRecordChapter"
}

// Represents the response body for the CreateRecordChapter request.
type CreateRecordChapterResponse struct {
	_response
}

/*
Adds a new chapter marker to the file currently being recorded.

Note: As of OBS 30.2.0, the only file format supporting this feature is Hybrid MP4.
*/
func (c *Client) CreateRecordChapter(paramss ...*CreateRecordChapterParams) (*CreateRecordChapterResponse, error) {
	if len(paramss) == 0 {
		paramss = []*CreateRecordChapterParams{{}}
	}
	params := paramss[0]
	data := &CreateRecordChapterResponse{}
	return data, c.client.SendRequest(params, data)
}
