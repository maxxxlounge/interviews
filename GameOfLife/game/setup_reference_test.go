package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCell_SetupReference(t *testing.T) {
	tests := map[string]struct {
		errMsg              string
		inputAroundCells    map[direction]*Cell
		expectedAroundCells map[direction]*Cell
		hasError            bool
	}{
		"NilLeftCell": {
			inputAroundCells: map[direction]*Cell{
				Left: nil,
			},
			errMsg:   ErrMsgNilCellAtLeft,
			hasError: true,
		},
		"OnlyLeftNoTop": {
			inputAroundCells: map[direction]*Cell{
				Left: {
					label:  "OnlyLeftNoTop_left",
					Status: false,
					Cells:  map[direction]*Cell{},
				},
			},
			expectedAroundCells: map[direction]*Cell{
				Left: {
					label:  "OnlyLeftNoTop_left",
					Status: false,
					Cells:  map[direction]*Cell{},
				},
			},
			errMsg:   "",
			hasError: false,
		},
		"OnlyLeftAndTop": {
			inputAroundCells: map[direction]*Cell{
				Left: {
					label:  "OnlyLeftAndTop_Left",
					Status: false,
					Cells: map[direction]*Cell{
						Top: {
							label:  "OnlyLeftAndTop_LeftTop",
							Status: false,
							Cells:  map[direction]*Cell{},
						},
					},
				},
			},
			expectedAroundCells: map[direction]*Cell{
				Left: {
					Status: false,
					label:  "OnlyLeftAndTop_Left",
				},
				LeftTop: {
					Status: false,
					label:  "OnlyLeftAndTop_LeftTop",
				},
			},
			errMsg:   "",
			hasError: false,
		},
		"HasLeftTopRight": {
			inputAroundCells: map[direction]*Cell{
				Left: {
					label:  "HasLeftTopRight_Left",
					Status: false,
					Cells: map[direction]*Cell{
						Top: {
							label:  "HasLeftTopRight_LeftTop",
							Status: false,
							Cells: map[direction]*Cell{
								Right: {
									label:  "HasLeftTopRight_LeftTopRight",
									Status: false,
									Cells:  map[direction]*Cell{},
								},
							},
						},
					},
				},
			},
			expectedAroundCells: map[direction]*Cell{
				Left: {
					Status: false,
					label:  "HasLeftTopRight_Left",
				},
				LeftTop: {
					Status: false,
					label:  "HasLeftTopRight_LeftTop",
				},
				Top: {
					Status: false,
					label:  "HasLeftTopRight_LeftTopRight",
				},
			},
			errMsg:   "",
			hasError: false,
		},
		"HasLeftTopRightRight": {
			inputAroundCells: map[direction]*Cell{
				Left: {
					label:  "HasLeftTopRightRight_Left",
					Status: false,
					Cells: map[direction]*Cell{
						Top: {
							label:  "HasLeftTopRightRight_LeftTop",
							Status: false,
							Cells: map[direction]*Cell{
								Right: {
									label:  "HasLeftTopRightRight_LeftTopRight",
									Status: false,
									Cells: map[direction]*Cell{
										Right: {
											label:  "HasLeftTopRightRight_LeftTopRightRight",
											Status: false,
											Cells:  map[direction]*Cell{},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedAroundCells: map[direction]*Cell{
				Left: {
					Status: false,
					label:  "HasLeftTopRightRight_Left",
				},
				LeftTop: {
					Status: false,
					label:  "HasLeftTopRightRight_LeftTop",
				},
				Top: {
					Status: false,
					label:  "HasLeftTopRightRight_LeftTopRight",
				},
				RightTop: {
					Status: false,
					label:  "HasLeftTopRightRight_LeftTopRightRight",
				},
			},
			errMsg:   "",
			hasError: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCell()
			c.label = "mainCell"
			c.Cells[Left] = tc.inputAroundCells[Left]
			if tc.inputAroundCells[Left] != nil {
				tc.inputAroundCells[Left].Cells[Right] = c
			}
			require.NotNil(t, c)
			c.SetupReference()

			if tc.expectedAroundCells[Left] == nil {
				assert.Nil(t, c.Cells[Left], tc.expectedAroundCells[Left])
				return
			}
			require.NotNil(t, c.Cells[Left])

			if tc.expectedAroundCells[Left] != nil {
				assert.Equal(t, tc.expectedAroundCells[Left].label, c.Cells[Left].label)
				assert.Equal(t, c.label, c.Cells[Left].Cells[Right].label)

			}
			if tc.expectedAroundCells[LeftTop] != nil {
				assert.Equal(t, tc.expectedAroundCells[LeftTop].label, c.Cells[LeftTop].label)
				assert.Equal(t, c.label, c.Cells[LeftTop].Cells[RightBottom].label)
			}
			if tc.expectedAroundCells[Top] != nil {
				assert.Equal(t, tc.expectedAroundCells[Top].label, c.Cells[Top].label)
				assert.Equal(t, c.label, c.Cells[Top].Cells[Bottom].label)
			}
			if tc.expectedAroundCells[RightTop] != nil {
				assert.Equal(t, tc.expectedAroundCells[RightTop].label, c.Cells[RightTop].label)
				assert.Equal(t, c.label, c.Cells[RightTop].Cells[LeftBottom].label)
			}
		})
	}
}
