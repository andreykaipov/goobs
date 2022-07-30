package requests

// Params describes the behavior for any params-like object. Used to abstract
// the functionality of any request that embeds ParamsBasic within their fields.
type Params interface {
	GetRequestType() string
	SetRequestType(string)

	// The name of the actual request, i.e. "Abc" for "AbcParams". Used to
	// set the RequestType, so it's essentially an alias for GetRequestType.
	GetSelfName() string
}

// ParamsBasic represents common parameters for any request.
type ParamsBasic struct {
	RequestType string `json:"request-type,omitempty"`
}

// GetRequestType does what it says.
func (o *ParamsBasic) GetRequestType() string {
	return o.RequestType
}

// SetRequestType does what it says.
func (o *ParamsBasic) SetRequestType(x string) {
	o.RequestType = x
}

// Response describes the behavior for any response-like object. Used to
// abstract the functionality of any request's response that embeds
// ResponseBasic within their fields.
type Response interface {
	GetMessageID() string
	GetStatus() RequestResponseStatus
}

// ResponseBasic represents common fields on any returned response.
type ResponseBasic struct {
	MessageID string                `json:"requestId"`
	Status    RequestResponseStatus `json:"requestStatus"`
}

type RequestResponseData interface{}

type RequestResponseStatus struct {
	Code    int    `json:"code"`
	Result  bool   `json:"result"`
	Comment string `json:"comment,omitempty"`
}

// GetMessageID does what it says.
func (o *ResponseBasic) GetMessageID() string {
	return o.MessageID
}

// GetStatus does what it says.
func (o *ResponseBasic) GetStatus() RequestResponseStatus {
	return o.Status
}
