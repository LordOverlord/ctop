package compact

import (
	ui "github.com/gizak/termui"
)

type CompactGrid struct {
	ui.GridBufferer
	header *CompactHeader
	footer *CompactFooter
	cols   []CompactCol // reference columns
	Rows   []RowBufferer
	X, Y   int
	Width  int
	Height int
	Offset int // starting row offset
}

func NewCompactGrid() *CompactGrid {
	cg := &CompactGrid{
		header: NewCompactHeader(),
		footer: NewCompactFooter(),
	}
	cg.rebuildHeader()
	return cg
}

func (cg *CompactGrid) Align() {
	y := cg.Y

	if cg.Offset >= len(cg.Rows) || cg.Offset < 0 {
		cg.Offset = 0
	}

	colWidths := cg.calcWidths()

	// 1) header arriba
	cg.header.SetY(y)
	cg.header.SetWidths(cg.Width, colWidths)
	y += cg.header.GetHeight()

	// 2) body rows
	for _, r := range cg.Rows[cg.Offset:] {
		// corta si ya no cabe (dejando espacio para footer)
		if y+r.GetHeight() > ui.TermHeight()-cg.footer.GetHeight() {
			break
		}
		r.SetY(y)
		r.SetWidths(cg.Width, colWidths)
		y += r.GetHeight()
	}

	// 3) footer fijo abajo
	cg.footer.SetY(ui.TermHeight() - cg.footer.GetHeight())
	cg.footer.SetWidths(cg.Width, nil)
}

func (cg *CompactGrid) Clear() {
	cg.Rows = []RowBufferer{}
	cg.rebuildHeader()
}

func (cg *CompactGrid) GetHeight() int { return len(cg.Rows) + cg.header.Height + cg.footer.Height }
func (cg *CompactGrid) SetX(x int)     { cg.X = x }
func (cg *CompactGrid) SetY(y int)     { cg.Y = y }
func (cg *CompactGrid) SetWidth(w int) { cg.Width = w }
func (cg *CompactGrid) MaxRows() int {
	return ui.TermHeight() - cg.header.Height - cg.footer.Height - cg.Y
}

// calculate and return per-column width
func (cg *CompactGrid) calcWidths() []int {
	var autoCols int
	width := cg.Width
	colWidths := make([]int, len(cg.cols))

	for n, w := range cg.cols {
		colWidths[n] = w.FixedWidth()
		width -= w.FixedWidth()
		if w.FixedWidth() == 0 {
			autoCols++
		}
	}

	spacing := colSpacing * len(cg.cols)
	autoWidth := (width - spacing) / autoCols
	for n, val := range colWidths {
		if val == 0 {
			colWidths[n] = autoWidth
		}
	}
	return colWidths
}
func (cg *CompactGrid) visibleRows() (rows []RowBufferer) {
	// El header ocupa arriba, arrancamos después de él
	y := cg.Y + cg.header.GetHeight()
	yLimit := ui.TermHeight() - cg.footer.GetHeight()

	for _, r := range cg.Rows[cg.Offset:] {
		if y+r.GetHeight() > yLimit {
			break
		}
		rows = append(rows, r)
		y += r.GetHeight()
	}
	return rows
}

func (cg *CompactGrid) pageRows() (rows []RowBufferer) {
	rows = append(rows, cg.header)
	rows = append(rows, cg.visibleRows()...)
	rows = append(rows, cg.footer)
	return rows
}

func (cg *CompactGrid) Buffer() ui.Buffer {
	buf := ui.NewBuffer()
	buf.Merge(cg.header.Buffer())
	for _, r := range cg.visibleRows() {
		buf.Merge(r.Buffer())
	}
	buf.Merge(cg.footer.Buffer())
	return buf
}

func (cg *CompactGrid) AddRows(rows ...RowBufferer) {
	cg.Rows = append(cg.Rows, rows...)
}

func (cg *CompactGrid) rebuildHeader() {
	cg.cols = newRowWidgets()
	cg.header.clearFieldPars()
	for _, col := range cg.cols {
		cg.header.addFieldPar(col.Header())
	}
}
