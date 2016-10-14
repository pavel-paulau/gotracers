package tracers

import (
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

// StartTracer emits a entry tracer event. If unique span ID is not provided,
// StartTracer generates a random 64-bit integer ID.
func StartTracer(tag string, spanId int64) (int64, error) {
	if spanId < 0 {
		spanId = rand.Int63()
	}
	tracer := []byte(">:" + strconv.FormatInt(spanId, 10) + ":" + tag + "::")

	err := writeTracer(tracer)

	return spanId, err
}

// EndTracer emits an exit tracer event for user-provided span ID.
func EndTracer(tag string, spanId int64) error {
	tracer := []byte("<:" + strconv.FormatInt(spanId, 10) + ":" + tag + "::")

	return writeTracer(tracer)
}
