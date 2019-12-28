package ovim

/*
 * Buffer contains the implementaion of the text being manipulated by the editor.
 * It consists of an array of variable length strings of runes that can be manpipulated
 * using one or more cursors which identify one or more positions within the buffer
 */

// https://github.com/golang/go/wiki/SliceTricks

// Line implements a sequence of Runes
type Line []rune

func (l Line) GetRunes(start, end int) []rune {
	if start > len(l) {
		return nil
	}
	if start > end {
		return nil
	}
	if end > len(l) {
		end = len(l)
	}
	return l[start:end]
}

type Buffer struct {
	Lines []Line
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) Length() int {
	return len(b.Lines)
}

// GetLines attempts to retun the lines between start/end
func (b *Buffer) GetLines(start, end int) []Line {
	if start > b.Length() {
		return nil
	}
	if start > end {
		return nil
	}
	if end > b.Length() {
		end = b.Length()
	}
	return b.Lines[start:end]

}

// AddLine adds a line to the bottom of the buffer
func (b *Buffer) AddLine(line Line) {
	b.Lines = append(b.Lines, line)
}

/* PutRuneAtCursor
 * Does not update cursors
 */
func (b *Buffer) PutRuneAtCursors(cs Cursors, r rune) {
	for _, c := range cs {
		line := b.Lines[c.Line]
		line = append(line[:c.Pos], append(Line{r}, line[c.Pos:]...)...)
		b.Lines[c.Line] = line
	}
}

func (b *Buffer) RemoveRuneBeforeCursors(cs Cursors) {
	// We can't really do all cursors at once. Perhaps let caller always loop?
	// optionally, Cursors.all(func() {})
	for _, c := range cs {
		if c.Pos > 0 {
			line := b.Lines[c.Line]
			line = append(line[:c.Pos-1], line[c.Pos:]...)
			b.Lines[c.Line] = line
		}
	}
}

/* SplitLines
 *
 * Split lines ae position of cursors.
 * This is tricky since it will create extra lines, which may affect cursors below
 */
func (b *Buffer) SplitLines(cs Cursors) {
	linesAdded := 0
	for _, c := range cs {
		line := b.Lines[c.Line+linesAdded]
		before, after := line[c.Pos:], line[:c.Pos]
		b.Lines = append(b.Lines[:c.Line+linesAdded],
			append([]Line{after, before}, b.Lines[c.Line+linesAdded+1:]...)...)
		linesAdded++
	}
}
