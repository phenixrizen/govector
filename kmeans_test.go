package govector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMeans(t *testing.T) {
	assert := assert.New(t)

	RandSeed = 0

	v := Vector{0.1, 2.3, 1.0}

	a := Vector{0.1, 2.3, 1.0, 5.0}
	b := Vector{1.1, 2.6, 4.0, 5.0}
	c := Vector{3.1, 2.3, 3.0, 5.0}
	d := Vector{0.6, 2.2, 5.0, 5.0}
	e := Vector{0.1, 5.3, 4.0, 5.0}
	f := Vector{0.1, 2.3, 4.0, 5.0}
	vects := Nodes{a, b, c, d, e, f}
	vvects := Nodes{a, b, c, d, e, f, v}

	nodes, err := Train(vects, 12, 10)
	assert.Equal(ErrorClusterLength, err, "Error kmeans training should fail")
	assert.Nil(nodes, "Error kmeans nodes should be nil")

	nodes, err = Train(vvects, 3, 10)
	assert.Equal(ErrorVectorLengths, err, "Error kmeans training should fail")
	assert.Nil(nodes, "Error kmeans nodes should be nil")

	nodes, err = Train(vects, 3, 10)
	_nodes := Nodes{
		Vector{1.6, 2.3, 2, 5},
		Vector{0.6000000000000001, 2.3666666666666667, 4.333333333333333, 5},
		Vector{0.1, 5.3, 4, 5},
	}
	assert.Nil(err, "Error did not find kmeans centroids")
	assert.Equal(_nodes, nodes, "Error calculating centroids")

	n := Nearest(a, nodes)
	assert.Equal(0, n, "Error finding nearest centroid")
	_n := Nearest(e, nodes)
	assert.Equal(2, _n, "Error finding nearest centroid")
}
