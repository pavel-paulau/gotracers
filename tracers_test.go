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

func TestStart(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	id, err := Start("mytag")
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

func TestStartInt(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	err = StartInt("mytag", 1234567890)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	bs, err := ioutil.ReadFile(traceFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := fmt.Sprintf(">:%d:mytag::", 1234567890)
	if string(bs) != expected {
		t.Fatalf("exptected %v, got %v", expected, string(bs))
	}
}

func TestStartStr(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	id, err := StartStr("mytag", "fd4086fa-9ed0-465d-9a99-422c5d8e9506")
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

func TestTracer(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	if err = End("mytag", 1234567890); err != nil {
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

func BenchmarkStart(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Start("mytag")
	}
}

func BenchmarkStartInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StartInt("mytag", 1234567890)
	}
}

func BenchmarkStartStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StartStr("mytag", "fd4086fa-9ed0-465d-9a99-422c5d8e9506")
	}
}

func BenchmarkEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		End("mytag", 1234567890)
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

func BenchmarkHash(b *testing.B) {
	s := []byte("fd4086fa-9ed0-465d-9a99-422c5d8e9506")
	for n := 0; n < b.N; n++ {
		hash(s)
	}
}
