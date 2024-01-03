// This file has been automatically generated. Don't edit it.

package goobs

import (
	config "github.com/andreykaipov/goobs/api/requests/config"
	filters "github.com/andreykaipov/goobs/api/requests/filters"
	general "github.com/andreykaipov/goobs/api/requests/general"
	inputs "github.com/andreykaipov/goobs/api/requests/inputs"
	mediainputs "github.com/andreykaipov/goobs/api/requests/mediainputs"
	outputs "github.com/andreykaipov/goobs/api/requests/outputs"
	record "github.com/andreykaipov/goobs/api/requests/record"
	sceneitems "github.com/andreykaipov/goobs/api/requests/sceneitems"
	scenes "github.com/andreykaipov/goobs/api/requests/scenes"
	sources "github.com/andreykaipov/goobs/api/requests/sources"
	stream "github.com/andreykaipov/goobs/api/requests/stream"
	transitions "github.com/andreykaipov/goobs/api/requests/transitions"
	ui "github.com/andreykaipov/goobs/api/requests/ui"
)

type Categories struct {
	Config      *config.Client
	Filters     *filters.Client
	General     *general.Client
	Inputs      *inputs.Client
	MediaInputs *mediainputs.Client
	Outputs     *outputs.Client
	Record      *record.Client
	SceneItems  *sceneitems.Client
	Scenes      *scenes.Client
	Sources     *sources.Client
	Stream      *stream.Client
	Transitions *transitions.Client
	Ui          *ui.Client
}

func setClients(c *Client) {
	c.Config = config.NewClient(c.client)
	c.Filters = filters.NewClient(c.client)
	c.General = general.NewClient(c.client)
	c.Inputs = inputs.NewClient(c.client)
	c.MediaInputs = mediainputs.NewClient(c.client)
	c.Outputs = outputs.NewClient(c.client)
	c.Record = record.NewClient(c.client)
	c.SceneItems = sceneitems.NewClient(c.client)
	c.Scenes = scenes.NewClient(c.client)
	c.Sources = sources.NewClient(c.client)
	c.Stream = stream.NewClient(c.client)
	c.Transitions = transitions.NewClient(c.client)
	c.Ui = ui.NewClient(c.client)
}
