package elementary

import "testing"

func Test_quickFind(t *testing.T) {
	pairs := [][]int{
		{3, 4},
		{4, 9},
		{8, 0},
		{2, 3},
		{5, 6},
		{2, 9},
		{5, 9},
		{7, 3},
		{4, 8},
		{5, 6},
		{0, 2},
		{6, 1},
	}

	quickFind(pairs)
}
