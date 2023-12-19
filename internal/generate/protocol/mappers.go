package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

/*
The following mappers are for fields with "complex types" in the OBS websocket
protocol, i.e. those that are of Object or Array<Object>. We have to map these
to Go types to the best of our ability, and the following functions do just
that.

We can use the following variables to help us make our decisions:

- the origin being one of 'request', 'response', or 'event'
- the name of the request, response, or event
- the field (notably its name) being mapped

Sometimes the best we can do is just map[string]any though!
*/

var typedefs = goobs + "/api/typedefs"

func mapObject(origin, name string, field Field) *Statement {
	fvn := field.GetValueName()

	switch fvn {
	case "inputAudioTracks":
		return Op("*").Qual(typedefs, "InputAudioTracks")
	case "streamServiceSettings":
		return Op("*").Qual(typedefs, "StreamServiceSettings")
	case "sceneItemTransform":
		return Op("*").Qual(typedefs, "SceneItemTransform")
	case "keyModifiers":
		return Op("*").Qual(typedefs, "KeyModifiers")

	// settings are generic enough to merit the use of map[string]any
	// for individual inputs, filters, outputs, etc.
	case "outputSettings":
		fallthrough
	case "inputSettings", "defaultInputSettings":
		fallthrough
	case "filterSettings", "defaultFilterSettings":
		fallthrough
	case "transitionSettings":
		return Map(String()).Any()

	// from CallVendorRequest* and VendorEvent, CustomEvent
	// same as above, the data here can be anything
	case "requestData", "responseData", "eventData":
		return Map(String()).Any()

	default:
		fmt.Printf("!! unhandled Object type for field %s from %s:%s\n", fvn, origin, name)
		return Map(String()).Any()
	}
}

func mapArrayObject(origin, name string, field Field) *Statement {
	fvn := field.GetValueName()

	switch {
	case fvn == "filters":
		return Index().Op("*").Qual(typedefs, "Filter")
	case fvn == "inputs" && origin == "event" && name == "InputVolumeMeters":
		return Index().Op("*").Qual(typedefs, "InputVolumeMeter")
	case fvn == "inputs":
		return Index().Op("*").Qual(typedefs, "Input")
	case fvn == "outputs":
		return Index().Op("*").Qual(typedefs, "Output")
	case fvn == "sceneItems" && origin == "event" && name == "SceneItemListReindexed":
		return Index().Op("*").Qual(typedefs, "SceneItemBasic")
	case fvn == "sceneItems":
		return Index().Op("*").Qual(typedefs, "SceneItem")
	case fvn == "scenes":
		return Index().Op("*").Qual(typedefs, "Scene")
	case fvn == "propertyItems":
		return Index().Op("*").Qual(typedefs, "PropertyItem")
	case fvn == "transitions":
		return Index().Op("*").Qual(typedefs, "Transition")
	case fvn == "monitors":
		return Index().Op("*").Qual(typedefs, "Monitor")
	default:
		fmt.Printf("!! unhandled Array<Object> type for field %s from %s:%s\n", fvn, origin, name)
		return Index().Map(String()).Any()
	}
}
