package formingmagicsquare

import "math"

// FormingMagicSquare ...
func FormingMagicSquare(s [][]int32) int32 {
	possiblePermutations := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}

	minCost := float64(math.MaxInt32)
	for _, permutation := range possiblePermutations {
		var permutationCost float64
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				permutationCost += math.Abs(float64(s[i][j]) - float64(permutation[i][j]))
			}
		}
		minCost = math.Min(minCost, permutationCost)
	}

	return int32(minCost)
}
