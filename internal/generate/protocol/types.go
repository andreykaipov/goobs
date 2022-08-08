package main

// Protocol represents the generated protocol.json available at
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.json
//
// Put that into https://mholt.github.io/json-to-go/ and we're all good. v5's
// protocol is super clean for us already!
type Protocol struct {
	Enums    []*Enum    `json:"enums"`
	Requests []*Request `json:"requests"`
	Events   []*Event   `json:"events"`
}

type Enum struct {
	EnumType        string            `json:"enumType"`
	EnumIdentifiers []*EnumIdentifier `json:"enumIdentifiers"`
}

type Event struct {
	Description       string       `json:"description"`
	EventType         string       `json:"eventType"`
	EventSubscription string       `json:"eventSubscription"`
	Complexity        int          `json:"complexity"`
	RPCVersion        string       `json:"rpcVersion"`
	Deprecated        bool         `json:"deprecated"`
	InitialVersion    string       `json:"initialVersion"`
	Category          string       `json:"category"`
	DataFields        []*DataField `json:"dataFields"`
}

type Request struct {
	Description    string           `json:"description"`
	RequestType    string           `json:"requestType"`
	Complexity     int              `json:"complexity"`
	RPCVersion     string           `json:"rpcVersion"`
	Deprecated     bool             `json:"deprecated"`
	InitialVersion string           `json:"initialVersion"`
	Category       string           `json:"category"`
	RequestFields  []*RequestField  `json:"requestFields"`
	ResponseFields []*ResponseField `json:"responseFields"`
}

type EnumIdentifier struct {
	Description    string `json:"description"`
	EnumIdentifier string `json:"enumIdentifier"`
	Deprecated     bool   `json:"deprecated"`
	InitialVersion string `json:"initialVersion"`

	// this is not well defined either ... :/
	// it's an int in ObsMediaInputAction but strings everywhere elseRequestBatchExecuionTypet
	RPCVersion interface{} `json:"rpcVersion"`

	// because of some ints showing up as "1<<2" should the protocol
	// generate a well defined field here? file a bug upstream?
	EnumValue interface{} `json:"enumValue"`
}

type RequestField struct {
	FieldCommon
	ValueRestrictions     interface{} `json:"valueRestrictions"`
	ValueOptional         bool        `json:"valueOptional"`
	ValueOptionalBehavior interface{} `json:"valueOptionalBehavior"`
}

type ResponseField struct {
	FieldCommon
}

type DataField struct {
	FieldCommon
}

type Field interface {
	GetValueName() string
	GetValueType() string
	GetValueDescription() string
}

type FieldCommon struct {
	ValueName        string `json:"valueName"`
	ValueType        string `json:"valueType"`
	ValueDescription string `json:"valueDescription"`
}

func (o *FieldCommon) GetValueName() string {
	return o.ValueName
}
func (o *FieldCommon) GetValueType() string {
	return o.ValueType
}
func (o *FieldCommon) GetValueDescription() string {
	return o.ValueDescription
}
