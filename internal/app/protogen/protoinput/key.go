package protoinput

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	SelectFile key.Binding
}

var defaultKeyMap = KeyMap{
	SelectFile: key.NewBinding(key.WithKeys("enter")),
}
