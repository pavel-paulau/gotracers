package gotracers

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

func writeFile(data []byte) error {
	n, err := traceFile.Write(data)
	if err == nil && n < len(data) {
		return io.ErrShortWrite
	}
	return err
}

func StartTracer(tag string, id int64) (int64, error) {
	if id < 0 {
		id = rand.Int63()
	}
	tracer := []byte(">:" + strconv.FormatInt(id, 10) + ":" + tag + "::")

	err := writeFile(tracer)

	return id, err
}

func EndTracer(tag string, id int64) error {
	tracer := []byte("<:" + strconv.FormatInt(id, 10) + ":" + tag + "::")

	return writeFile(tracer)
}
