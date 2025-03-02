package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func Test_greeting(t *testing.T) {

	buff := new(bytes.Buffer)
	name := "Joseph"

	greeting(buff, name)

	expected := "Hello,Joseph\n"

	got := buff.String()

	if expected != got {
		t.Errorf("Expected value:%v but got value:%v", expected, got)
	}
}

func Test_greeting2(t *testing.T) {
	originalStdout := os.Stdout

	r, w, err := os.Pipe()

	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w

	name := "Janet"
	greeting2(name)

	w.Close()
	os.Stdout = originalStdout
	expected := "Hello,Janet\n"

	var buff bytes.Buffer

	_, err = io.Copy(&buff, r)

	if err != nil {
		t.Fatal(err)
	}

	got := buff.String()

	if expected != got {
		t.Errorf("Expected: %q; Got: %q", expected, got)
	}
}
