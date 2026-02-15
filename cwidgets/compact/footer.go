package compact

import (
	"strings"

	ui "github.com/gizak/termui"
)

type CompactFooter struct {
	X, Y     int
	Width    int
	Height   int
	par      *ui.Par
	keybinds []Keybinding
}

func NewCompactFooter() *CompactFooter {
	f := &CompactFooter{
		X:        rowPadding,
		Height:   1,
		keybinds: DefaultKeybindings(),
	}

	p := ui.NewPar(" ")
	p.Height = f.Height
	p.Border = false

	p.TextBgColor = ui.ColorWhite
	p.TextFgColor = ui.ColorBlack

	f.par = p
	return f
}

func (f *CompactFooter) GetHeight() int { return f.Height }

func (f *CompactFooter) SetWidths(totalWidth int, _ []int) {
	f.par.SetX(f.X)
	f.par.SetWidth(totalWidth)
	f.Width = totalWidth

	f.par.Text = f.buildFooterText(totalWidth)

}

func (f *CompactFooter) SetX(x int)       { f.X = x; f.par.SetX(x) }
func (f *CompactFooter) SetY(y int)       { f.Y = y; f.par.SetY(y) }
func (f *CompactFooter) SetText(s string) { f.par.Text = s }

func (f *CompactFooter) Buffer() ui.Buffer {
	buf := ui.NewBuffer()
	buf.Merge(f.par.Buffer())
	return buf
}

func (f *CompactFooter) SetKeybindings(kb []Keybinding) { f.keybinds = kb }

func (f *CompactFooter) buildFooterText(width int) string {
	if width <= 0 {
		return ""
	}

	sep := " | "

	out := ""
	plainLen := 0

	for i, kb := range f.keybinds {
		keyToken, keyPlain := formatToken(kb.Key, kb.KeyFg, kb.KeyBg)
		actToken, actPlain := formatToken(kb.Action, kb.ActionFg, kb.ActionBg)

		partToken := keyToken + " " + actToken
		partPlainLen := keyPlain + 1 + actPlain

		if i == 0 {
			if partPlainLen > width {
				break
			}
			out = partToken
			plainLen = partPlainLen
			continue
		}

		if plainLen+len(sep)+partPlainLen > width {
			break
		}

		out += sep + partToken
		plainLen += len(sep) + partPlainLen
	}

	if plainLen < width {
		out += strings.Repeat(" ", width-plainLen)
	}

	return out
}

func attrName(a ui.Attribute) string {
	switch a {
	case ui.ColorBlack:
		return "black"
	case ui.ColorRed:
		return "red"
	case ui.ColorGreen:
		return "green"
	case ui.ColorYellow:
		return "yellow"
	case ui.ColorBlue:
		return "blue"
	case ui.ColorMagenta:
		return "magenta"
	case ui.ColorCyan:
		return "cyan"
	case ui.ColorWhite:
		return "white"
	default:
		return "white"
	}
}
