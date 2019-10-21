package govector

import (
	"math"
)

// Equal determines if two nodes have the same values.
func Equal(a, b Vector) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// CosineSimilarity calculates the cosine similarity of two vectors
func CosineSimilarity(a, b Vector) (float64, error) {
	var sum float64
	if len(a) != len(b) {
		return sum, ErrorVectorLengths
	}
	var x, y int
	for n := len(a); n > 0; n-- {
		sum += a[x] * b[y]
		x++
		y++
	}
	return sum, nil
}

// EuclideanDistance calculates the euclidian distance between two vectors
func EuclideanDistance(a, b Vector) (float64, error) {
	var sum float64
	if len(a) != len(b) {
		return sum, ErrorVectorLengths
	}
	for i := 0; i < len(a); i++ {
		sum += math.Pow((a[i] - b[i]), 2.0)
	}
	return math.Sqrt(sum), nil
}

// ManhattenDistance calculates the Manhatten distance between two vectors
func ManhattenDistance(a, b Vector) (float64, error) {
	var sum float64
	if len(a) != len(b) {
		return sum, ErrorVectorLengths
	}
	for i := 0; i < len(a); i++ {
		sum += math.Abs(a[i] - b[i])
	}
	return sum, nil
}

// MinowskiDistance calculates the Minowski distance between two vectors
func MinowskiDistance(a, b Vector, pow int) (float64, error) {
	var sum float64
	if len(a) != len(b) {
		return sum, ErrorVectorLengths
	}
	for i := 0; i < len(a); i++ {
		sum += math.Pow(math.Abs(a[i]-b[i]), float64(pow))
	}
	return NthRoot(sum, pow), nil
}

// NthRoot calculates the nth root
func NthRoot(a float64, n int) float64 {
	z := a / float64(n)
	for i := 0; i < 20; i++ {
		if math.Pow(z, float64(n)) == a {
			return z
		}
		z -= (math.Pow(z, float64(n)) - a) / (float64(n) * math.Pow(z, float64(n-1)))
	}
	return z
}

// Probability calculates the the probability two vectors are the same.
// Use for face vectors, mathematically if the probability is 0.85 or
// greater it most likely the same person.
func Probability(a, b Vector) (float64, error) {
	if len(a) != len(b) {
		return 0, ErrorVectorLengths
	}
	dist, _ := EuclideanDistance(a, b)
	return (1 - (dist / 4)), nil

}
