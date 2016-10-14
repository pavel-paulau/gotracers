package gotracers

import (
	"io"
	"os"
)

var traceFile *os.File

func init() {
	traceFile, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
}

func StartTracer(tag string) error {
	bs := []byte(">::" + tag + "::")

	n, err := traceFile.Write(bs)
	if err == nil && n < len(bs) {
		return io.ErrShortWrite
	}
	return err
}

func EndTracer(tag string) error {
	bs := []byte("<::" + tag + "::")

	n, err := traceFile.Write(bs)
	if err == nil && n < len(bs) {
		return io.ErrShortWrite
	}
	return err
}
