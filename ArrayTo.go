package gotool

func ArrayToArray[SourceType any, ResultType any](sourceArray []SourceType, convertFn func(item SourceType) ResultType) (result []ResultType) {
	result = make([]ResultType, len(sourceArray))
	for index := range sourceArray {
		item := sourceArray[index]
		result[index] = convertFn(item)
	}
	return
}

func ArrayDistinct[SourceType comparable](sourceArray []SourceType) (result []SourceType) {
	result = make([]SourceType, 0, len(sourceArray))
	distinctMap := make(map[SourceType]bool)
	for index := range sourceArray {
		item := sourceArray[index]
		distinctMap[item] = true
	}
	for item := range distinctMap {
		result = append(result, item)
	}
	return
}

func ArrayToMap[ArrayType any, KeyType comparable, ValueType any](array []ArrayType, getKeyAndValue func(ArrayType) (KeyType, ValueType)) (result map[KeyType]ValueType) {
	result = make(map[KeyType]ValueType)
	for index := range array {
		item := array[index]
		key, value := getKeyAndValue(item)
		result[key] = value
	}
	return
}

func GetMapKeys[KeyType comparable, ValueType any](source map[KeyType]ValueType) (result []KeyType) {
	result = make([]KeyType, 0)
	for key := range source {
		result = append(result, key)
	}
	return
}

func GetMapValues[KeyType comparable, ValueType any](source map[KeyType]ValueType) []ValueType {
	result := make([]ValueType, 0)
	for _, value := range source {
		result = append(result, value)
	}
	return result
}
