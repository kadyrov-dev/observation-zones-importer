package coordinatesmatrix

import (
	"math"

	"github.com/kadyrov-dev/observation-zones-importer/internal/types"
)

func Build() []types.Coordinates {
	// Differences between map squares horizontally and vertically
	const (
		SquareWidth  = 0.02210140228
		SquareHeight = 0.00825229714263287
	)

	const (
		MinLeftX  = 37.12226483781682
		MinLeftY  = 55.46991536980882
		MaxRightX = 37.904366240096456
		MaxRightY = 56.018167666951456
	)

	horizontalSquaresCount := int(math.Round((MaxRightX - MinLeftX) / SquareWidth))
	verticalSquaresCount := int(math.Round((MaxRightY - MinLeftY) / SquareHeight))

	// Top-left square's coordinates
	// If zone will be expanded, then we need to update this values to actual
	leftX := 37.14226483781682
	leftY := 56.00991536980882
	rightX := 37.164366240096456
	rightY := 56.018167666951456

	prevX := leftX
	nextX := rightX
	prevY := leftY
	nextY := rightY

	var coords []types.Coordinates

	for vs := 0; vs < verticalSquaresCount; vs++ {
		for hs := 0; hs < horizontalSquaresCount; hs++ {
			coords = append(
				coords,
				types.Coordinates{
					LeftX:  prevX,
					LeftY:  prevY,
					RightX: nextX,
					RightY: nextY,
				},
			)

			prevX = nextX
			nextX += SquareWidth
		}

		nextX = rightX
		prevX = leftX
		nextY = prevY
		prevY -= SquareHeight
	}

	return coords
}
