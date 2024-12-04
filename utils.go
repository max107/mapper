package mapper

func Keys[K, V comparable](mapping map[K]V) []K {
	keys := make([]K, len(mapping))

	i := 0
	for k, _ := range mapping {
		keys[i] = k
		i++
	}

	return keys
}

func Values[K, V comparable](mapping map[K]V) []V {
	result := make([]V, len(mapping))

	i := 0
	for _, v := range mapping {
		result[i] = v
		i++
	}

	return result
}
