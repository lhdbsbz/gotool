package gotool

import (
	"github.com/spf13/cast"
	"strings"
)

func JoinStringInt64(array []int64, split string) string {
	itemArray := make([]string, len(array))
	for index, item := range array {
		itemArray[index] = cast.ToString(item)
	}
	return strings.Join(itemArray, split)
}
