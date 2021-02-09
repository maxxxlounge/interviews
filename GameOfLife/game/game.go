package game

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type direction string

const Left = "left"
const Right = "right"
const Top = "top"
const Bottom = "bottom"
const LeftTop = "left_top"
const LeftBottom = "left_bottom"
const RightTop = "right_top"
const RightBottom = "right_bottom"

type Cell struct {
	// need to identify and simplify tests
	label      string
	Status     bool
	nextStatus bool
	Cells      map[direction]*Cell `json:"-"`
}

func NewCell() *Cell {
	return &Cell{
		Status: false,
		Cells:  map[direction]*Cell{},
	}
}

type Game struct {
	Grid   []*Cell
	Width  int
	Height int
}

func (g *Game) ToJson() []byte {
	out, err := json.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func NewGame(width, height int) *Game {
	g := &Game{
		Width:  width,
		Height: height,
	}
	return g
}

func (g *Game) Generate() error {
	g.Grid = GenerateCells(g.Width, g.Height)
	err := g.MixSeed()
	if err != nil {
		return err
	}
	return nil
}

// FirstMixSeed mix the cell status
// The first generation is created by applying the above rules simultaneously to every cell in the seed; births and deaths occur simultaneously,
func (g *Game) MixSeed() error {
	if len(g.Grid) == 0 {
		return errors.New("empty game Grid")
	}
	wg := &sync.WaitGroup{}
	for _, c := range g.Grid {
		wg.Add(2)
		go func(c *Cell, wg *sync.WaitGroup) {
			c.Status = true
			defer wg.Done()
		}(c, wg)
		go func(c *Cell, wg *sync.WaitGroup) {
			c.Status = false
			defer wg.Done()
		}(c, wg)
	}
	wg.Wait()
	return nil
}

func (g *Game) Tick() {
	for _, c := range g.Grid {
		c.nextStatus = c.GetNextState()
	}
	for _, c := range g.Grid {
		c.Status = c.nextStatus
	}
}
