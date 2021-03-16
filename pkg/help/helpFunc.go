package help

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)


func NowTime(str string) string {
	str = strings.Replace(str, "Y", "2006", -1)
	str = strings.Replace(str, "y", "06", -1)
	str = strings.Replace(str, "m", "01", -1)
	str = strings.Replace(str, "d", "02", -1)
	str = strings.Replace(str, "h", "15", -1)
	str = strings.Replace(str, "i", "04", -1)
	str = strings.Replace(str, "s", "05", -1)
	return time.Now().Format(str)
}

func StructToMap(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	byteStr, _ := json.Marshal(obj)
	json.Unmarshal(byteStr, &data)
	return data
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
func P(v ...interface{}) {
	dateTime := NowTime("Y-m-d h:i:s")
	for k, m := range v {
		typeOf:=reflect.TypeOf(m)
		fmt.Println(dateTime,k,typeOf,m)
	}
	fmt.Println("")
}
func Pmap(m map[string]interface{}) {
	for k, v := range m {
		switch value := v.(type) {
		case nil:
			fmt.Println(k, "is nil", "null")
		case string:
			fmt.Println(k, "is string", value)
		case int:
			fmt.Println(k, "is int", value)
		case float64:
			fmt.Println(k, "is float64", value)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range value {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			Pmap(value)
		default:
			fmt.Println(k, "is unknown type", fmt.Sprintf("%T", v))
		}
	}
	fmt.Println("")
}

