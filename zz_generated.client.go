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

type subclients struct {
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
	c.Config = &config.Client{Client: c.Client}
	c.Filters = &filters.Client{Client: c.Client}
	c.General = &general.Client{Client: c.Client}
	c.Inputs = &inputs.Client{Client: c.Client}
	c.MediaInputs = &mediainputs.Client{Client: c.Client}
	c.Outputs = &outputs.Client{Client: c.Client}
	c.Record = &record.Client{Client: c.Client}
	c.SceneItems = &sceneitems.Client{Client: c.Client}
	c.Scenes = &scenes.Client{Client: c.Client}
	c.Sources = &sources.Client{Client: c.Client}
	c.Stream = &stream.Client{Client: c.Client}
	c.Transitions = &transitions.Client{Client: c.Client}
	c.Ui = &ui.Client{Client: c.Client}
}
