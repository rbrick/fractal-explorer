package engine

import "gonum.org/v1/gonum/mat"

var (
	identityMatrix2x2 = []float64{
		1.0, 0.0,
		0.0, 1.0,
	}

	identityMatrix3x3 = []float64{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
	}

	identityMatrix4x4 = []float64{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
)

//IdentityMatrix2x2
func IdentityMatrix2x2() []float64 {
	return identityMatrix2x2
}

func IdentityMatrix3x3() []float64 {
	return identityMatrix3x3
}

func IdentityMatrix4x4() []float64 {
	return identityMatrix4x4
}

func MatrixToArray(m mat.Matrix, array []float32) {
	r, c := m.Dims()

	for i := 0; i < r; i++ {
		row := mat.Row(nil, i, m)
		for j := 0; j < c; j++ {
			array[i+j*c] = float32(row[i])
		}
	}
}
