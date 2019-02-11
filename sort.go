package typedict

func DataTypeSorter(dataTypes []*DataType) func(int, int) bool {
	return func(i, j int) bool {
		return (dataTypes[i].PkgPath + "." + dataTypes[i].Name) < (dataTypes[j].PkgPath + "." + dataTypes[j].Name)
	}
}
