package gotool

import (
	"fmt"
	"testing"
)

func TestArrayDistinct(t *testing.T) {
	abc := []int64{
		1, 1, 2, 2, 3, 3,
	}
	result := ArrayDistinct(abc)
	fmt.Println(fmt.Sprintf("%v", result))
}

func TestArrayToArray(t *testing.T) {
	abc := []int64{
		1, 1, 2, 2, 3, 3,
	}
	result := ArrayToArray(abc, func(item int64) string {
		return fmt.Sprintf("string:%d", item)
	})
	fmt.Println(fmt.Sprintf("%v", result))
}

func TestArrayToMap(t *testing.T) {
	abc := []int64{
		1, 1, 2, 2, 3, 3,
	}
	result := ArrayToMap(abc, func(item int64) (int64, string) {
		return item, fmt.Sprintf("string:%d", item)
	})
	fmt.Println(fmt.Sprintf("%v", result))
}

func TestGetMapKeys(t *testing.T) {
	abc := map[int64]string{
		1: "11", 2: "22", 3: "33",
		11: "11", 22: "22", 33: "33",
	}
	result := GetMapKeys(abc)
	fmt.Println(fmt.Sprintf("%v", result))
}

func TestGetMapValues(t *testing.T) {
	abc := map[int64]string{
		1: "11", 2: "22", 3: "33",
		11: "11", 22: "22", 33: "33",
	}
	result := GetMapValues(abc)
	fmt.Println(fmt.Sprintf("%v", result))
}
