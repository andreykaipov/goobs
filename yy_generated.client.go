// This file has been automatically generated. Don't edit it.

package goobs

import (
	general "github.com/andreykaipov/goobs/api/general"
	profiles "github.com/andreykaipov/goobs/api/profiles"
	recording "github.com/andreykaipov/goobs/api/recording"
	replaybuffer "github.com/andreykaipov/goobs/api/replay_buffer"
	scenecollections "github.com/andreykaipov/goobs/api/scene_collections"
	sceneitems "github.com/andreykaipov/goobs/api/scene_items"
	scenes "github.com/andreykaipov/goobs/api/scenes"
	sources "github.com/andreykaipov/goobs/api/sources"
	streaming "github.com/andreykaipov/goobs/api/streaming"
	studiomode "github.com/andreykaipov/goobs/api/studio_mode"
	transitions "github.com/andreykaipov/goobs/api/transitions"
)

type Subclients struct {
	Profiles         *profiles.Client
	ReplayBuffer     *replaybuffer.Client
	SceneCollections *scenecollections.Client
	SceneItems       *sceneitems.Client
	Scenes           *scenes.Client
	Sources          *sources.Client
	Transitions      *transitions.Client
	General          *general.Client
	Recording        *recording.Client
	Streaming        *streaming.Client
	StudioMode       *studiomode.Client
}

func setClients(c *Client) {
	c.Profiles = profiles.NewClient(profiles.WithConn(c.conn))
	c.ReplayBuffer = replaybuffer.NewClient(replaybuffer.WithConn(c.conn))
	c.SceneCollections = scenecollections.NewClient(scenecollections.WithConn(c.conn))
	c.SceneItems = sceneitems.NewClient(sceneitems.WithConn(c.conn))
	c.Scenes = scenes.NewClient(scenes.WithConn(c.conn))
	c.Sources = sources.NewClient(sources.WithConn(c.conn))
	c.Transitions = transitions.NewClient(transitions.WithConn(c.conn))
	c.General = general.NewClient(general.WithConn(c.conn))
	c.Recording = recording.NewClient(recording.WithConn(c.conn))
	c.Streaming = streaming.NewClient(streaming.WithConn(c.conn))
	c.StudioMode = studiomode.NewClient(studiomode.WithConn(c.conn))
}
