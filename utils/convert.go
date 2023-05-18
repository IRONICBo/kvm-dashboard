package utils

import (
	"encoding/json"
	"errors"
	"reflect"
)

// flatten map[string]interface{} to map[string]xxx, with only one level
// {
// "sensor": {
// 	"temperature": 25,
// 	"humidity": 50
// }
// TO
// }
// {
// "sensor+temperature": 25,
// "sensor+humidity": 50
// }
func FlattenMap(input map[string]interface{}, infix string) map[string]interface{} {
	flattened := make(map[string]interface{})
	for k, v := range input {
		switch v := v.(type) {
		case map[string]interface{}:
			dfsMap(flattened, v, k, infix)
		default:
			flattened[k] = v
		}
	}
	return flattened
}

func dfsMap(flattened, data map[string]interface{}, path string, infix string) {
	for k, v := range data {
		newPath := path + infix + k
		switch v := v.(type) {
		case map[string]interface{}:
			dfsMap(flattened, v, newPath, infix)
		default:
			flattened[newPath] = v
		}
	}
}

func StructToMapWithJSONTag(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, errors.New("data parameter must be a struct or a pointer to struct")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// use json tag to convert struct to map recursively
// deprecated
func _StructToMapWithJSONTag(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, errors.New("data parameter must be a struct or a pointer to struct")
	}

	valueType := value.Type()
	numFields := value.NumField()

	for i := 0; i < numFields; i++ {
		field := valueType.Field(i)
		fieldValue := value.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			if fieldValue.Kind() == reflect.Struct {
				nestedStructMap, err := _StructToMapWithJSONTag(fieldValue.Interface())
				if err != nil {
					return nil, err
				}
				result[jsonTag] = nestedStructMap
			} else {
				if fieldValue.CanInterface() {

					// judge is struct or pointer to struct
					tmp := reflect.ValueOf(fieldValue.Interface())
					if tmp.Kind() == reflect.Ptr {
						tmp = tmp.Elem()
					}
					if tmp.Kind() == reflect.Struct {
						nestedStructMap, err := _StructToMapWithJSONTag(tmp.Interface())
						if err != nil {
							return nil, err
						}
						result[jsonTag] = nestedStructMap
					}

					result[jsonTag] = fieldValue.Interface()
				} else {
					result[jsonTag] = fieldValue
				}
			}
		}
	}

	return result, nil
}
