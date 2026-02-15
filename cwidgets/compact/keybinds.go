package compact

import ui "github.com/gizak/termui"

type Keybinding struct {
	Key    string
	Action string

	// estilos para la tecla
	KeyFg ui.Attribute
	KeyBg ui.Attribute

	// estilos para la acción (opcional)
	ActionFg ui.Attribute
	ActionBg ui.Attribute
}

func DefaultKeybindings() []Keybinding {
	// Fondo del footer (debe coincidir con tu footer Par TextBgColor)
	bg := ui.ColorWhite

	normalKeyFg := ui.ColorBlue
	normalActionFg := ui.ColorBlack

	return []Keybinding{
		{Key: "<ENTER>", Action: "Menu", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "a", Action: "All", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "f", Action: "Filter", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "H", Action: "Header", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "h", Action: "Help", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "s", Action: "Sort", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "r", Action: "Reverse", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "o", Action: "Single", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "l", Action: "Logs", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "e", Action: "Exec", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "c", Action: "Columns", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
		{Key: "S", Action: "Save", KeyFg: normalKeyFg, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},

		// 🔴 Quit en rojo (solo la tecla)
		{Key: "q", Action: "Quit", KeyFg: ui.ColorRed, KeyBg: bg, ActionFg: normalActionFg, ActionBg: bg},
	}
}
