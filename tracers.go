package tracers

import (
	"context"
	"errors"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var traceFile *os.File

func init() {
	traceFile, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
	rand.Seed(time.Now().UnixNano())
}

func writeTracer(data []byte) error {
	n, err := traceFile.Write(data)
	if err == nil && n < len(data) {
		return io.ErrShortWrite
	}
	return err
}

// Start emits an entry tracer event. If unique span identifier is not provided,
// StartTracer generates a random 64-bit integer ID.
func Start(tag string, spanId int64) (int64, error) {
	if spanId < 0 {
		spanId = rand.Int63()
	}
	tracer := []byte(">:" + strconv.FormatInt(spanId, 10) + ":" + tag + "::")

	err := writeTracer(tracer)

	return spanId, err
}

// End emits an exit tracer event for the user-provided span identifier.
func End(tag string, spanId int64) error {
	tracer := []byte("<:" + strconv.FormatInt(spanId, 10) + ":" + tag + "::")

	return writeTracer(tracer)
}

// StartWithContext emits an entry tracer event and conditionally update the
// span identifier in the parent context.
func StartWithContext(ctx context.Context, tag string) (context.Context, error) {
	span, ok := ctx.Value("span").(int64)

	var spanId int64
	if ok {
		spanId = span
	} else {
		spanId = rand.Int63()
	}

	tracer := []byte(">:" + strconv.FormatInt(spanId, 10) + ":" + tag + "::")

	err := writeTracer(tracer)

	if ok {
		return ctx, err
	}
	return context.WithValue(ctx, "span", spanId), err
}

// EndWithContext emits an exit tracer event for the user-provided context.
// EndTracer returns an error if the parent context doesn't have a span identifier.
func EndWithContext(ctx context.Context, tag string) error {
	spanId, ok := ctx.Value("span").(int64)
	if !ok {
		return errors.New("missing span in context")
	}

	tracer := []byte("<:" + strconv.FormatInt(spanId, 10) + ":" + tag + "::")

	return writeTracer(tracer)
}
