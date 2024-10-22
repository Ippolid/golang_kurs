package storage

func Contains[T comparable](needle T, haystack []T) bool {
	k := 0
	for _, i := range haystack {
		if i == needle {
			k++
			break
		}
	}

	return k == 0
}
