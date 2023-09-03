package main

func SwapVariables(a int, b int) (c int, d int) {
	a, b = b, a

	return a, b
}
