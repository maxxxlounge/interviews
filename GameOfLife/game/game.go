package game

type direction string

const Left  = "left"
const Right  = "right"
const Top  = "top"
const Bottom  = "bottom"
const LeftTop  = "left_top"
const LeftBottom  = "left_bottom"
const RightTop  = "right_top"
const RightBottom  = "right_bottom"


type Cell struct {
	Xi int
	Yi int
	Status bool
	Cells map[direction]*Cell
}

func NewCell()*Cell{
	return &Cell{
		Status: false,
		Cells: map[direction]*Cell{},
	}
}

func (c *Cell) GetNextState()bool{
	activeNearCells := 0;

	for _,v := range c.Cells{
		if v.Status {
			activeNearCells++
			if activeNearCells >= 3 {
				return true
			}
		}
	}

	return false
}


