package docer

import (
	_ "embed"
	"encoding/json"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed template.gotempl
var templateStr string
var defaultTemplate *template.Template

func init() {
	defaultTemplate = template.Must(template.New("default").Funcs(map[string]any{
		"JSON": func(data any) string {
			res, _ := json.MarshalIndent(data, "", "\t")
			return string(res)
		},
		"arr": func(data ...any) []any {
			return data
		},
		"intRange": func(start, end int) []int {
			n := end - start + 1
			result := make([]int, n)
			for i := 0; i < n; i++ {
				result[i] = start + i
			}
			return result
		},
		"add": func(a, b int) int {
			return a + b
		},
	}).Parse(templateStr))
}

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

func (t *Type) Field(name string) *Field {
	for _, f := range t.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

type RequestSection struct {
	Description string  `json:"description"`
	Example     any     `json:"example"`
	Types       []*Type `json:"types"`
}

type Doc struct {
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	URL             string          `json:"url"`
	Endpoint        string          `json:"endpoint"`
	Method          string          `json:"method"`
	Headers         []string        `json:"headers"`
	CURL            string          `json:"curl"`
	SuccessResponse any             `json:"success_response"`
	ErrorResponse   any             `json:"error_response"`
	Body            *RequestSection `json:"body"`
	Param           *RequestSection `json:"param"`
	Query           *RequestSection `json:"query"`
}

func New() *Doc {
	doc := &Doc{
		URL:             "",
		Method:          "",
		Headers:         nil,
		SuccessResponse: nil,
		ErrorResponse:   nil,
		Body:            nil,
	}
	return doc
}

func (b *RequestSection) Type(name string) *Type {
	for _, t := range b.Types {
		if t.Name == name {
			return t
		}
	}
	return nil
}

func Read(input string) (*Doc, error) {
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
		doc, err := Read(output)
		if err == nil {
			final = mergeDoc(doc, final)
		}
	}
	*d = *final
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

func (d *Doc) ParseParam(param any, tag string) {
	d.Param = &RequestSection{
		Types: make([]*Type, 0),
	}
	d.Param.Types = NewParser(TagAsName(tag)).parse(param)
}

func (d *Doc) ParseQuery(query any, tag string) {
	d.Query = &RequestSection{
		Types: make([]*Type, 0),
	}
	d.Query.Types = NewParser(TagAsName(tag)).parse(query)
}

func (d *Doc) ParseBody(body any, tag string) {
	d.Body = &RequestSection{
		Types: make([]*Type, 0),
	}
	d.Body.Types = NewParser(TagAsName(tag)).parse(body)
}

func (d *Doc) Generate(output string) error {
	dir := filepath.Dir(output)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()
	return defaultTemplate.Execute(file, d)
}
