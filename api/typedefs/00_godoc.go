// Package typedefs defines both generated structs as specified under the
// [Typedefs
// section](https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#typedefs)
// of the protocol, and also manually defined "common" structs used across
// several events and requests.
//
// For example, `Crop` is used across:
//
//     request - GetSceneItemProperties
//     request - SetSceneItemProperties
//     typedef - SceneItemTransform
package typedefs
