package doc

import "testing"

func TestRead(t *testing.T) {
	doc, err := readJSON("test.json")
	if err != nil {
		t.Fatal(err)
	}
	if doc == nil {
		t.Fatal("doc is nil")
	}
}
