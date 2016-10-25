package tracers

import (
	"hash/fnv"
	"io"
	"math"
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

func hash(s []byte) uint64 {
	h := fnv.New64a()
	h.Write(s)
	return h.Sum64()
}

func iTou(i int64) uint64 {
	return uint64(math.MaxInt64 + i)
}

// Start generates a random span identifier and emits an entry tracer event.
func Start(tag string) (uint64, error) {
	spanID := iTou(rand.Int63())
	tracer := []byte(">:" + strconv.FormatUint(spanID, 10) + ":" + tag + "::")

	err := writeTracer(tracer)

	return spanID, err
}

// StartInt emits an entry tracer event using user-specified span identifier.
func StartInt(tag string, spanID uint64) error {
	tracer := []byte(">:" + strconv.FormatUint(spanID, 10) + ":" + tag + "::")

	return writeTracer(tracer)
}

// StartStr emits an entry tracer event using user-specified span identifier.
func StartStr(tag string, spanID string) (uint64, error) {
	intID := hash([]byte(spanID))
	tracer := []byte(">:" + strconv.FormatUint(intID, 10) + ":" + tag + "::")

	err := writeTracer(tracer)

	return intID, err
}

// End emits an exit tracer event for the user-provided span identifier.
func End(tag string, spanID uint64) error {
	tracer := []byte("<:" + strconv.FormatUint(spanID, 10) + ":" + tag + "::")

	return writeTracer(tracer)
}
