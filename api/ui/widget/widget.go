// Package widget defines the UI widgets within the Fyne toolkit
package widget

import "github.com/fyne-io/fyne/api/ui"

// Widget defines the standard behaviours of any widget. This extends the
// CanvasObject - a widget behaves in the same basic way but will encapsulate
// many child objects to create the rendered widget.
type Widget interface {
	CurrentSize() ui.Size
	Resize(ui.Size)
	CurrentPosition() ui.Position
	Move(ui.Position)

	MinSize() ui.Size

	// TODO should this move to a widget impl?... (private)
	Layout(ui.Size) []ui.CanvasObject
	ApplyTheme()
}

// A base widget class to define the standard widget behaviours.
type baseWidget struct {
	Size     ui.Size
	Position ui.Position

	objects []ui.CanvasObject
}

// Get the current size of this widget.
func (w *baseWidget) CurrentSize() ui.Size {
	return w.Size
}

// Set a new size for a widget.
// Note this hould not be used if the widget is being managed by a Layout within a Container.
func (w *baseWidget) Resize(size ui.Size) {
	w.Size = size
}

// Get the current position of this widget, relative to it's parent.
func (w *baseWidget) CurrentPosition() ui.Position {
	return w.Position
}

// Move the widget to a new position, relative to it's parent.
// Note this hould not be used if the widget is being managed by a Layout within a Container.
func (w *baseWidget) Move(pos ui.Position) {
	w.Position = pos
}

// FocussableWidget describes any widget that can respond to being focussed.
// It will receive the OnFocusGained and OnFocusLost events appropriately and,
// when focussed, it will also have OnKeyDown called as keys are pressed.
type FocussableWidget interface {
	OnFocusGained()
	OnFocusLost()
	OnKeyDown(*ui.KeyEvent)
}