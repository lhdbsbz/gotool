package gotool

import (
	"encoding/json"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

type DiffModel int64

const (
	DiffModelDefault              DiffModel = iota
	DiffModelCareFromNotZeroValue DiffModel = iota
	DiffModelCareToNotZeroValue   DiffModel = iota
)

func Diff(fromItem, toItem interface{}, diffModel DiffModel, excludedFields ...string) (diffFieldsFrom map[string]interface{}, diffFieldsTo map[string]interface{}, err error) {
	var excludedFieldsMap = make(map[string]bool)
	if len(excludedFields) > 0 {
		for _, field := range excludedFields {
			if field == "" {
				continue
			}
			excludedFieldsMap[toLowerCamelCase(field)] = true
		}
	}
	diffFieldsFrom, diffFieldsTo = compareStructs(fromItem, toItem, "", excludedFieldsMap)
	if diffModel == DiffModelDefault {
		return
	}
	if diffModel == DiffModelCareFromNotZeroValue {
		toTemp := make(map[string]interface{})
		for key := range diffFieldsFrom {
			if _, ok := diffFieldsTo[key]; ok {
				toTemp[key] = diffFieldsTo[key]
			}
		}
		diffFieldsTo = toTemp
	}

	if diffModel == DiffModelCareToNotZeroValue {
		toTemp := make(map[string]interface{})
		for key := range diffFieldsTo {
			if _, ok := diffFieldsFrom[key]; ok {
				toTemp[key] = diffFieldsFrom[key]
			}
		}
		diffFieldsFrom = toTemp
	}
	return
}

func DiffAndToJson(fromItem, toItem interface{}, diffModel DiffModel, excludedFields ...string) (from string, to string, err error) {
	var (
		diffFieldsFrom map[string]interface{}
		diffFieldsTo   map[string]interface{}
	)
	diffFieldsFrom, diffFieldsTo, err = Diff(fromItem, toItem, diffModel, excludedFields...)
	var (
		diffFieldsFromMarshal []byte
		diffFieldsToMarshal   []byte
	)
	diffFieldsFromMarshal, err = json.Marshal(diffFieldsFrom)
	if err != nil {
		return
	}
	from = string(diffFieldsFromMarshal)
	diffFieldsToMarshal, err = json.Marshal(diffFieldsTo)
	if err != nil {
		return
	}
	to = string(diffFieldsToMarshal)
	return
}

func compareStructs(from, to interface{}, prefix string, excludedFields map[string]bool) (map[string]interface{}, map[string]interface{}) {
	var diffFieldsFrom = make(map[string]interface{})
	var diffFieldsTo = make(map[string]interface{})
	typeFrom := reflect.TypeOf(from)
	valueFrom := reflect.ValueOf(from)
	typeTo := reflect.TypeOf(to)
	valueTo := reflect.ValueOf(to)
	if valueFrom.Kind() == reflect.Pointer {
		valueFrom = valueFrom.Elem()
	}
	if valueTo.Kind() == reflect.Pointer {
		valueTo = valueTo.Elem()
	}
	if typeFrom.Kind() == reflect.Pointer {
		typeFrom = typeFrom.Elem()
	}
	if typeTo.Kind() == reflect.Pointer {
		typeTo = typeTo.Elem()
	}
	for i := 0; i < typeFrom.NumField(); i++ {
		fieldFrom := typeFrom.Field(i)
		valueFieldFrom := valueFrom.Field(i)
		if valueFieldFrom.Kind() == reflect.Pointer {
			valueFieldFrom = valueFieldFrom.Elem()
		}
		if !valueFieldFrom.CanInterface() {
			// 忽略不能导出的字段
			continue
		}
		// 无需diff的字段直接跳过
		if excludedFields[getFieldFullPathName(prefix, fieldFrom.Name)] {
			continue
		}
		fieldTo, found := typeTo.FieldByName(fieldFrom.Name)
		if !found {
			if !isZeroValue(valueFieldFrom.Interface()) {
				diffFieldsFrom[toLowerCamelCase(fieldFrom.Name)] = valueFieldFrom.Interface()
			}
			continue
		}
		valueFieldTo := valueTo.FieldByName(fieldFrom.Name)
		if valueFieldTo.Kind() == reflect.Pointer {
			valueFieldTo = valueFieldTo.Elem()
		}
		if !valueFieldTo.CanInterface() {
			// 忽略不能导出的字段
			continue
		}
		if valueFieldFrom.Kind() == reflect.Struct && valueFieldTo.Kind() == reflect.Struct {
			nestedDiffFieldsFrom, nestedDiffFieldsTo := compareStructs(valueFieldFrom.Interface(), valueFieldTo.Interface(), getFieldFullPathName(prefix, fieldTo.Name), excludedFields)
			if !isZeroValue(nestedDiffFieldsFrom) {
				diffFieldsFrom[toLowerCamelCase(fieldFrom.Name)] = nestedDiffFieldsFrom
			}
			if !isZeroValue(nestedDiffFieldsTo) {
				diffFieldsTo[toLowerCamelCase(fieldTo.Name)] = nestedDiffFieldsTo
			}
		} else if !reflect.DeepEqual(valueFieldFrom.Interface(), valueFieldTo.Interface()) {
			if !isZeroValue(valueFieldFrom.Interface()) {
				diffFieldsFrom[toLowerCamelCase(fieldFrom.Name)] = valueFieldFrom.Interface()
			}
			if !isZeroValue(valueFieldTo.Interface()) {
				diffFieldsTo[toLowerCamelCase(fieldTo.Name)] = valueFieldTo.Interface()
			}
		}
	}

	// Check for additional fields in To
	for i := 0; i < typeTo.NumField(); i++ {
		fieldTo := typeTo.Field(i)
		// 无需diff的字段直接跳过
		if excludedFields[getFieldFullPathName(prefix, fieldTo.Name)] {
			continue
		}
		_, found := typeFrom.FieldByName(fieldTo.Name)
		if !found {
			valueFieldTo := valueTo.Field(i)
			// Check if the current field is excluded
			if excludedFields[getFieldFullPathName(prefix, fieldTo.Name)] {
				continue
			}
			if !valueFieldTo.CanInterface() {
				// 忽略不能导出的字段
				continue
			}
			if !isZeroValue(valueFieldTo.Interface()) {
				diffFieldsTo[toLowerCamelCase(fieldTo.Name)] = valueFieldTo.Interface()
			}
		}
	}
	return diffFieldsFrom, diffFieldsTo
}

func getFieldFullPathName(prefix, fieldName string) string {
	fieldName = toLowerCamelCase(fieldName)
	if prefix != "" {
		return prefix + "." + fieldName
	}
	return fieldName
}

func toLowerCamelCase(fieldName string) string {
	// 将下划线替换为空格
	fieldName = strings.ReplaceAll(fieldName, "_", " ")
	// 将字段名进行单词分割
	words := strings.Fields(fieldName)
	// 将每个单词的首字母大写
	for i := 0; i < len(words); i++ {
		words[i] = cases.Title(language.English).String(words[i])
	}
	// 将单词连接起来形成新的字段名
	fieldName = strings.Join(words, "")
	// 将首个单词的首字母小写
	fieldName = strings.ToLower(fieldName[:1]) + fieldName[1:]
	return fieldName
}

func isZeroValue(field interface{}) bool {
	value := reflect.ValueOf(field)
	zeroValue := reflect.Zero(reflect.TypeOf(field)).Interface()

	// 处理特殊情况：空数组、空字符串和空 map
	switch value.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map:
		return value.Len() == 0
	default:
		return reflect.DeepEqual(value.Interface(), zeroValue)
	}
}
