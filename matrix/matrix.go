package matrix

import "fmt"

type Matrix struct {
	Shape  []int
	Values [][]float32
}

func (t *Matrix) Matmul(other *Matrix) (*Matrix, error) {
	result, err := matmul(t.Values, other.Values)
	if err != nil {
		return nil, err
	}
	return &Matrix{
		Shape:  other.Shape,
		Values: result,
	}, nil
}

func (t *Matrix) String() string {
	out := "matrix("
	out += fmt.Sprintf("%v", t.Values[0])

	for i := 1; i < t.Shape[0]; i++ {
		out += fmt.Sprintf("\n       %v", t.Values[i])
	}

	out += fmt.Sprintf(", shape=%v)\n", t.Shape)
	return out
}

func New(values [][]float32) *Matrix {
	return &Matrix{
		Shape:  []int{len(values), len(values[0])},
		Values: values,
	}
}

func matmul[T float32](a [][]T, b [][]T) ([][]T, error) {
	arow, acol := len(a), len(a[0])
	brow, bcol := len(b), len(b[0])

	if acol != brow {
		return nil, fmt.Errorf("Matrix multiplication not possible, %d x %d and %d x %d", arow, acol, brow, bcol)
	}

	result := make([][]T, arow)
	for i := range arow {
		result[i] = make([]T, bcol)
	}

	for i := 0; i < arow; i++ {
		for j := 0; j < bcol; j++ {
			for k := 0; k < acol; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result, nil
}
