package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var sl []int

	for i := range input {
		sl = append(sl, i)
	}
	sort.Ints(sl)
	for _, v := range sl {
		result = append(result, input[v])
	}

	return result
}
