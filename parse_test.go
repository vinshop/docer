package doc

import (
	"testing"
)

type mock struct {
	ID  string   `json:"id"`
	Age int      `json:"age"`
	A   *mockA   `json:"a"`
	Bs  []*mockB `json:"bs"`
}

type mockA struct {
	Name string `json:"name"`
}

type mockB struct {
	Age int `json:"age"`
}

func TestParse(t *testing.T) {
	doc := Parse(mock{})
	if err := doc.JSON("test.json"); err != nil {
		t.Fatal(err)
	}
}
