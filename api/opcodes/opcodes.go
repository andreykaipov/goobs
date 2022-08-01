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

type Opcode interface{ id() int }

// Wrap wraps an opcode around an enclosing protocol Message
func Wrap(o Opcode) *Message { return &Message{Op: o.id(), D: marshal(o)} }

func ID(o Opcode) int              { return o.id() }
func (o *Hello) id() int           { return 0 }
func (o *Identify) id() int        { return 1 }
func (o *Identified) id() int      { return 2 }
func (o *Reidentify) id() int      { return 3 }
func (o *Event) id() int           { return 5 }
func (o *Request) id() int         { return 6 }
func (o *RequestResponse) id() int { return 7 }

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

// GetType does what it says.

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

// batch opcodes, no thx
// opcode 8
// opcode 9

// Wrap takes one of the above opcode structs and wrap it within a *Message, so
// that we don't have to specify the appropriate opcode number for each opcode
// struct. It then returns the marshalled message.
//
// Twice-encoding the data might look like a bit of a red flag, but it's
// actually not a problem at all because our wrapper Message's D field uses
// json.RawMessage instead of just raw []bytes. The former type is idempotent
// with respect to marshalling! I didn't believe it either!
//
//func Wrap(data interface{}) json.RawMessage {
//	op := -1
//
//	switch data.(type) {
//	case *Identify:
//		op = 1
//	case *Request:
//		op = 6
//	default:
//		panic(fmt.Errorf("Can't figure out opcode for given type %T: %#v", data, data))
//	}
//
//	return marshal(&Message{Op: op, D: marshal(data)})
//}

//func GetTypeForOpcode(o int) {
//	switch name {
//	case "CurrentSceneCollectionChanging":
//	}
//}

//func Unwrap(msg *Message) (json.RawMessage, error) {
//	switch msg.Op {
//	case 0:
//		opcode := &Hello{}
//		if err := json.Unmarshal(msg.D, opcode); err != nil {
//			return opcode, nil
//		}
//	}
//
//	switch data.(type) {
//	case *Identify:
//		op = 1
//	case *Request:
//		op = 6
//	default:
//		panic(fmt.Errorf("Can't figure out opcode for given type %T: %#v", data, data))
//	}
//
//	return marshal(&Message{Op: op, D: marshal(data)})
//}

// Should be safe ignoring any marshalling errors, since the only things we're
// marshalling are already well-typed, or have likely already been unmarshalled
// (e.g. we receive a message from server, process it, send back).
//
// Also see https://stackoverflow.com/q/33903552/4085283.
//
func marshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("failed marshalling %#v: %w", v, err))
	}

	return b
}

func (o *Message) String() string {
	return string(marshal(o))
}
