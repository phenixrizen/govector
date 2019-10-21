package govector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimilarity(t *testing.T) {
	assert := assert.New(t)

	// equal
	v1 := Vector{1, 1, 1, 1, 0, 0, 0, 0, 0}
	v3 := Vector{0, 0, 1, 1, 1, 1, 0, 0}
	e := Equal(v1, v3)
	assert.False(e, "Error vectors are not equal")

	// cosine similarity of two vectors
	v1 = Vector{1, 1, 1, 1, 0, 0, 0, 0, 0}
	v2 := Vector{0, 0, 1, 1, 1, 1, 0, 0, 0}
	_sim := 2.0
	sim, err := CosineSimilarity(v1, v2)
	assert.Nil(err)
	assert.Equal(_sim, sim)
	_, err = CosineSimilarity(v1, v3)
	assert.Equal(ErrorVectorLengths, err)

	// eucliudian distance of two vectors
	v1 = Vector{1, 1, 1, 1, 0, 0, 0, 0, 0}
	v2 = Vector{0, 0, 1, 1, 1, 1, 0, 0, 0}
	_dis := 2.0
	dis, err := EuclideanDistance(v1, v2)
	assert.Nil(err)
	assert.Equal(_dis, dis)
	_, err = EuclideanDistance(v1, v3)
	assert.Equal(ErrorVectorLengths, err)

	// manhatten distance of two vectors
	v1 = Vector{1, 1, 1, 1, 0, 0, 0, 0, 0}
	v2 = Vector{0, 0, 1, 1, 1, 1, 0, 0, 0}
	_dis = 4.0
	dis, err = ManhattenDistance(v1, v2)
	assert.Nil(err)
	assert.Equal(_dis, dis)
	_, err = ManhattenDistance(v1, v3)
	assert.Equal(ErrorVectorLengths, err)

	// minkowski distance of two vectors
	v1 = Vector{1, 1, 1, 1, 0, 0, 0, 0, 0}
	v2 = Vector{0, 0, 1, 1, 1, 1, 0, 0, 0}
	_dis = 2.0
	dis, err = MinowskiDistance(v1, v2, 2)
	assert.Nil(err)
	assert.Equal(_dis, dis)
	_, err = MinowskiDistance(v1, v3, 2)
	assert.Equal(ErrorVectorLengths, err)

	// nth root of a float
	_root := 1.772004514666935
	root := NthRoot(3.14, 2)
	assert.Equal(_root, root)

	// probability
	prob, err := Probability(v1, v2)
	_prob := 0.5
	assert.Nil(err, "Error calculating probability")
	assert.Equal(_prob, prob)
	prob, err = Probability(v1, v1)
	_prob = 1.0
	assert.Nil(err, "Error calculating probability")
	assert.Equal(_prob, prob)
	_, err = Probability(v1, v3)
	assert.Equal(ErrorVectorLengths, err)

}
