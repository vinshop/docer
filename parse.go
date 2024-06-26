package docer

import (
	"fmt"
	"reflect"
)

type ConfigOption func(p *Parser)

func TagAsName(tag string) ConfigOption {
	return func(p *Parser) {
		p.tagAsName = tag
	}
}

type Parser struct {
	memType   map[string]*Type
	types     []*Type
	tagAsName string
}

func NewParser(ops ...ConfigOption) *Parser {
	p := &Parser{
		memType: make(map[string]*Type),
		types:   make([]*Type, 0),
	}
	for _, op := range ops {
		op(p)
	}
	return p
}

func (p *Parser) parse(data any) []*Type {
	p.parseStruct(reflect.TypeOf(data))
	return p.types
}

func (p *Parser) parseStruct(data reflect.Type) *Type {
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}
	if data.Kind() != reflect.Struct {
		return nil
	}
	if t, ok := p.memType[data.Name()]; ok {
		return t
	}
	fmt.Println("type", data.Name(), data.Kind())
	t := &Type{
		Name:        data.Name(),
		DisplayName: "",
		Description: "",
		Fields:      make([]*Field, 0),
	}
	p.memType[data.Name()] = t
	p.types = append(p.types, t)

	for i := 0; i < data.NumField(); i++ {
		f := data.Field(i)
		fmt.Println("-", data.Name()+"."+f.Name, f.Type.Kind())
		jsonTag := f.Tag.Get(p.tagAsName)
		field := &Field{
			Name:        jsonTag,
			Type:        "",
			Required:    false,
			Ref:         "",
			Description: "",
		}
		subT := f.Type
		k := subT.Kind()
		if k == reflect.Ptr {
			subT = subT.Elem()
			k = subT.Kind()
		}

		switch k {
		case reflect.Struct:
			sub := p.parseStruct(subT)
			field.Type = "object"
			field.Ref = sub.Name
		case reflect.Slice, reflect.Array:
			sub := p.parseStruct(subT.Elem())
			if sub != nil {
				field.Ref = sub.Name
				field.Type = "array of object"
			} else {
				field.Type = "array of " + subT.Elem().String()
			}
		default:
			field.Type = subT.String()
		}
		t.Fields = append(t.Fields, field)
	}
	return t
}
