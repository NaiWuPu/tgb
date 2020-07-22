package zip

import (
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	f1, err := os.Open("/home/zzw/test_data/ziptest/gophercolor16x16.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()
	f2, err := os.Open("/home/zzw/test_data/ziptest/readme.notzip")
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()
	f3, err := os.Open("/home/zzw/test_data")
	if err != nil {
		t.Fatal(err)
	}
	defer f3.Close()
	var files = []*os.File{f1, f2, f3}
	dest := "/home/zzw/test_data/test.zip"
	err = Compress(files, dest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeCompress(t *testing.T) {
	err := DeCompress("/home/zzw/test_data/test.zip", "/home/zzw/test_data/de")
	if err != nil {
		t.Fatal(err)
	}
}
