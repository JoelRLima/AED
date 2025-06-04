package sorting

import (
	//"math/rand"
)
//in place
func SelectionSort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		menorindice := i

		for j := i + 1; j < len(v); j++ {
			if v[j] < v[menorindice] {
				menorindice = j
			}
		}
		v[i], v[menorindice] = v[menorindice], v[i]
	}
}

