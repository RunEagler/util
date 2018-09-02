package util

import (
	"reflect"
	"sort"
	"time"
)

func MapStringIntSort(mapStringInt map[string]int) []string {

	keys := MapStringKeys(Value(mapStringInt))
	sort.Slice(keys, func(a, b int) bool {
		return mapStringInt[keys[a]] < mapStringInt[keys[b]]
	})
	return keys
}

func MapStringTimeSort(mapStringTime map[string]time.Time) []string {

	keys := MapStringKeys(Value(mapStringTime))
	sort.Slice(keys, func(a, b int) bool {
		return mapStringTime[keys[a]].Before(mapStringTime[keys[b]])
	})
	return keys
}

func MapStringKeys(hash reflect.Value) []string {
	var keys []string
	keys = []string{}
	if !hash.IsValid() {
		return keys
	}
	hash = PtrElem(hash)
	if hash.Type().Kind() == reflect.Map {
		for _, value := range hash.MapKeys() {
			if value.Type().Kind() == reflect.String {
				keys = append(keys, value.Interface().(string))
			}
		}
		sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
	}
	return keys
}

func MapIntKeys(hash reflect.Value) []int {
	var keys []int
	keys = []int{}
	if !hash.IsValid() {
		return keys
	}
	hash = PtrElem(hash)
	if hash.Type().Kind() == reflect.Map {
		for _, value := range hash.MapKeys() {
			if value.Type().Kind() == reflect.Int {
				keys = append(keys, value.Interface().(int))
			}
		}
		sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
	}
	return keys
}

func MapTimeKeys(hash reflect.Value) []time.Time {
	var keys []time.Time
	keys = []time.Time{}

	if !hash.IsValid() {
		return keys
	}
	hash = PtrElem(hash)
	if hash.Type().Kind() == reflect.Map {
		for _, value := range hash.MapKeys() {
			switch t := value.Interface().(type) {
			case time.Time:
				keys = append(keys, t)
			default:

			}
		}
		sort.Slice(keys, func(a, b int) bool { return keys[a].Before(keys[b]) })
	}
	return keys
}
