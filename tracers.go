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

// Start emits an entry tracer event. If unique span identifier is not provided,
// StartTracer generates a random 64-bit integer ID.
func Start(tag string, spanID int64) (int64, error) {
	if spanID < 0 {
		spanID = rand.Int63()
	}
	tracer := []byte(">:" + strconv.FormatInt(spanID, 10) + ":" + tag + "::")

	err := writeTracer(tracer)

	return spanID, err
}

// End emits an exit tracer event for the user-provided span identifier.
func End(tag string, spanID int64) error {
	tracer := []byte("<:" + strconv.FormatInt(spanID, 10) + ":" + tag + "::")

	return writeTracer(tracer)
}
