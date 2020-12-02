package main

var _ = declareDay(1, day01)

func day01(partB bool, inputUntyped interface{}) interface{} {
	input := inputUntyped.([]int)
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if partB {
				for k := j + 1; k < len(input); k++ {
					if input[i]+input[j]+input[k] == 2020 {
						return input[i] * input[j] * input[k]
					}
				}
			} else {
				if input[i]+input[j] == 2020 {
					return input[i] * input[j]
				}
			}
		}
	}
	panic("no solution")
}
