package dp

// this function utilises a matrix

// the idea here is that multiply the matrix M = {{1,1},{1,0}} to itself

// then we get the (n+1)th Fibonacci number

func FibonnaciWithMatrix(n int) int {

	var M = [][]int{{1, 1}, {1, 0}}

	var F = [][]int{{1}, {0}}

	for i := 2; i < n+1; i++ {
		F = multiplyMatrix(M, F)
	}
	return F[0][0]
}

func multiplyMatrix(A, B [][]int) [][]int {
	if len(A[0]) != len(B) {
		return nil
	}

	a := len(A[0])
	b := len(A)
	c := len(B[0])
	C := make([][]int, b)

	for i := range C {
		C[i] = make([]int, b)
	}
	for i := 0; i < b; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k < a; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C

}
