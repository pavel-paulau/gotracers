package gotracers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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

	if err = StartTracer("mytag"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	bs, err := ioutil.ReadFile(traceFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(bs) != ">::mytag::" {
		t.Fatalf("unexpected string: %v", string(bs))
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

	if err = EndTracer("mytag"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	bs, err := ioutil.ReadFile(traceFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(bs) != "<::mytag::" {
		t.Fatalf("unexpected string: %v", string(bs))
	}
}

func BenchmarkConcat(b *testing.B) {
	var bs []byte
	for n := 0; n < b.N; n++ {
		str := ">::" + "mytag" + "::"
		bs = []byte(str)
	}

	if string(bs) != ">::mytag::" {
		b.Errorf("unexpected string: %v", string(bs))
	}
}

func BenchmarkFmt(b *testing.B) {
	var bs []byte
	for n := 0; n < b.N; n++ {
		str := fmt.Sprintf(">::%s::", "mytag")
		bs = []byte(str)
	}

	if string(bs) != ">::mytag::" {
		b.Errorf("unexpected string: %v", string(bs))
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buffer := bytes.Buffer{}
		buffer.WriteString(">::")
		buffer.WriteString("mytag")
		buffer.WriteString("::")
	}
}
func BenchmarkWriteToFile(b *testing.B) {
	data := []byte(">::mytag::")

	for n := 0; n < b.N; n++ {
		ioutil.WriteFile(os.DevNull, data, 0644)
	}
}

func BenchmarkStartTracer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StartTracer("mytag")
	}
}
