package main

func Min(input []int) (result int) {

	result = input[0]

	for _, val := range input {
		if val < result {
			result = val
		}

	}

	return
}

func Max(input []int) (result int) {

	result = input[0]

	for _, val := range input {
		if val > result {
			result = val
		}

	}

	return
}

func main() {

}