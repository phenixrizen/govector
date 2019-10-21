package govector

import (
	"errors"
	"math/rand"
	"time"
)

// kmeans is a simple k-means clusterer that determines centroids with the Train function,
// and then classifies additional observations with the Nearest function.

// Node represents an observation of floating point values
type Nodes []Vector

var RandSeed = time.Now().UnixNano()

var ErrorClusterLength = errors.New("the length of requested clusters is larger than the vectors nodes count")

// Train takes an array of Nodes (observations), and produces as many centroids as specified by
// clusterCount. It will stop adjusting centroids after maxRounds is reached. If there are less
// observations than the number of centroids requested, then Train will return (false, nil).
func Train(nodes Nodes, clusterCount int, maxRounds int) (Nodes, error) {
	if len(nodes) < clusterCount {
		return nil, ErrorClusterLength
	}

	// Check to make sure everything is consistent, dimension-wise
	stdLen := 0
	for i, node := range nodes {
		curLen := len(node)
		if i == 0 {
			stdLen = curLen
		}
		if i > 0 && len(node) != stdLen {
			return nil, ErrorVectorLengths
		}

	}

	centroids := make(Nodes, clusterCount)

	r := rand.New(rand.NewSource(RandSeed))

	// Pick centroid starting points from Nodes
	for i := 0; i < clusterCount; i++ {
		srcIndex := r.Intn(len(nodes))
		srcLen := len(nodes[srcIndex])
		n := make(Vector, srcLen)
		centroids[i] = n
		copy(centroids[i], nodes[r.Intn(len(nodes))])
	}

	// Train centroids
	movement := true
	for i := 0; i < maxRounds && movement; i++ {
		movement = false

		groups := make(map[int][]Vector)

		for _, node := range nodes {
			near := Nearest(node, centroids)
			groups[near] = append(groups[near], node)
		}

		for key, group := range groups {
			newNode := meanNode(group)

			if !Equal(centroids[key], newNode) {
				centroids[key] = newNode
				movement = true
			}
		}
	}

	return centroids, nil
}

// Nearest return the index of the closest centroid from nodes
func Nearest(in Vector, nodes Nodes) int {
	count := len(nodes)

	results := make(Vector, count)
	cnt := make(chan int)
	for i, node := range nodes {
		go func(i int, node Vector, in Vector) {
			results[i] = distance(in, node)
			cnt <- 1
		}(i, node, in)
	}

	wait(cnt, results)

	mindex := 0
	curdist := results[0]

	for i, dist := range results {
		if dist < curdist {
			curdist = dist
			mindex = i
		}
	}

	return mindex
}

// Distance determines the square Euclidean distance between two nodes
func distance(node1, node2 Vector) float64 {
	length := len(node1)
	squares := make(Vector, length, length)

	cnt := make(chan int)

	for i, _ := range node1 {
		go func(i int) {
			diff := node1[i] - node2[i]
			squares[i] = diff * diff
			cnt <- 1
		}(i)
	}

	wait(cnt, squares)

	sum := 0.0
	for _, val := range squares {
		sum += val
	}

	return sum
}

// meanNode takes an array of Nodes and returns a node which represents the average
// value for the provided nodes. This is used to center the centroids within their cluster.
func meanNode(values Nodes) Vector {
	newNode := make(Vector, len(values[0]))

	for _, value := range values {
		for j := 0; j < len(newNode); j++ {
			newNode[j] += (value)[j]
		}
	}

	for i, value := range newNode {
		newNode[i] = value / float64(len(values))
	}

	return newNode
}

// wait stops a function from continuing until the provided channel has processed as
// many items as there are dimensions in the provided Node.
func wait(c chan int, values Vector) {
	count := len(values)

	<-c
	for respCnt := 1; respCnt < count; respCnt++ {
		<-c
	}
}
