package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	res := make([]int64, len(input))
	for i, v := range input {
		res[len(res)-1-i] = v
	}
	copy(result, res)
	return
}
