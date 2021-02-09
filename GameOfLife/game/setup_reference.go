package game

var ErrMsgNilCellAtLeft = "Nil cell at left, impossible to retrieve other references"
var ErrMsgNilCellAtTop = "Nil cell at top, impossible to retrieve other references"

func (c *Cell) SetupReference() {
	if c.Cells[Left] == nil {
		return
	}
	if c.Cells[Left].Cells[Top] == nil {
		return
	}

	// find left top - right bottom relation
	if c.Cells[Left].Cells[Top] != nil {
		c.Cells[LeftTop] = c.Cells[Left].Cells[Top]
		c.Cells[LeftTop].Cells[RightBottom] = c
	}

	// find top - bottom relation
	if c.Cells[Left].Cells[Top].Cells[Right] != nil {
		c.Cells[Top] = c.Cells[Left].Cells[Top].Cells[Right]
		c.Cells[Top].Cells[Bottom] = c

		if c.Cells[Left].Cells[Top].Cells[Right].Cells[Right] != nil {
			c.Cells[RightTop] = c.Cells[Left].Cells[Top].Cells[Right].Cells[Right]
			c.Cells[RightTop].Cells[LeftBottom] = c
		}
	}
	return
}
