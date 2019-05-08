package typedict

import (
	"reflect"
	"regexp"
	"sort"
)

func SerializableTypes(types []reflect.Type) []*DataType {
	dataTypes := []*DataType{}
	for _, t := range types {
		dt := NewDataType(t)
		dataTypes = append(dataTypes, dt)
	}
	sort.Slice(dataTypes, DataTypeSorter(dataTypes))
	return dataTypes
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
	Kinds   []string
	Size    uintptr
	Fields  []*DataField
	Representation string
}

func NewDataType(t reflect.Type) *DataType {
	kindNames, end := KindNamesAndEnd(t)
	r := &DataType{
		Name:    end.Name(),
		PkgPath: end.PkgPath(),
		Kinds:   kindNames,
		Size:    t.Size(),
		Representation: t.String(),
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
	kindNames, end := KindNamesAndEnd(t)
	return &DataFieldType{
		PkgPath:        end.PkgPath(),
		Name:           end.Name(),
		Kinds:          kindNames,
		Representation: t.String(),
	}
}

func KindNamesAndEnd(t reflect.Type) ([]string, reflect.Type) {
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		kindNames, end := KindNamesAndEnd(t.Elem())
		return append(kindNames, t.Kind().String()), end
	default:
		return []string{t.Kind().String()}, t
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
