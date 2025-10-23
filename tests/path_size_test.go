package testdata

import (
	"code"
	"testing"
)

func TestGetPathSize_File(t *testing.T) {
	path := "../testdata/rss-aggregator/index.html"
	want := "3589B	../testdata/rss-aggregator/index.html"
	msg, err := code.GetSize(path)

	if want != msg || err != nil {
		t.Errorf("unexpected result: got %s, want %s", msg, want)
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	path := "../testdata/rss-aggregator/"
	want := "242785B	../testdata/rss-aggregator/"
	msg, err := code.GetSize(path)

	if want != msg || err != nil {
		t.Errorf("unexpected result: got %s, want %s", msg, want)
	}
}

func TestGetPathSize_EmptyPath(t *testing.T) {
	path := ""
	want := "The path to the file or directory has not been entered. Run the program with the -h flag to read the help."
	msg, err := code.GetSize(path)

	if want != msg || err != nil {
		t.Errorf("unexpected result: got %s, want %s", msg, want)
	}
}
