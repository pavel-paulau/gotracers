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

func StartTracer(tag string, id int64) (int64, error) {
	if id < 0 {
		id = rand.Int63()
	}
	bs := []byte(">:" + strconv.FormatInt(id, 10) + ":" + tag + "::")

	n, err := traceFile.Write(bs)
	if err == nil && n < len(bs) {
		return id, io.ErrShortWrite
	}
	return id, err
}

func EndTracer(tag string, id int64) error {
	bs := []byte("<:" + strconv.FormatInt(id, 10) + ":" + tag + "::")

	n, err := traceFile.Write(bs)
	if err == nil && n < len(bs) {
		return io.ErrShortWrite
	}
	return err
}
