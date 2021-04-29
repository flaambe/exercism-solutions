// Package grains provides methods for calculating the number of wheat grains on a chessboard.
package grains

import "errors"

// Square returns the number of wheat grains on the given square of the chessboard.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("square must be between 1 and 64")
	}

	return 1 << (square - 1), nil
}

// Total returns the total number of wheat grains on the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
