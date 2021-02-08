package game

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCells(t *testing.T) {
	tests := map[string]struct {
		width  int
		height int
	}{
		"emptyGrid": {
			width:  0,
			height: 0,
		},
		"3x3Grid": {
			width:  3,
			height: 3,
		},
		"4x4Grid": {
			width:  4,
			height: 4,
		},
		"6x6Grid": {
			width:  6,
			height: 6,
		},
		"10x10Grid": {
			width:  10,
			height: 10,
		},
	}

	for name, tc := range tests {
		fmt.Println(name)
		t.Run(name, func(t *testing.T) {
			grid := GenerateCells(tc.width, tc.height)
			assert.Equal(t, tc.width*tc.height, len(grid))
			if len(grid) == 0{
				return
			}

			for hi:=0;hi<tc.height;hi++ {
				for ci := 0; ci < tc.width; ci++ {
					index := (hi*tc.width)+ci

					if hi>1 && hi < tc.height-1 && ci >1 && ci < tc.width-1 {
						assert.NotNil(t,grid[index].Cells[LeftTop])
						assert.NotNil(t,grid[index].Cells[LeftTop])
						assert.NotNil(t,grid[index].Cells[RightTop])
						assert.NotNil(t,grid[index].Cells[RightBottom])
						assert.NotNil(t,grid[index].Cells[RightBottom])
						assert.NotNil(t,grid[index].Cells[Bottom])
						assert.NotNil(t,grid[index].Cells[LeftBottom])
					}

					if hi == 0{
						assert.Nil(t, grid[index].Cells[Top])
						assert.NotNil(t, grid[index].Cells[Bottom])
					}
					// last line
					if hi == tc.height-1 {
						assert.Nil(t, grid[index].Cells[Bottom])
						assert.NotNil(t, grid[index].Cells[Top])
					}

					// first row
					if ci == 0 {
						/* Check for first ROW */
						assert.Nil(t, grid[index].Cells[Left])
						assert.NotNil(t, grid[index].Cells[Right])
					}

					// last row
					if ci == tc.width-1 {
						assert.Nil(t, grid[index].Cells[Right])
						assert.NotNil(t, grid[index].Cells[Left])
					}

					/*
					assert.Nil(t, Grid[index].Cells[Top])
					assert.Nil(t, Grid[index].Cells[LeftTop])
					assert.Nil(t, Grid[index].Cells[RightTop])
					assert.Equal(t, Grid[index].Cells[Left], Grid[index-1])
					assert.NotNil(t, Grid[index].Cells[Bottom])
					assert.NotNil(t, Grid[index].Cells[Right])
					assert.NotNil(t, Grid[index].Cells[RightBottom])

					assert.Equal(t, Grid[index-1].Cells[Right], Grid[index])*/

				}
			}
			// Check for other cases: row 2 start from index

		})
	}

}
