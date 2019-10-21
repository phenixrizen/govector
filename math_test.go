package govector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert := assert.New(t)

	x, err := AsVector([]int{1, 2, 3, 4, 6, 5})
	_x := Vector{1, 2, 3, 4, 6, 5}
	errMsg := "Error casting integer array to vector"
	assert.Equal(x, _x, errMsg)
	assert.Nil(err, errMsg)

	y, err := AsVector([]float64{2.6, 1.1, 3, 4.2, 5, 6})
	_y := Vector{2.6, 1.1, 3, 4.2, 5, 6}
	errMsg = "Error casting integer array to vector"
	assert.Equal(y, _y, errMsg)
	assert.Nil(err, errMsg)

	z, _ := AsVector([]int{1, 2})

	v, err := Product(x, y)
	_v := Vector{2.6, 2.2, 9, 16.8, 30, 30}
	errMsg = "Error calculating vector product"
	assert.Equal(_v, v, errMsg)
	assert.Nil(err, errMsg)
	_, err = Product(x, z)
	assert.Equal(ErrorVectorLengths, err)

	f, err := DotProduct(x, y)
	_f := 90.6
	errMsg = "Error calculating dot product"
	assert.Equal(_f, f, errMsg)
	assert.Nil(err, errMsg)
	_, err = DotProduct(x, z)
	assert.Equal(ErrorVectorLengths, err)

	f, err = Cosine(x, y)
	_f = 0.9713054876814755
	errMsg = "Error calculating cosine similarity"
	assert.Equal(_f, f, errMsg)
	assert.Nil(err, errMsg)
	_, err = Cosine(x, z)
	assert.Equal(ErrorVectorLengths, err)

	f, err = Correlation(x, y)
	_f = 0.8422701479042108
	errMsg = "Error calculating vector correlation"
	assert.Equal(_f, f, errMsg)
	assert.Nil(err, errMsg)
	_, err = Correlation(x, z)
	assert.Equal(ErrorVectorLengths, err)

	v, err = Average(x, y)
	_v = Vector{1.8, 1.55, 3, 4.1, 5.5, 5.5}
	assert.Equal(_v, v)
	assert.Nil(err, "Error calculating vector average")
	_, err = Average(x, z)
	assert.Equal(ErrorVectorLengths, err)

}
