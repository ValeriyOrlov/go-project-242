package tests

import (
	"code"
	"testing"
)

func TestGetPathSize_File(t *testing.T) {
	path := "../testdata/fixtures/main.go"
	want := "1436B"
	bytes, err := code.GetPathSize(path, false, false, false)

	if want != bytes || err != nil {
		t.Errorf("unexpected result: got %s, want %s", bytes, want)
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	path := "../testdata/fixtures"
	want := "1471B"
	bytes, err := code.GetPathSize(path, false, false, false)

	if want != bytes || err != nil {
		t.Errorf("unexpected result: got %s, want %s", bytes, want)
	}
}

func TestGetPathSize_EmptyPath(t *testing.T) {
	path := ""
	want := "the path to the file or directory has not been entered"
	_, err := code.GetPathSize(path, false, false, false)

	if err.Error() != want && err != nil {
		t.Errorf("unexpected result: got %s, want %s", err, want)
	}
}

func TestFormatSize_HumanFormat(t *testing.T) {
	testBytes := 123
	testOneKbBytes := 1024
	testOneEbBytes := 1152921504606846976
	wantBytes := "123B"
	resBytes := code.FormatSize(int64(testBytes), false)
	wantOneKb := "1.0KB"
	resOneKb := code.FormatSize(int64(testOneKbBytes), true)
	wantOneEb := "1.0EB"
	resOneEb := code.FormatSize(int64(testOneEbBytes), true)

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

func TestGetSize_AllFormat(t *testing.T) {
	path := "../testdata/fixtures/.hiddenfiletwo"
	want := "2508B"
	hiddenFileSize, err := code.GetPathSize(path, false, false, true)

	if want != hiddenFileSize || err != nil {
		t.Errorf("unexpected result: got %s, want %s.", hiddenFileSize, want)
	}

	want = "0B"
	hiddenFileSizeWithoutFlag, err := code.GetPathSize(path, false, false, false)

	if want != hiddenFileSizeWithoutFlag || err != nil {
		t.Errorf("unexpected result: got %s, want %s.", hiddenFileSizeWithoutFlag, want)
	}
}

func TestGetSize_RecursiveMode(t *testing.T) {
	path := "../testdata/fixtures/"
	want := "1471B"
	withoutRecursiveModeSize, err := code.GetPathSize(path, false, false, false)
	if want != withoutRecursiveModeSize || err != nil {
		t.Errorf("unexpected result: got %s, want %s.", withoutRecursiveModeSize, want)
	}
	want = "2080B"
	withRecursiveModeSize, err := code.GetPathSize(path, true, false, false)
	if want != withRecursiveModeSize || err != nil {
		t.Errorf("unexpected result: got %s, want %s.", withRecursiveModeSize, want)
	}
}
