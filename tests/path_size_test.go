package tests

import (
	"code"
	"testing"
)

func TestGetPathSize_File(t *testing.T) {
	path := "../testdata/fixtures/main.go"
	want := int64(1436)
	bytes, err := code.GetSize(path, []string{})

	if want != bytes || err != nil {
		t.Errorf("unexpected result: got %d, want %d", bytes, want)
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	path := "../testdata/fixtures"
	want := int64(1471)
	bytes, err := code.GetSize(path, []string{})

	if want != bytes || err != nil {
		t.Errorf("unexpected result: got %d, want %d", bytes, want)
	}
}

func TestGetPathSize_EmptyPath(t *testing.T) {
	path := ""
	want := "the path to the file or directory has not been entered"
	_, err := code.GetSize(path, []string{})

	if err.Error() != want && err != nil {
		t.Errorf("unexpected result: got %s, want %s", err, want)
	}
}

func TestFormatSize_HumanFormat(t *testing.T) {
	testBytes := 123
	testOneKbBytes := 1024
	testOneEbBytes := 1152921504606846976
	wantBytes := "123B"
	resBytes := code.FormatSize(int64(testBytes), []string{})
	wantOneKb := "1.0KB"
	resOneKb := code.FormatSize(int64(testOneKbBytes), []string{"human"})
	wantOneEb := "1.0EB"
	resOneEb := code.FormatSize(int64(testOneEbBytes), []string{"human"})

	if resBytes != wantBytes {
		t.Errorf("unexpected result: got %s, want %s", resBytes, wantBytes)
	}

	if resOneKb != wantOneKb {
		t.Errorf("unexpected result: got %s, want %s", resOneKb, wantOneKb)
	}

	if resOneEb != wantOneEb {
		t.Errorf("unexpected result: got %s, want %s", resOneEb, wantOneEb)
	}
}
