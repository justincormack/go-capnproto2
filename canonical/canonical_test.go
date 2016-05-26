package canonical

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
	"testing"

	"zombiezen.com/go/capnproto2"
)

func readTestFile(name string) ([]byte, error) {
	path := filepath.Join("testdata", name)
	return ioutil.ReadFile(path)
}

func TestConvert(t *testing.T) {
	tests := []struct {
		inFile  string
		outFile string
	}{
	// TODO(light): add tests
	}
	for _, test := range tests {
		in, err := readTestFile(test.inFile)
		if err != nil {
			t.Errorf("readTestFile(%q): %v", test.inFile, err)
			continue
		}
		want, err := readTestFile(test.outFile)
		if err != nil {
			t.Errorf("readTestFile(%q): %v", test.outFile, err)
			continue
		}
		msg, err := capnp.Unmarshal(in)
		if err != nil {
			t.Errorf("capnp.Unmarshal(readTestFile(%q)): %v", test.inFile, err)
			continue
		}

		cmsg, err := Convert(msg)
		if err != nil {
			t.Errorf("canonical.Convert(capnp.Unmarshal(readTestFile(%q))): %v", test.inFile, err)
			continue
		}

		if cmsg.NumSegments() != 1 {
			t.Errorf("canonical.Convert(capnp.Unmarshal(readTestFile(%q))) num segments = %d; want 1", test.inFile, cmsg.NumSegments())
			continue
		}
		seg, err := cmsg.Segment(0)
		if err != nil {
			t.Errorf("canonical.Convert(capnp.Unmarshal(readTestFile(%q))).Segment(0): %v", test.inFile, err)
			continue
		}
		if !bytes.Equal(seg.Data(), want) {
			t.Errorf("canonical.Convert(capnp.Unmarshal(readTestFile(%q))).Segment(0) =\n%s\n; want:\n%s", test.inFile, hex.Dump(seg.Data()), hex.Dump(want))
		}
	}
}
