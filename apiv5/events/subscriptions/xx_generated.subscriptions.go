// This file has been automatically generated. Don't edit it.

package subscriptions

// Subcription value used to disable all events.
const None = 0

// Subscription value to receive events in the `General` category.
const General = (1 << 0)

// Subscription value to receive events in the `Config` category.
const Config = (1 << 1)

// Subscription value to receive events in the `Scenes` category.
const Scenes = (1 << 2)

// Subscription value to receive events in the `Inputs` category.
const Inputs = (1 << 3)

// Subscription value to receive events in the `Transitions` category.
const Transitions = (1 << 4)

// Subscription value to receive events in the `Filters` category.
const Filters = (1 << 5)

// Subscription value to receive events in the `Outputs` category.
const Outputs = (1 << 6)

// Subscription value to receive events in the `SceneItems` category.
const SceneItems = (1 << 7)

// Subscription value to receive events in the `MediaInputs` category.
const MediaInputs = (1 << 8)

// Subscription value to receive the `VendorEvent` event.
const Vendors = (1 << 9)

// Subscription value to receive events in the `Ui` category.
const Ui = (1 << 10)

// Helper to receive all non-high-volume events.
const All = (General | Config | Scenes | Inputs | Transitions | Filters | Outputs | SceneItems | MediaInputs | Vendors)

// Subscription value to receive the `InputVolumeMeters` high-volume event.
const InputVolumeMeters = (1 << 16)

// Subscription value to receive the `InputActiveStateChanged` high-volume event.
const InputActiveStateChanged = (1 << 17)

// Subscription value to receive the `InputShowStateChanged` high-volume event.
const InputShowStateChanged = (1 << 18)

// Subscription value to receive the `SceneItemTransformChanged` high-volume event.
const SceneItemTransformChanged = (1 << 19)
