package formatter

import "encoding/json"

type JsonFormatter struct{}

func (f *JsonFormatter) Format(v any) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
