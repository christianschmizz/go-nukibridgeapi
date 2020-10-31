package testhelper

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

// GetGoldenFile returns the golden file content. If the `update` is specified, it updates the
// file with the current output and returns it.
func GetGoldenFile(t *testing.T, actual []byte, fileName string, update bool) []byte {
	golden := filepath.Join("testdata", fileName)
	if update {
		if err := ioutil.WriteFile(golden, actual, 0644); err != nil {
			t.Fatalf("Error writing golden file for filename=%s: %s", fileName, err)
		}
	}
	expected, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	return expected
}
