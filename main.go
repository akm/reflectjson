package reflectjson

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

func process(objectMap map[string][]interface{}, ptn *regexp.Regexp) {
	res := map[string][]*DataType{}

	for key, objects := range objectMap {
		types := []reflect.Type{}
		for _, obj := range objects {
			types = append(types, reflect.TypeOf(obj))
		}

		types = digTypes(types)

		dataTypes := []*DataType{}
		for _, t := range types {
			dt := newDataType(t)
			if ptn.MatchString(dt.PkgPath) {
				dataTypes = append(dataTypes, dt)
			}
		}

		res[key] = dataTypes
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to json.MarshalIndent because of %v\n", err)
		return
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write because of %v\n", err)
		return
	}
}

type DataFieldType struct {
	PkgPath        string   `json:"PkgPath,omitempty"`
	Name           string   `json:"Name,omitempty"`
	Kinds          []string `json:"Kinds,omitempty"`
	Representation string
}

type DataField struct {
	Name      string
	Type      *DataFieldType
	RawTag    string            `json:"RawTag,omitempty"`
	Tag       map[string]string `json:"Tag,omitempty"`
	Anonymous bool
}

type DataType struct {
	Name    string
	PkgPath string
	Size    uintptr
	Fields  []*DataField
}

func digTypes(types []reflect.Type) []reflect.Type {
	m := map[string]reflect.Type{}
	for _, t := range types {
		key := t.PkgPath() + "." + t.Name()
		m[key] = t
	}

	for _, t := range types {
		digType(m, t)
	}

	r := []reflect.Type{}
	for _, t := range m {
		r = append(r, t)
	}
	return r
}

func digType(m map[string]reflect.Type, t reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		digStruct(m, t)
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		digType(m, t.Elem())
	}
}

func digStruct(m map[string]reflect.Type, t reflect.Type) {
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		f := t.Field(i)
		ft := f.Type
		if ft.PkgPath() != "" {
			key := ft.PkgPath() + "." + ft.Name()
			_, ok := m[key]
			if !ok {
				m[key] = ft
				switch ft.Kind() {
				case reflect.Struct:
					digType(m, ft)
				}
			}
		}
	}
}

func newDataType(t reflect.Type) *DataType {
	r := &DataType{
		Name:    t.Name(),
		PkgPath: t.PkgPath(),
		Size:    t.Size(),
	}
	if t.Kind() != reflect.Struct {
		return r
	}

	fields := []*DataField{}
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		f := t.Field(i)
		ft := f.Type
		fields = append(fields, &DataField{
			Name:      f.Name,
			Anonymous: f.Anonymous,
			Type:      DataFieldTypeFromType(ft),
			Tag:       parseTag(string(f.Tag)),
			RawTag:    string(f.Tag),
		})
	}
	r.Fields = fields
	return r
}

func DataFieldTypeFromType(t reflect.Type) *DataFieldType {
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		r := DataFieldTypeFromType(t.Elem())
		r.Kinds = append(r.Kinds, t.Kind().String())
		return r
	default:
		return &DataFieldType{
			PkgPath:        t.PkgPath(),
			Name:           t.Name(),
			Kinds:          []string{t.Kind().String()},
			Representation: t.String(),
		}
	}
}

var TagParserRE = regexp.MustCompile(`\s*([^:\s]+?):"(.+?)"`)

func parseTag(src string) map[string]string {
	parsed := TagParserRE.FindAllStringSubmatch(src, -1)
	r := map[string]string{}
	for _, parts := range parsed {
		r[parts[1]] = parts[2]
	}
	return r
}
