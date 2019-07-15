package iniconfig

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Marshal(data interface{}) (result []byte, err error) {
	return
}
func UnMarshal(data []byte, result interface{}) (err error) {

	linwArr := strings.Split(string(data), "\n")
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}
	var lastFieldName string
	for index, line := range linwArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line[0] == ';' || line[0] == '#' {
			return
		}
		if line[0] == '[' {
			lastFieldName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v lineno:%d", err, index+1)
				return
			}
			continue
		}
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno:%d", err, index+1)
			return
		}
	}
	return
}

func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("sytax error,line;%s", line)
		return
	}
	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index:])
	if len(key) == 0 {
		err = fmt.Errorf("sytax error,line;%s", line)
		return
	}
	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFieldName)
	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", lastFieldName)
		return
	}
	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) == 0 {
		return
	}
	fmt.Println(keyFieldName)
	return
}

func parseSection(line string, typeInfo reflect.Type) (fieldName string, err error) {

	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error,invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error,invalid sevtion:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)+1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error,invalid sevtion:%s", line)
			return
		}
		for i := 0; i < typeInfo.NumField(); i++ {
			field := typeInfo.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				fieldName = field.Name
				break
			}
		}
	}
	return fieldName, err
}
