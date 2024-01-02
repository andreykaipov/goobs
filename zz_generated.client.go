// This file has been automatically generated. Don't edit it.

package goobs

import (
	api "github.com/andreykaipov/goobs/api"
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
	c.Config = (*config.Client)(&api.Service{Client: c.client})
	c.Filters = (*filters.Client)(&api.Service{Client: c.client})
	c.General = (*general.Client)(&api.Service{Client: c.client})
	c.Inputs = (*inputs.Client)(&api.Service{Client: c.client})
	c.MediaInputs = (*mediainputs.Client)(&api.Service{Client: c.client})
	c.Outputs = (*outputs.Client)(&api.Service{Client: c.client})
	c.Record = (*record.Client)(&api.Service{Client: c.client})
	c.SceneItems = (*sceneitems.Client)(&api.Service{Client: c.client})
	c.Scenes = (*scenes.Client)(&api.Service{Client: c.client})
	c.Sources = (*sources.Client)(&api.Service{Client: c.client})
	c.Stream = (*stream.Client)(&api.Service{Client: c.client})
	c.Transitions = (*transitions.Client)(&api.Service{Client: c.client})
	c.Ui = (*ui.Client)(&api.Service{Client: c.client})
}
