package util

import (
	"reflect"
)




func HasKey(hash interface{}, targetKey interface{}) (bool) {

	switch reflect.TypeOf(hash).Kind() {
	case reflect.Map:

		switch reflect.TypeOf(hash).Key().Kind() {
		case reflect.String:
			if !IsString(targetKey) {
				//return false, fmt.Errorf("type of hash key should be same type of targetKey")
				return false
			}
			return hasKeyString(hash.(map[string]interface{}), targetKey.(string))
		case reflect.Int:
			if !IsInt(targetKey) {
				//return false, fmt.Errorf("type of hash key should be same type of targetKey")
				return false
			}
			return hasKeyInt(hash.(map[int]interface{}), targetKey.(int))
		}
	default:
		//return false, fmt.Errorf("hash should be map")
		return false
	}
	return false
}

func hasKeyString(hash map[string]interface{}, key string) bool {

	if _, exist := hash[key]; exist {
		return true
	}
	return false
}

func hasKeyInt(hash map[int]interface{}, key int) bool {

	if _, exist := hash[key]; exist {
		return true
	}
	return false
	reflect.DeepEqual()
}


