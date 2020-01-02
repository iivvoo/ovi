package basicemu

import (
	"gitlab.com/iivvoo/ovim/ovim"
)

// Move moves the cursor in the specified direction
func Move(b *ovim.Buffer, c *ovim.Cursor, movement ovim.CursorDirection) {
	switch movement {
	case ovim.CursorUp:
		if c.Line > 0 {
			c.Line--
			if c.Pos > len(b.Lines[c.Line]) {
				c.Pos = len(b.Lines[c.Line])
			}
		}
	case ovim.CursorDown:
		// weirdness because empty last line that we want to position on
		if c.Line < len(b.Lines)-1 {
			c.Line++
			if c.Pos > len(b.Lines[c.Line]) {
				c.Pos = len(b.Lines[c.Line])
			}
		}
	case ovim.CursorLeft:
		if c.Pos > 0 {
			c.Pos--
		} else if c.Line > 0 {
			c.Line--
			c.Pos = len(b.Lines[c.Line])
		}
	case ovim.CursorRight:
		if c.Pos < len(b.Lines[c.Line]) {
			c.Pos++
		} else if c.Line < len(b.Lines)-1 {
			c.Line++
			c.Pos = 0
		}
	case ovim.CursorBegin:
		c.Pos = 0
	case ovim.CursorEnd:
		// move *past* the end
		c.Pos = len(b.Lines[c.Line])
	}
}
