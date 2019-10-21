package govector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversion(t *testing.T) {
	assert := assert.New(t)

	v, err := AsVector([]uint8{0, 1, 2, 3, 4, 5})
	errMsg := "Error converting to vector type"
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]uint16{0, 1, 2, 3, 4, 5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]uint32{0, 1, 2, 3, 4, 5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]uint64{0, 1, 2, 3, 4, 5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]int8{0, 1, 2, 3, 4, 5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]int16{0, 1, 2, 3, 4, 5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]int32{0, 1, 2, 3, 4, 5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]int64{0, 1, 2, 3, 4, 5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1, 2, 3, 4, 5}, v, errMsg)

	v, err = AsVector([]float32{0, 1.1, 2.2, 3.3, 4.4, 5.5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1.100000023841858, 2.200000047683716, 3.299999952316284, 4.400000095367432, 5.5}, v, errMsg)

	v, err = AsVector([]float64{0, 1.1, 2.2, 3.3, 4.4, 5.5})
	assert.Nil(err, errMsg)
	assert.Equal(Vector{0, 1.1, 2.2, 3.3, 4.4, 5.5}, v, errMsg)

	_, err = AsVector(1)
	assert.Equal(ErrorCastToVector, err)
	assert.NotNil(err, "Integer should return error in vector conversion")
}
