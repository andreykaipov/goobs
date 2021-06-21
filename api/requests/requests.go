package requests

// Params describes the behavior for any params-like object. Used to abstract
// the functionality of any request that embeds ParamsBasic within their fields.
type Params interface {
	GetRequestType() string
	SetRequestType(string)
	GetMessageID() string
	SetMessageID(string)

	// The name of the actual request, i.e. "Abc" for "AbcParams". Used to
	// set the RequestType.
	GetSelfName() string
}

// ParamsBasic represents common parameters for any request.
type ParamsBasic struct {
	RequestType string `json:"request-type,omitempty"`
	MessageID   string `json:"message-id,omitempty"`
}

// GetRequestType does what it says.
func (o *ParamsBasic) GetRequestType() string {
	return o.RequestType
}

// SetRequestType does what it says.
func (o *ParamsBasic) SetRequestType(x string) {
	o.RequestType = x
}

// GetMessageID does what it says.
func (o *ParamsBasic) GetMessageID() string {
	return o.MessageID
}

// SetMessageID does what it says.
func (o *ParamsBasic) SetMessageID(x string) {
	o.MessageID = x
}

// Response describes the behavior for any response-like object. Used to
// abstract the functionality of any request's response that embeds
// ResponseBasic within their fields.
type Response interface {
	GetMessageID() string
	GetStatus() string
	GetError() string
}

// ResponseBasic represents common fields on any returned response.
type ResponseBasic struct {
	MessageID string `json:"message-id,omitempty"`
	Status    string `json:"status,omitempty"`
	Error     string `json:"error,omitempty"`
}

// GetMessageID does what it says.
func (o *ResponseBasic) GetMessageID() string {
	return o.MessageID
}

// GetStatus does what it says.
func (o *ResponseBasic) GetStatus() string {
	return o.Status
}

// GetError does what it says.
func (o *ResponseBasic) GetError() string {
	return o.Error
}
