package opcodes

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Op int             `json:"op"`
	D  json.RawMessage `json:"d"`
}

func (o *Message) Bytes() json.RawMessage { return marshal(o) }
func (o *Message) String() string         { return string(o.Bytes()) }

type Opcode interface{ id() int }

// Wrap wraps an opcode around an enclosing protocol Message
func Wrap(o Opcode) *Message { return &Message{Op: o.id(), D: marshal(o)} }

func (o *Hello) id() int           { return 0 }
func (o *Identify) id() int        { return 1 }
func (o *Identified) id() int      { return 2 }
func (o *Reidentify) id() int      { return 3 }
func (o *Event) id() int           { return 5 }
func (o *Request) id() int         { return 6 }
func (o *RequestResponse) id() int { return 7 }

func GetOpcodeForOp(code int) Opcode {
	switch code {
	case 0:
		return &Hello{}
	case 1:
		return &Identify{}
	case 2:
		return &Identified{}
	case 3:
		return &Reidentify{}
	case 4:
		// noop
	case 5:
		return &Event{}
	case 6:
		return &Request{}
	case 7:
		return &RequestResponse{}
	case 8:
		// request batch
	case 9:
		// request batch response
	}

	return nil
}

// Should be safe ignoring any marshalling errors, since the only things we're
// marshalling are already well-typed, or have likely already been unmarshalled
// (e.g. we receive a message from server, process it, send back).
//
// Also see https://stackoverflow.com/q/33903552/4085283.
//
func marshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("marshalling %#v: %w", v, err))
	}

	return b
}

// server -> client
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#hello-opcode-0
type Hello struct {
	ObsWebSocketVersion string         `json:"obsWebSocketVersion"`
	RPCVersion          int            `json:"rpcVersion"`
	Authentication      Authentication `json:"authentication"`
}
type Authentication struct {
	Challenge string `json:"challenge"`
	Salt      string `json:"salt"`
}

// client -> server
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#identify-opcode-1
type Identify struct {
	RPCVersion         int    `json:"rpcVersion"`
	Authentication     string `json:"authentication"`
	EventSubscriptions int    `json:"eventSubscriptions"`
}

// server -> client
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#identified-opcode-2
type Identified struct {
	NegotiatedRPCVersion int `json:"negotiatedRpcVersion"`
}

// client -> server
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#reidentify-opcode-3
type Reidentify struct {
	EventSubscriptions int `json:"eventSubscriptions"`
}

// server -> client
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#event-opcode-5
type Event struct {
	EventType   string          `json:"eventType"`
	EventIntent int             `json:"eventIntent"`
	EventData   json.RawMessage `json:"eventData"`
}

// client -> server
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#request-opcode-6
type Request struct {
	RequestType string      `json:"requestType"`
	RequestID   string      `json:"requestId"`
	RequestData interface{} `json:"requestData,omitempty"`
}

// server -> client
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-7
type RequestResponse struct {
	RequestType   string                `json:"requestType"`
	RequestID     string                `json:"requestId"`
	RequestStatus RequestResponseStatus `json:"requestStatus"`
	ResponseData  json.RawMessage       `json:"responseData,omitempty"`
}
type RequestResponseStatus struct {
	Code    int    `json:"code"`
	Result  bool   `json:"result"`
	Comment string `json:"comment,omitempty"`
}
