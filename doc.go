package doc

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Ref         string `json:"ref"`
	Description string `json:"description"`
}

type Type struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Description string   `json:"description"`
	Fields      []*Field `json:"fields"`
}

type Doc struct {
	URL             string   `json:"url"`
	Method          string   `json:"method"`
	Headers         []string `json:"headers"`
	ExampleBody     any      `json:"example_body"`
	SuccessResponse any      `json:"success_response"`
	ErrorResponse   any      `json:"error_response"`
	Types           []*Type  `json:"types"`
}

func readJSON(input string) (*Doc, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	doc := &Doc{}
	err = dec.Decode(doc)
	return doc, err
}

func (d *Doc) JSON(output string) error {
	final := d
	_, err := os.Stat(output)
	if err == nil {
		// file exist try mergeDoc
		doc, err := readJSON(output)
		if err == nil {
			final = mergeDoc(doc, final)
		}
	}
	// file not exist create new
	dir := filepath.Dir(output)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	return enc.Encode(final)
}
