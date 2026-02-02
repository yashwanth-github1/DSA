package main

func reverse(a []int, i, j int) {
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func rotateRight(a []int, k int) {
	n := len(a)
	if n == 0 {
		return
	}

	k = k % n

	reverse(a, 0, n-1)
	reverse(a, 0, k-1)
	reverse(a, k, n-1)
}

func rotateLeft(a []int, k int) {
	n := len(a)
	if n == 0 {
		return
	}

	k = k % n

	reverse(a, 0, k-1)
	reverse(a, k, n-1)
	reverse(a, 0, n-1)
}
