package json_path_scanner

import (
	"strconv"
	"strings"
)

type PathValue struct {
	Path  string
	Value interface{}
}

func newPathValue(key string, value interface{}) *PathValue {
	return &PathValue{
		Path:  key,
		Value: value,
	}
}

func Scan(value interface{}, ch chan<- *PathValue) {
	defer close(ch)
	scanJson("$", value, ch)
}

func scanJson(label string, value interface{}, ch chan<- *PathValue) {
	switch value.(type) {
	case int, float64, string, nil:
		ch <- newPathValue(label, value)
	case map[string]interface{}:
		m := value.(map[string]interface{})
		for k, v := range m {
			if strings.Contains(k, ".") {
				scanJson(label+"['"+k+"']", v, ch)
			} else {
				scanJson(label+"."+k, v, ch)
			}
		}
	case []interface{}:
		for i, v := range value.([]interface{}) {
			scanJson(label+"["+strconv.Itoa(i)+"]", v, ch)
		}
	default:
		panic("Unsupported type in json")
	}
}
