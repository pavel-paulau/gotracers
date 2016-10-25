# gotracers

[![codebeat badge](https://codebeat.co/badges/95dc70af-969e-404e-954f-42bcf337725b)](https://codebeat.co/projects/github-com-pavel-paulau-gotracers)
[![Go Report Card](https://goreportcard.com/badge/github.com/pavel-paulau/gotracers)](https://goreportcard.com/report/github.com/pavel-paulau/gotracers)
[![Build Status](https://travis-ci.org/pavel-paulau/gotracers.svg?branch=master)](https://travis-ci.org/pavel-paulau/gotracers)
[![Coverage Status](https://coveralls.io/repos/github/pavel-paulau/gotracers/badge.svg?branch=master)](https://coveralls.io/github/pavel-paulau/gotracers?branch=master)
[![GoDoc](https://godoc.org/github.com/pavel-paulau/gotracers?status.svg)](https://godoc.org/github.com/pavel-paulau/gotracers)

gotracer provides emitters for [Sysdig tracers](https://github.com/draios/sysdig/wiki/Tracers).

# Usage

Auto-generated span identifiers:

```
func myFunc() {
	spanID, err := tracers.Start("mytag")
	handleErr(err)
	defer tracers.End("mytag", spanID)

	...
}
```

Custom span identifiers:

```
func myFunc(uniqueId uint64) {
	err := tracers.StartInt("mytag", uniqueId)
	handleErr(err)
	defer tracers.End("mytag", uniqueId)

	...
}
```

```
func myFunc(uniqueId string) {
	spanID, err := tracers.StartStr("mytag", uniqueId)
	handleErr(err)
	defer tracers.End("mytag", spanID)

	...
}
```

# Benchmarks

```
BenchmarkStart-8          	 3000000	       386 ns/op	      32 B/op	       1 allocs/op
BenchmarkStartInt-8       	 5000000	       329 ns/op	      16 B/op	       1 allocs/op
BenchmarkStartStr-8       	 3000000	       523 ns/op	      88 B/op	       3 allocs/op
BenchmarkEnd-8            	 5000000	       363 ns/op	      16 B/op	       1 allocs/op
```
