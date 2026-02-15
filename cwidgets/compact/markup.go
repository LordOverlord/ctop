package compact

import ui "github.com/gizak/termui"

func formatToken(s string, fg ui.Attribute, bg ui.Attribute) (string, int) {
	// Si no quieres markup, regresa plano
	if fg == 0 && bg == 0 {
		return s, len(s)
	}
	return "[" + s + "](fg-" + attrName(fg) + ",bg-" + attrName(bg) + ")", len(s)
}
