// This file has been automatically generated. Don't edit it.

package subscriptions

const (
	// Subcription value used to disable all events.
	None = 0

	// Subscription value to receive events in the `General` category.
	General = (1 << 0)

	// Subscription value to receive events in the `Config` category.
	Config = (1 << 1)

	// Subscription value to receive events in the `Scenes` category.
	Scenes = (1 << 2)

	// Subscription value to receive events in the `Inputs` category.
	Inputs = (1 << 3)

	// Subscription value to receive events in the `Transitions` category.
	Transitions = (1 << 4)

	// Subscription value to receive events in the `Filters` category.
	Filters = (1 << 5)

	// Subscription value to receive events in the `Outputs` category.
	Outputs = (1 << 6)

	// Subscription value to receive events in the `SceneItems` category.
	SceneItems = (1 << 7)

	// Subscription value to receive events in the `MediaInputs` category.
	MediaInputs = (1 << 8)

	// Subscription value to receive the `VendorEvent` event.
	Vendors = (1 << 9)

	// Subscription value to receive events in the `Ui` category.
	Ui = (1 << 10)

	// Helper to receive all non-high-volume events.
	All = (General | Config | Scenes | Inputs | Transitions | Filters | Outputs | SceneItems | MediaInputs | Vendors | Ui)

	// Subscription value to receive the `InputVolumeMeters` high-volume event.
	InputVolumeMeters = (1 << 16)

	// Subscription value to receive the `InputActiveStateChanged` high-volume event.
	InputActiveStateChanged = (1 << 17)

	// Subscription value to receive the `InputShowStateChanged` high-volume event.
	InputShowStateChanged = (1 << 18)

	// Subscription value to receive the `SceneItemTransformChanged` high-volume event.
	SceneItemTransformChanged = (1 << 19)
)
