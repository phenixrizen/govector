# govector

[![Build Status](https://travis-ci.org/phenixrizen/govector.svg?branch=master)](https://travis-ci.org/phenixrizen/govector) [![GoDoc](https://godoc.org/github.com/phenixrizen/govector?status.svg)](https://godoc.org/github.com/phenixrizen/govector) [![Coverage Status](https://coveralls.io/repos/github/phenixrizen/govector/badge.svg?branch=master)](https://coveralls.io/github/phenixrizen/govector?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/phenixrizen/govector)](https://goreportcard.com/report/github.com/phenixrizen/govector)

Provides a nice vector API for Golang.

### Get Started

#### Installation
```bash
$ go get github.com/phenixrizen/govector
```

#### Usage

```go
package main

import (
	"fmt"
	"math"

	"github.com/phenixrizen/govector"
)

func main() {
	// create a Vector type from an int array
	x, err := govector.AsVector([]int{1, 2, 3, 4, 6, 5})
	if err != nil {
		panic(err)
	}

	// create a Vector type from a float64 array, to be used for weights
	w, err := govector.AsVector([]float64{1.0, 1.0, 1.0, 1.0, 1.0, 4.0})
	if err != nil {
		panic(err)
	}

	// find the differences of the Vector x
	dx := x.Diff()
	fmt.Printf("%+v\n", dx)

	// Generate the empirical CDF function for x
	empirical := x.Ecdf()

	// Calculate the percentile from the empirical CDF of x
	percentile := empirical(2.4)
	fmt.Println(percentile)

	// Calculate the weighted mean of x using weights w
	wm, err := x.WeightedMean(w)
	if err != nil {
		panic(err)
	}
	fmt.Println(wm)

	// Calculate the 5% and 95% quantiles of x
	q, err := govector.AsVector([]float64{0.05, 0.95})
	if err != nil {
		panic(err)
	}
	quantiles := x.Quantiles(q)
	fmt.Printf("%+v\n", quantiles)

	// cumulative sum
	sum := x.Cumsum()
	fmt.Printf("%+v\n", sum)

	// shuffle x
	shuffled := x.Shuffle()
	fmt.Printf("%+v\n", shuffled)

	// Apply arbitrary functions to vectors
	y := x.Apply(empirical)
	fmt.Printf("%+v\n", y)
	z := x.Apply(math.Sqrt)
	fmt.Printf("%+v\n", z)

	nodes := govector.Nodes{x, w, w, y, z}
	// Perform kmeans clustering
	centroids, err := govector.Train(nodes, 2, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", centroids)
	idx := govector.Nearest(w, centroids)
	fmt.Printf("The nearest centroid is: %+v\n", centroids[idx])

	dist, err := govector.EuclideanDistance(x, y)
	if err != nil {
		panic(err)
	}
	fmt.Println(dist)

	dist, err = govector.EuclideanDistance(x, x)
	if err != nil {
		panic(err)
	}
	fmt.Println(dist)
}
```

### Credits
This library was originally a fork of:
- https://github.com/drewlanenga/govector