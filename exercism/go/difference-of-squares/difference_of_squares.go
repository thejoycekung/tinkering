package diffsquares

// SquareOfSum takes an integer n, sums 1 to n, and squares the sum.
func SquareOfSum(n int) (square int) {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares takes an integer and returns the sum of squares (1 to n).
func SumOfSquares(n int) (sum int) {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference finds the difference between the sum of squares and the square of a number's sum.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
