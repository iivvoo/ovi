package viemu

import "github.com/iivvoo/novi/novi"

// StartSelection initializes start/end with the current cursor position
func (em *Vi) StartSelection() {
	em.SelectionStart = *em.Editor.Cursors[0]
	em.Editor.Selection.Enable()
	em.UpdateSelection()
}

func (em *Vi) CancelSelection() {
	em.Selection = SelectionNone
	em.Editor.Selection.Disable()
}

// UpdateSelection updates the end of the selection
func (em *Vi) UpdateSelection() {
	if em.Selection != SelectionNone {
		// There's a difference between the emulation selection and the UI selection.
		// In order to properly switch between the different types of selection, we need
		// to properly preserve the actual start/end, and only set the desired selection
		// (line,block,fluid) on em.Selection XXX
		// we have start/emuStart on the selection for that.

		em.SelectionEnd = *em.Editor.Cursors[0]

		s, e := em.SelectionStart, em.SelectionEnd
		if e.Line < s.Line || (e.Line == s.Line && e.Pos < s.Pos) {
			// swap start, end
			s, e = e, s
		}

		switch em.Selection {
		case SelectionLines:
			s.Pos = 0
			e.Pos = em.Editor.Buffer.GetLine(e.Line).Len() - 1
			fallthrough
		case SelectionBlock:
			em.Editor.Selection.SetBlock(true)
			fallthrough
		case SelectionFluid:
			em.Editor.Selection.SetStart(s)
			em.Editor.Selection.SetEnd(e)
		}
		log.Printf("Selection %s", em.Editor.Selection.ToString())
	}
}

// HandleSelectionBlock handles the block select key
func (em *Vi) HandleSelectionBlock(ev novi.Event) bool {
	// check key, set seleciton mode appropriately. Set start/end cursor
	// based on selection. In cursor movement, update selection based on selectionmode
	em.Selection = SelectionBlock
	em.StartSelection()
	return true
}

// HandleSelectionFluid handles the fluid select key
func (em *Vi) HandleSelectionFluid(ev novi.Event) bool {
	// check key, set seleciton mode appropriately. Set start/end cursor
	// based on selection. In cursor movement, update selection based on selectionmode
	em.Selection = SelectionFluid
	em.StartSelection()
	return true
}

// HandleSelectionLines handles the block select key
func (em *Vi) HandleSelectionLines(ev novi.Event) bool {
	// check key, set seleciton mode appropriately. Set start/end cursor
	// based on selection. In cursor movement, update selection based on selectionmode
	em.Selection = SelectionLines
	em.StartSelection()
	return true
}