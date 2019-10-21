package govector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectors(t *testing.T) {
	assert := assert.New(t)

	x, err := AsVector([]int{2, 2, 2, 4, 2, 5})
	_x := Vector{2, 2, 2, 4, 2, 5}
	errMsg := "Error casting integer array to vector"
	assert.Equal(_x, x, errMsg)
	assert.Nil(err, errMsg)

	w, err := AsVector([]float64{1.0, 1.0, 1.0, 1.0, 1.0, 4.0})
	_w := Vector{1.0, 1.0, 1.0, 1.0, 1.0, 4.0}
	errMsg = "Error casting float64 array to vector"
	assert.Equal(_w, w, errMsg)
	assert.Nil(err, errMsg)

	q, err := AsVector([]float64{0.05, 0.95})
	_q := Vector{0.05, 0.95}
	errMsg = "Error casting float64 array to vector"
	assert.Equal(_q, q, errMsg)
	assert.Nil(err, errMsg)

	dx := x.Diff()
	_dx := Vector{0, 0, 2, -2, 3}
	errMsg = "Error calculating diff"
	assert.Equal(_dx, dx, errMsg)
	dw := w.Diff()
	_dw := Vector{0, 0, 0, 0, 3}
	assert.Equal(_dw, dw, errMsg)
	e := Vector{}
	_e := e.Diff()
	assert.Equal(Vector{NA}, _e)

	rdx := x.RelDiff()
	_rdx := Vector{0, 0, 0.5, -1, 0.6}
	assert.Equal(_rdx, rdx)
	_e = e.RelDiff()
	assert.Equal(Vector{NA}, _e)

	max := x.Max()
	_max := 5.0
	assert.Equal(_max, max, "Error calculating max")

	min := x.Min()
	_min := 2.0
	assert.Equal(_min, min, "Error calculating min")

	min, max = x.MinMax()
	assert.Equal(_min, min, "Error calculating min/max")
	assert.Equal(_max, max, "Error calculating min/max")

	zz := Vector{0, -1.22}
	_ = zz.Min()
	_, _ = zz.MinMax()
	zz = Vector{0, 1}
	_ = zz.Max()

	empirical := x.Ecdf()
	_empirical := empirical(0.25)
	assert.Equal(0.0, _empirical, "Error calculating empirical")

	percentile := empirical(2.4)
	_percentile := 2.0 / 3.0
	assert.Equal(_percentile, percentile, "Error in CDF calculation")

	m, vr := Vector{1., 2., 3.}.MeanVar()
	_m, _vr := 2.0, 1.0
	assert.Equal(_m, m, "Incorrect mean calculation")
	assert.Equal(_vr, vr, "Incorrect variance calculation")

	f, err := x.weightedSum(q)
	assert.Equal(NA, f, "Weighted sum should fail")
	assert.Equal(ErrorVectorLengths, err, "Weighted sum should fail")

	f, err = dx.WeightedMean(dw)
	_f := 3.0
	errMsg = "Error calculating weighted mean"
	assert.Equal(_f, f)
	assert.Nil(err, errMsg)
	_, err = dx.WeightedMean(q)
	assert.Equal(ErrorVectorLengths, err, "Weighted Mean should fail")

	r := Vector{}
	f = r.variance(1.1)
	errMsg = "Error calculating variance"
	assert.Equal(0.0, f, errMsg)
	r = Vector{1.0}
	f = r.variance(1.1)
	assert.Equal(0.0, f, errMsg)

	v := x.Quantiles(q)
	_v := Vector{2, 4.699999999999999}
	assert.Equal(_v, v, "Error calculating quantiles")
	nn := Vector{}
	n := nn.Quantiles(x)
	assert.Equal(Vector{0, 0, 0, 0, 0, 0}, n)

	cumsum := x.Cumsum()
	assert.Equal(Vector{2, 4, 6, 10, 12, 17}, cumsum, "Error calculating cumulative sum")

	ranks := x.Rank()
	assert.Equal(Vector{0, 0, 0, 4, 0, 5}, ranks, "Error calculating ranks")

	order := x.Order()
	assert.Equal(Vector{0, 1, 2, 4, 3, 5}, order, "Error calculating order")

	shuffled := x.Shuffle()
	assert.Equal(x.Len(), shuffled.Len(), "Error shuffling vector")

	y, err := AsVector([]int{-2, 2, -1, 4, 2, 5})
	assert.Nil(err, "Error casting negative integer array to vector")

	abs := y.Abs()
	assert.Equal(Vector{2, 2, 1, 4, 2, 5}, abs, "Error finding absolute values")

	scaled := x.FeatureScale()
	_scaled := Vector{0, 0, 0, 0.6666666666666666, 0, 1}
	assert.Equal(_scaled, scaled, "Error in feature scaling")

	v = x.Apply(empirical)
	_v = Vector{0.6666666666666666, 0.6666666666666666, 0.6666666666666666, 0.8333333333333334, 0.6666666666666666, 1}
	assert.Equal(_v, v, "Error applying function to vector")

	l := x.Len()
	x.Push(50)
	assert.Equal(l+1, x.Len(), "Error appending value to vector")

	xw := Join(x, w)
	assert.Equal(x.Len()+w.Len(), xw.Len(), "Error joining vectors")

	filtered := xw.Filter(func(x float64) bool {
		if x < 10 {
			return false
		}
		return true
	})
	assert.Equal(12, len(filtered), "Error filtering vector")

	z, err := AsVector([]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18})
	assert.Nil(err, "Error casting to vector")

	smoothed := z.Smooth(0, 0)
	assert.Equal(z, smoothed)

	smoothed = z.Smooth(1, 1)
	expected := Vector{1, 2, 4, 6, 8, 10, 12, 14, 16, 17}
	assert.Equal(expected, smoothed, "Error smoothing vector")

	x.Sort()
	assert.Equal(Vector{2, 2, 2, 2, 4, 5, 50}, x, "Error sorting vector")
}

func TestFixedPush(t *testing.T) {
	assert := assert.New(t)

	arr := make([]float64, 3, 3)

	v := Vector(arr)
	v.PushFixed(5.0)
	v.PushFixed(25.0)
	v.PushFixed(125.0)
	assert.Equal(v[2], 125.0)
}
