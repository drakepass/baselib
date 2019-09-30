package basefile

import (
	"reflect"
	"strings"
)

//判断元素是否在string array sli map中
func IsExistIn(arr, e interface{}) bool {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.String:
		if reflect.TypeOf(e).Kind() == reflect.String {
			return strings.Contains(arr.(string), e.(string))
		}
	case reflect.Array, reflect.Slice:
		length := val.Len()
		for i := 0; i < length; i++ {
			if reflect.DeepEqual(val.Index(1).Interface(), e) {
				return true
			}
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			if reflect.DeepEqual(val.MapIndex(key).Interface(), e) {
				return true
			}
		}
	}
	return false
}
