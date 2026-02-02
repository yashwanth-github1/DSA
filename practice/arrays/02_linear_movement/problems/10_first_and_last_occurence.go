package main

func firstAndLastOccurunceinArray(a []int, target int) (int, int) {
	first, last := -1, -1
	for i, v := range a {
		if v == target {
			if first == -1 {
				first = i
			}
			last = i
		}
	}

	return first, last

}
