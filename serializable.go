package typedict

import (
	"reflect"
	"regexp"
	"sort"
)

func SerializableTypes(types []reflect.Type) DataTypes {
	dataTypes := DataTypes{}
	for _, t := range types {
		dt := NewDataType(t)
		dataTypes = append(dataTypes, dt)
	}
	sort.Slice(dataTypes, DataTypeSorter(dataTypes))
	return dataTypes
}

type DataFieldType struct {
	PkgPath        string `json:",omitempty"`
	Name           string `json:",omitempty"`
	Kind           string `json:",omitempty"`
	Representation string `json:",omitempty"`
}

type DataField struct {
	Name      string
	Type      *DataFieldType
	RawTag    string            `json:",omitempty"`
	Tag       map[string]string `json:",omitempty"`
	Anonymous bool
}

type DataType struct {
	Name           string
	PkgPath        string
	Kind           string
	Size           uintptr
	Fields         []*DataField `json:",omitempty"`
	Elem           *DataType    `json:",omitempty"`
	Representation string       `json:",omitempty"`
}

type DataTypes []*DataType

func (s DataTypes) DetectByName(name string) *DataType {
	for _, i := range s {
		if i.Name == name {
			return i
		}
	}
	return nil
}

func NewDataType(t reflect.Type) *DataType {
	r := &DataType{
		Name:           t.Name(),
		PkgPath:        t.PkgPath(),
		Kind:           t.Kind().String(),
		Size:           t.Size(),
		Representation: t.String(),
	}
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		if elem := t.Elem(); elem != nil {
			r.Elem = NewDataType(elem)
		}
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
	r := &DataFieldType{
		PkgPath:        t.PkgPath(),
		Name:           t.Name(),
		Kind:           t.Kind().String(),
		Representation: t.String(),
	}
	// switch t.Kind() {
	// case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
	// 	if elem := t.Elem(); elem != nil {
	// 		r.Elem = NewDataType(elem)
	// 	}
	// }
	return r
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
