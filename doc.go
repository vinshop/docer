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
			res, _ := json.MarshalIndent(data, "", "    ")
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

type Example struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

type Collection struct {
	Description string    `json:"description"`
	Examples    []Example `json:"examples"`
	Types       []*Type   `json:"types"`
}

type Doc struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	Endpoint    string      `json:"endpoint"`
	Method      string      `json:"method"`
	Headers     []string    `json:"headers"`
	Examples    []Example   `json:"examples"`
	Body        *Collection `json:"body"`
	Param       *Collection `json:"param"`
	Query       *Collection `json:"query"`
	Response    *Collection `json:"response"`
}

func New() *Doc {
	doc := &Doc{
		URL:     "",
		Method:  "",
		Headers: nil,
		Body:    nil,
	}
	return doc
}

func (b *Collection) Type(name string) *Type {
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
	enc.SetIndent("", "    ")
	return enc.Encode(final)
}

func (d *Doc) HasResponse(response any, tag string) *Doc {
	d.Response = &Collection{
		Types:    make([]*Type, 0),
		Examples: make([]Example, 0),
	}
	if response != nil {
		d.Response.Types = NewParser(TagAsName(tag)).parse(response)
	}
	return d
}

func (d *Doc) HasParam(param any, tag string) *Doc {
	d.Param = &Collection{
		Types: make([]*Type, 0),
	}
	if param != nil {
		d.Param.Types = NewParser(TagAsName(tag)).parse(param)
	}
	return d
}

func (d *Doc) HasQuery(query any, tag string) *Doc {
	d.Query = &Collection{
		Types: make([]*Type, 0),
	}
	if query == nil {
		d.Query.Types = NewParser(TagAsName(tag)).parse(query)
	}
	return d
}

func (d *Doc) HasBody(body any, tag string) *Doc {
	d.Body = &Collection{
		Types: make([]*Type, 0),
	}
	if body != nil {
		d.Body.Types = NewParser(TagAsName(tag)).parse(body)
	}
	return d
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
