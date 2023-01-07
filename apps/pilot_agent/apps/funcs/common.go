package funcs

import (
	"bytes"
	"encoding/json"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// "not used" 오류 회피
func UNUSED(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

// 값을 가져올때 NULL 검사하여 가져옴
func GetNC(m interface{}) interface{} {
	if m == nil {
		return ""
	}
	return m
}

// if then else 문법 처리
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// 문자열에서 CRLF 및 모든 공백 제거
func OutputOmit(data string) string {
	re := regexp.MustCompile(`\r?\n|\s+`)
	return re.ReplaceAllString(data, "")
}

// Json 처리 자료형
type JSON_DATA map[string]interface{}

// Json 문자열 출력
func JsonDumpsIndent(data interface{}, indent string) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", indent)

	err := encoder.Encode(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	return strings.TrimRight(buffer.String(), "\n")
}

// Json 문자열 출력
func JsonDumps(data interface{}) string {
	return JsonDumpsIndent(data, "  ")
}

// Json 특수문자 Escape
func JsonEscape(value string) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(value)

	return buffer.String()
}

func JsonValue(json_data JSON_DATA, keys ...string) (interface{}, string) {
	size := len(keys)

	if 1 == size {
		if val, ok := json_data[keys[0]]; ok {
			return val, reflect.TypeOf(val).String()
		}
	} else if 2 <= size {
		if val, ok := json_data[keys[0]]; ok {
			return JsonValue(val.(map[string]interface{}), keys[1:]...)
		}
	}

	return nil, "nil"
}
