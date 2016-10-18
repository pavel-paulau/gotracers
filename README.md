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
	spanID, err := tracers.Start("mytag", -1)
	handleErr(err)
	defer tracers.End("mytag", spanID)

	...
}
```

Custom span identifiers:

```
func myFunc(uniqueId int64) {
	_, err := tracers.Start("mytag", uniqueId)
	handleErr(err)
	defer tracers.End("mytag", uniqueId)

	...
}
```

Using context:

```
func myFunc(ctx context.Context) {
	_, err := tracers.StartWithContext(ctx, "mytag")
	handleErr(err)
	defer tracers.EndWithContext(ctx, "mytag")

	...
}

```

```
func myFunc() {
	ctx, err := tracers.StartWithContext(context.TODO(), "mytag")
	handleErr(err)
	defer tracers.EndWithContext(ctx, "mytag")

	...
}

```
