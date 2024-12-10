package main

import (
	"testing"
	"bytes"
	"io"
	"os"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old // restore the real stdout
	io.Copy(&buf, r)
	return buf.String()
}

func TestMain(t *testing.T) {
	output := captureOutput(main)
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("Unexpected output in main(): expected %q, got %q", expected, output)
	}
}
