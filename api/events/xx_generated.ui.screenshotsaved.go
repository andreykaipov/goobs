// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the ScreenshotSaved event.

A screenshot has been saved.

Note: Triggered for the screenshot feature available in `Settings -> Hotkeys -> Screenshot Output` ONLY.
Applications using `Get/SaveSourceScreenshot` should implement a `CustomEvent` if this kind of inter-client
communication is desired.
*/
type ScreenshotSaved struct {
	// Path of the saved image file
	SavedScreenshotPath string `json:"savedScreenshotPath,omitempty"`
}
