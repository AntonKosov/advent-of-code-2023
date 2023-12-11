package aoc

func GCD(nums ...int) int {
	if len(nums) < 2 {
		panic("incorrect number of parameters")
	}

	a, b := nums[0], nums[1]
	if len(nums) > 2 {
		b = GCD(nums[1:]...)
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(nums ...int) int {
	if len(nums) < 2 {
		panic("incorrect number of parameters")
	}

	a, b := nums[0], nums[1]
	if len(nums) > 2 {
		b = LCM(nums[1:]...)
	}

	return a * b / GCD(a, b)
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
