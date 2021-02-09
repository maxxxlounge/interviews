package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCell_GetNextState(t *testing.T) {

	tests := map[string]struct {
		cell              *Cell
		expectedNextState bool
	}{
		"ZeroActiveCell": {
			expectedNextState: false,
			cell: &Cell{
				label:  "",
				Status: false,
				Cells: map[direction]*Cell{
					Left:        NewCell(),
					LeftTop:     NewCell(),
					Top:         NewCell(),
					RightTop:    NewCell(),
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
		"NilAroundCell": {
			expectedNextState: false,
			cell: &Cell{
				label:  "",
				Status: false,
				Cells:  map[direction]*Cell{},
			},
		},
		"OneActiveCell": {
			expectedNextState: false,
			cell: &Cell{
				label:  "",
				Status: false,
				Cells: map[direction]*Cell{
					Left:        {Status: true},
					LeftTop:     NewCell(),
					Top:         NewCell(),
					RightTop:    NewCell(),
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
		"TwoDeadActiveCell": {
			expectedNextState: false,
			cell: &Cell{
				label:  "",
				Status: false,
				Cells: map[direction]*Cell{
					Left:        {Status: true},
					LeftTop:     {Status: true},
					Top:         NewCell(),
					RightTop:    NewCell(),
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
		"TwoAliveActiveCell": {
			expectedNextState: true,
			cell: &Cell{
				label:  "",
				Status: true,
				Cells: map[direction]*Cell{
					Left:        {Status: true},
					LeftTop:     {Status: true},
					Top:         NewCell(),
					RightTop:    NewCell(),
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
		"ThreeDeadActiveCell": {
			expectedNextState: true,
			cell: &Cell{
				label:  "",
				Status: false,
				Cells: map[direction]*Cell{
					Left:        {Status: true},
					LeftTop:     {Status: true},
					Top:         {Status: true},
					RightTop:    NewCell(),
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
		"ThreeAliveActiveCell": {
			expectedNextState: true,
			cell: &Cell{
				label:  "",
				Status: true,
				Cells: map[direction]*Cell{
					Left:        {Status: true},
					LeftTop:     {Status: true},
					Top:         {Status: true},
					RightTop:    NewCell(),
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
		"FourAliveActiveCell": {
			expectedNextState: false,
			cell: &Cell{
				label:  "",
				Status: true,
				Cells: map[direction]*Cell{
					Left:        {Status: true},
					LeftTop:     {Status: true},
					Top:         {Status: true},
					RightTop:    {Status: true},
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
		"FourDeadActiveCell": {
			expectedNextState: false,
			cell: &Cell{
				label:  "",
				Status: false,
				Cells: map[direction]*Cell{
					Left:        {Status: true},
					LeftTop:     {Status: true},
					Top:         {Status: true},
					RightTop:    {Status: true},
					Right:       NewCell(),
					RightBottom: NewCell(),
					Bottom:      NewCell(),
					LeftBottom:  NewCell(),
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedNextState, tc.cell.GetNextState())
		})
	}
}
