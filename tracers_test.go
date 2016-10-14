package tracers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestStartTracer(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	id, err := StartTracer("mytag", -1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	bs, err := ioutil.ReadFile(traceFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := fmt.Sprintf(">:%d:mytag::", id)
	if string(bs) != expected {
		t.Fatalf("exptected %v, got %v", expected, string(bs))
	}
}

func TestEndTracer(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	if err = EndTracer("mytag", 1234567890); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	bs, err := ioutil.ReadFile(traceFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(bs) != "<:1234567890:mytag::" {
		t.Fatalf("unexpected string: %v", string(bs))
	}
}

func BenchmarkConcat(b *testing.B) {
	var bs []byte
	id := "1234567890"
	tag := "mytag"
	for n := 0; n < b.N; n++ {
		str := ">:" + id + ":" + tag + "::"
		bs = []byte(str)
	}

	if string(bs) != ">:1234567890:mytag::" {
		b.Errorf("unexpected string: %v", string(bs))
	}
}

func BenchmarkFmt(b *testing.B) {
	var bs []byte
	id := "1234567890"
	tag := "mytag"
	for n := 0; n < b.N; n++ {
		str := fmt.Sprintf(">:%s:%s::", id, tag)
		bs = []byte(str)
	}

	if string(bs) != ">:1234567890:mytag::" {
		b.Errorf("unexpected string: %v", string(bs))
	}
}

func BenchmarkBuffer(b *testing.B) {
	id := "1234567890"
	tag := "mytag"
	for n := 0; n < b.N; n++ {
		buffer := bytes.Buffer{}
		buffer.WriteString(">:")
		buffer.WriteString(id)
		buffer.WriteString(":")
		buffer.WriteString(tag)
		buffer.WriteString("::")
	}
}

func BenchmarkWriteToFile(b *testing.B) {
	data := []byte(">:1234567890:mytag::")
	traceFile, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0644)

	for n := 0; n < b.N; n++ {
		writeTracer(data)
	}
}

func BenchmarkUtilWriteToFile(b *testing.B) {
	data := []byte(">:1234567890:mytag::")

	for n := 0; n < b.N; n++ {
		ioutil.WriteFile(os.DevNull, data, 0644)
	}
}

func BenchmarkStartTracer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StartTracer("mytag", -1)
	}
}

func BenchmarkEndTracer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EndTracer("mytag", 1234567890)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rand.Int63()
	}
}

func BenchmarkItoa(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.FormatInt(1234567890, 10)
	}
}
