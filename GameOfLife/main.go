package main

import (
	"encoding/json"
	"fmt"
	"github.com/maxxxlounge/interviews/GameOfLife/game"
	"log"
)


func main(){
	cells:=GenerateCells(10,10)
	out,err := json.Marshal(cells)
	if err != nil {
		log.Fatal(err)
	}

	for i,c := range cells{
		fmt.Printf("%v %v" + " " + c.Status)
	}

}

func GenerateCells(width, height int) ([]*game.Cell){
	var (
		cells []*game.Cell
		prevCell *game.Cell
	)
	for hi := 0; hi < height; hi++{
		for wi := 0; wi < width; wi++{
			c := game.NewCell()
			cells = append(cells,c)
			if wi % height == 0 {
				c.Cells[game.Left] = nil
				c.Cells[game.LeftTop] = nil
			}
			if prevCell != nil {
				prevCell.Cells[game.Right]=c
			}
			if i%height==0{
				c.Cells[game.Top] = nil
			}else{
				prevCell.Cells[game.Bottom]=c
			}

			if i%width == 0 && i&height ==0{
				c.Cells[game.LeftTop] = nil
			}else {
				prevCell.Cells[game.RightBottom] = c
			}
			prevCell = c
		}
	}


	for i := 0; i < totalCell; i++{
		c := game.NewCell()
		cells = append(cells,c)
		if i % width == 0 {
			c.Cells[game.Left] = nil
		}else{
			prevCell.Cells[game.Right]=c
		}
		if i%height==0{
			c.Cells[game.Top] = nil
		}else{
			prevCell.Cells[game.Bottom]=c
		}

		if i%width == 0 && i&height ==0{
			c.Cells[game.LeftTop] = nil
		}else {
			prevCell.Cells[game.RightBottom] = c
		}
		prevCell = c
	}
	return cells
}
