package main

// Took
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/comments.json
// and put it into something like https://mholt.github.io/json-to-go/, deduping
// and cleaning up the structs manually.

// Param represents individual entry in a request's Params field, or in an event
// or request's Returns field.
type Param struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

// Event represents the specification of an event. Its fields are also common
// for a Request, and embedded in that definition.
type Event struct {
	// API can either be "events" or "requests"
	API string `json:"api"`

	// Category is the subsection of the Event or Request, e.g. "general" or
	// "sources". We don't really need this as we already know which
	// category an event or request belongs to since the parent key of this
	// event or request is also the category.
	Category string `json:"category"`

	// Categories is like Category because it's only ever of length one, but
	// I imagine it's a slice for future-proofing reasons.
	Categories []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"categories"`

	// Lead is a preface for the description. If it's present, we should use
	// it first.
	Lead string `json:"lead"`

	// Description is... the description.
	Description string `json:"description"`

	// Name is the name of the event or request.
	Name string `json:"name"`

	// Names is like Name because it's only ever of length one, but
	// I imagine it's a slice for future-proofing reasons.
	Names []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"names"`

	// Return is a neat representation of the fields in the JSON response of
	// the event or request, I guess for documentation purposes. It could be
	// either a `[]string` or just a `string`. For programmatic and parsing
	// purposes, Returns is easier to use though.
	Return interface{} `json:"return"`

	// Returns is a representation of the fields in the JSON response.
	Returns []*Param `json:"returns"`

	// The following are kinda unnecessary fields, but keep them here
	// because I already did the work

	// Since is the version the event or request was introduced in.
	Since string `json:"since"`

	// Sinces is like Since because it's only ever of length one, but
	// I imagine it's a slice for future-proofing reasons.
	Sinces []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"sinces"`

	// Examples looks to be always empty, but I guess it's for examples.
	Examples []interface{} `json:"examples"`

	// Heading is just for documentation generation on Palakis' end.
	Heading struct {
		Level int    `json:"level"`
		Text  string `json:"text"`
	} `json:"heading"`

	// Subheads looks to be always empty, but it's probably for
	// documentation purposes again.
	Subheads []interface{} `json:"subheads"`

	// Type looks to be always "class" for events and requests.
	Type string `json:"type"`
}

// Request surprisingly represents a Request.
type Request struct {
	*Event

	// Param is a neat representation of the fields in the JSON request,
	// I guess for documentation purposes. It could be either a `[]string`
	// or just a `string`. For programmatic and parsing purposes, Params is
	// easier to use though.
	Param interface{} `json:"param,omitempty"`

	// Params is a representation of the fields in the JSON request.
	Params []*Param `json:"params"`

	// Deprecated is deprecation text for the request. If it's present, we
	// know the request is deprecated. We'd have to parse the string to find
	// out what exact version it was deprecated though.
	Deprecated string `json:"deprecated,omitempty"`

	// Deprecateds is like deprecated. Probably for future-proofing, but
	// I can't imagine why since it's not like you can have two different
	// deprecation versions. Who cares -- I don't need it.
	Deprecateds []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"deprecateds,omitempty"`
}

// TypeDef defines more complex types
type TypeDef struct {
	Examples []interface{} `json:"examples"`
	Heading  struct {
		Level int    `json:"level"`
		Text  string `json:"text"`
	} `json:"heading"`
	Name       string `json:"name"`
	Properties []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"properties"`
	Property interface{}   `json:"property"` // could be []string or just string; not so important since Properties has all the info we'd need
	Subheads []interface{} `json:"subheads"`
	Typedef  string        `json:"typedef"`
	Typedefs []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"typedefs"`
}

// Comments represents the generated comments.json from Palakis
type Comments struct {
	Events   map[string][]*Event   `json:"events"`
	Requests map[string][]*Request `json:"requests"`
	Typedefs []*TypeDef            `json:"typedefs"`
}
