package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var s float32
	for _, v := range input {
		s += v
	}
	return s / (float32(len(input)))
}
