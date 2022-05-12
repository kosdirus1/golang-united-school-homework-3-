package homework

func reverse(input []int64) (result []int64) {
	inputLen := len(input)
	resSlice := make([]int64, inputLen)
	for i, v := range input {
		ii := inputLen - 1 - i
		resSlice[ii] = v
	}
	return resSlice
}
