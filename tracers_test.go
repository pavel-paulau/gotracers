package tracers

import (
	"bytes"
	"context"
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

	id, err := Start("mytag", -1)
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

func TestStartWithContext(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	ctx := context.WithValue(context.TODO(), "span", int64(1234567890))

	newCtx, err := StartWithContext(ctx, "mytag")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	spanID, ok := newCtx.Value("span").(int64)
	if !ok {
		t.Fatal("span is missing")
	}

	if spanID != int64(1234567890) {
		t.Fatalf("unxpected span id: %v", spanID)
	}
}

func TestStartWithEmptyContext(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	ctx := context.TODO()

	newCtx, err := StartWithContext(ctx, "mytag")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := newCtx.Value("span").(int64); !ok {
		t.Fatal("span is missing")
	}
}

func TestEndWithContext(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	ctx := context.WithValue(context.TODO(), "span", int64(1234567890))

	err = EndWithContext(ctx, "mytag")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEndWithBadContext(t *testing.T) {
	var err error
	traceFile, err = ioutil.TempFile(".", "tracers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer traceFile.Close()
	defer os.Remove(traceFile.Name())

	ctx := context.TODO()

	err = EndWithContext(ctx, "1234567890")
	if err == nil {
		t.Fatal("unexpected success")
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
		Start("mytag", -1)
	}
}

func BenchmarkEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		End("mytag", 1234567890)
	}
}

func BenchmarkStartWithContext(b *testing.B) {
	ctx := context.WithValue(context.TODO(), "span", 1234567890)

	for n := 0; n < b.N; n++ {
		StartWithContext(ctx, "mytag")
	}
}

func BenchmarkStartWithEmptyContext(b *testing.B) {
	ctx := context.TODO()

	for n := 0; n < b.N; n++ {
		StartWithContext(ctx, "mytag")
	}
}

func BenchmarkEndWithContext(b *testing.B) {
	ctx := context.WithValue(context.TODO(), "span", 1234567890)

	for n := 0; n < b.N; n++ {
		EndWithContext(ctx, "mytag")
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
