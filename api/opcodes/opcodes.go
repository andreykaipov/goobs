package opcodes

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Op int             `json:"op"`
	D  json.RawMessage `json:"d"`
}

// opcode 0 (server -> client)
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-0
type Hello struct {
	ObsWebSocketVersion string         `json:"obsWebSocketVersion"`
	RPCVersion          int            `json:"rpcVersion"`
	Authentication      Authentication `json:"authentication"`
}
type Authentication struct {
	Challenge string `json:"challenge"`
	Salt      string `json:"salt"`
}

// opcode 1 (client -> server)
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-1
type Identify struct {
	RPCVersion         int    `json:"rpcVersion"`
	Authentication     string `json:"authentication"`
	EventSubscriptions int    `json:"eventSubscriptions"`
}

// opcode 2 (server -> client)
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-2
type Identified struct {
	NegotiatedRPCVersion int `json:"negotiatedRpcVersion"`
}

// opcode 3 (client -> server)
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-3
type Reidentify struct {
	EventSubscriptions int `json:"eventSubscriptions"`
}

// opcode 5 (server -> client)
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-5
type Event struct {
	EventType   string      `json:"eventType"`
	EventIntent int         `json:"eventIntent"`
	EventData   interface{} `json:"eventData"`
}

// opcode 6 (client -> server)
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-6
type Request struct {
	RequestType string      `json:"requestType"`
	RequestID   string      `json:"requestId"`
	RequestData interface{} `json:"requestData"`
}

// opcode 7 (server -> client)
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requestresponse-opcode-7
type RequestResponse struct {
	RequestType   string      `json:"requestType"`
	RequestID     string      `json:"requestId"`
	RequestStatus interface{} `json:"requestStatus"`
}

// opcode 8
// opcode 9

// should be pretty safe ignoring our marshalling errors since the only things
// we're marshalling have all either likely already been unmarshalled (server
// sends, we process, we send back), or everything we send is well-defined.
//
// Twice-encoding the data might look a bit of a red flag, but it's actually
// not a problem at all because the wrapping Message's D field uses
// json.RawMessage instead of just raw []bytes. The former is idempotent with
// respect to marshalling! I didn't believe it either!
//
func Wrap(data interface{}) json.RawMessage {
	op := -1

	switch data.(type) {
	case Identify:
		op = 1
	case *Identify:
		op = 1
	default:
		panic(fmt.Errorf("I don't know what to do with this data: %#v", data))
	}

	return marshal(&Message{Op: op, D: marshal(data)})
}

func marshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}
