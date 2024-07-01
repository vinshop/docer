package docer

import (
	"reflect"
)

func getNew[T any](old, new T) T {
	v := reflect.ValueOf(new)
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		if v.IsNil() || v.Len() == 0 {
			return old
		}
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return old
		}
	case reflect.Invalid:
		return old
	}

	if reflect.DeepEqual(new, reflect.Zero(reflect.TypeOf(new)).Interface()) {
		return old
	}
	return new
}

// mergeDoc two doc
func mergeDoc(old, new *Doc) *Doc {
	d := &Doc{
		Name:        getNew(old.Name, new.Name),
		Description: getNew(old.Description, new.Description),
		Endpoint:    getNew(old.Endpoint, new.Endpoint),
		URL:         getNew(old.URL, new.URL),
		Method:      getNew(old.Method, new.Method),
		Headers:     getNew(old.Headers, new.Headers),
		Examples:    getNew(old.Examples, new.Examples),
		Body:        mergeRequestSection(old.Body, new.Body),
		Param:       mergeRequestSection(old.Param, new.Param),
		Query:       mergeRequestSection(old.Query, new.Query),
		Response:    mergeRequestSection(old.Response, new.Response),
	}

	return d
}

func mergeRequestSection(old, new *Collection) *Collection {
	if new == nil {
		return nil
	}
	if old == nil {
		return new
	}
	b := &Collection{
		Types:       nil,
		Description: getNew(old.Description, new.Description),
		Examples:    getNew(old.Examples, new.Examples),
	}
	mpNew := make(map[string]*Type)
	for _, t := range new.Types {
		mpNew[t.Name] = t
	}
	for _, t := range old.Types {
		_, ok := mpNew[t.Name]
		if ok {
			mpNew[t.Name] = mergeType(t, mpNew[t.Name])
		}
	}
	for i := range new.Types {
		b.Types = append(b.Types, mpNew[new.Types[i].Name])
	}

	return b
}

func mergeType(old, new *Type) *Type {
	t := &Type{
		Name:        getNew(old.Name, new.Name),
		DisplayName: getNew(old.DisplayName, new.DisplayName),
		Description: getNew(old.Description, new.Description),
		Fields:      nil,
	}
	mpNew := make(map[string]*Field)
	for _, f := range new.Fields {
		mpNew[f.Name] = f
	}
	for _, f := range old.Fields {
		_, ok := mpNew[f.Name]
		if ok {
			mpNew[f.Name] = mergeField(f, mpNew[f.Name])
		}
	}
	for i := range new.Fields {
		t.Fields = append(t.Fields, mpNew[new.Fields[i].Name])
	}

	return t
}

func mergeField(old, new *Field) *Field {
	return &Field{
		Name:        getNew(old.Name, new.Name),
		Type:        getNew(old.Type, new.Type),
		Required:    getNew(old.Required, new.Required),
		Ref:         getNew(old.Ref, new.Ref),
		Description: getNew(old.Description, new.Description),
	}
}
