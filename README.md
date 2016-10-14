# gotracers

[![codebeat badge](https://codebeat.co/badges/95dc70af-969e-404e-954f-42bcf337725b)](https://codebeat.co/projects/github-com-pavel-paulau-gotracers)
[![Go Report Card](https://goreportcard.com/badge/github.com/pavel-paulau/gotracers)](https://goreportcard.com/report/github.com/pavel-paulau/gotracers)
[![Build Status](https://travis-ci.org/pavel-paulau/gotracers.svg?branch=master)](https://travis-ci.org/pavel-paulau/gotracers)
[![Coverage Status](https://coveralls.io/repos/github/pavel-paulau/gotracers/badge.svg?branch=master)](https://coveralls.io/github/pavel-paulau/gotracers?branch=master)

gotracer provides emitters for [Sysdig tracers](https://github.com/draios/sysdig/wiki/Tracers).

# Usage

Auto-generated span identifiers:

```
func myFunc() {
	spanID, err := tracers.StartTracer("mytag", -1)
	handleErr(err)
	defer tracers.EndTracer("mytag", spanID)

	...
}
```

Custom span identifiers:

```
func myFunc(uniqueId int64) {
	_, err := tracers.StartTracer("mytag", uniqueId)
	handleErr(err)
	defer tracers.EndTracer("mytag", uniqueId)

	...
}
```
