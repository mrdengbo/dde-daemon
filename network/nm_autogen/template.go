package main

// get key type
const tplGetKeyType = `
// Get key type
func get{{.FieldName | ToFieldFuncBaseName}}KeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown{{range .Keys}}
	case {{.Name}}:
		t = {{.Type}}{{end}}
	}
	return
}
`

// get key's default json value
const tplGetDefaultValueJSON = `{{$fieldFuncBaseName := .FieldName | ToFieldFuncBaseName}}
// Get key's default value
func get{{$fieldFuncBaseName}}KeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		LOGGER.Error("invalid key:", key){{range .Keys}}{{$default := ToKeyTypeDefaultValueJSON .Type .Default}}
	case {{.Name}}:
		valueJSON = ` + "`{{$default}}`" + `{{end}}
	}
	return
}
`

// get json value generally
const tplGeneralGetterJSON = `
{{$fieldFuncBaseName := .FieldName | ToFieldFuncBaseName}}
// Get JSON value generally
func generalGet{{$fieldFuncBaseName}}KeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		LOGGER.Error("generalGet{{.FieldName | ToFieldFuncBaseName}}KeyJSON: invalide key", key){{range .Keys}}
	case {{.Name}}:
		value = get{{.Name | ToKeyFuncBaseName}}JSON(data){{end}}
	}
	return
}
`

// check if key exists
const tplCheckExists = `
// Check if key exists{{$fieldName := .FieldName}}{{range $index, $key := .Keys}}
func is{{$key.Name | ToKeyFuncBaseName}}Exists(data _ConnectionData) bool {
	return isConnectionDataKeyExists(data, {{$fieldName}}, {{$key.Name}})
}{{end}}
`

// getter
const tplGetter = `
// Getter{{$fieldName := .FieldName}}{{range $index, $key := .Keys}}
func get{{$key.Name | ToKeyFuncBaseName}}(data _ConnectionData) (value {{$key.Type | ToKeyTypeRealData}}) {
	value, _ = getConnectionDataKey(data, {{$fieldName}}, {{$key.Name}}).({{$key.Type | ToKeyTypeRealData}})
	return
}{{end}}
`

// setter
const tplSetter = `
// Setter{{$fieldName := .FieldName}}{{range $index, $key := .Keys}}
func set{{$key.Name | ToKeyFuncBaseName}}(data _ConnectionData, value {{$key.Type | ToKeyTypeRealData}}) {
	setConnectionDataKey(data, {{$fieldName}}, {{$key.Name}}, value)
}{{end}}
`

// json getter
const tplJSONGetter = `
// JSON Getter{{$fieldName := .FieldName}}{{range $index, $key := .Keys}}
func get{{$key.Name | ToKeyFuncBaseName}}JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, {{$fieldName}}, {{$key.Name}}, get{{$fieldName | ToFieldFuncBaseName}}KeyType({{$key.Name}}))
	return
}{{end}}
`

// json setter
const tplJSONSetter = `
// JSON Setter{{$fieldName := .FieldName}}{{range $index, $key := .Keys}}
func set{{$key.Name | ToKeyFuncBaseName}}JSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, {{$fieldName}}, {{$key.Name}}, valueJSON, get{{$fieldName | ToFieldFuncBaseName}}KeyType({{$key.Name}}))
}{{end}}
`

// remover
const tplRemover = `
// Remover{{$fieldName := .FieldName}}{{range $index, $key := .Keys}}
func remove{{$key.Name | ToKeyFuncBaseName}}(data _ConnectionData) {
	removeConnectionDataKey(data, {{$fieldName}}, {{$key.Name}})
}{{end}}
`

// TODO
const tplGetAvaiableValues = `// Get avaiable values`
