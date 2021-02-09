package game

import (
	"fmt"
)

// GenerateCells provide the cells generation using setup reference method for cell and manage special border's case
// return an array of cell reference
func GenerateCells(width, height int) []*Cell {
	var (
		prevCell *Cell
		// head and tail cell maintains a quick reference to start and ending cell
		firstCellOfTheRow *Cell
	)

	out := []*Cell{}
	for hi := 0; hi < height; hi++ {
		for wi := 0; wi < width; wi++ {

			c := NewCell()

			c.label = fmt.Sprintf("cell_w%v_h%v", wi, hi)
			out = append(out, c)

			if firstCellOfTheRow == nil {
				firstCellOfTheRow = c
				prevCell = c
				continue
			}

			if wi == 0 {
				c.Cells[Top] = firstCellOfTheRow
				firstCellOfTheRow.Cells[Bottom] = c
				firstCellOfTheRow = c
				prevCell = c
				continue
			}

			c.Cells[Left] = prevCell

			if prevCell != nil {
				c.Cells[Left].Cells[Right] = c
			}

			c.SetupReference()

			if wi == width-1 {
				prevCell = nil
				continue
			}

			prevCell = c

		}
	}
	return out
}
