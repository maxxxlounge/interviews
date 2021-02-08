package game

// GetNextState elaborate the next state for cell
// Any live cell with two or three neighbors survives.
// Any dead cell with three live neighbors becomes a live cell.
// All other live Cells die in the next generation.
func (c *Cell) GetNextState() bool {
	activeNearCells := 0
	for _, v := range c.Cells {
		if v.Status {
			activeNearCells++
		}
	}
	if activeNearCells == 2 {
		return c.Status
	}
	if activeNearCells == 3 {
		return true
	}

	return false
}
