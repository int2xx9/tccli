package formatter

import (
	"html/template"
	"reflect"
	"strings"

	"github.com/int2xx9/tccli/teamcity/schema"
	"github.com/moznion/go-optional"
)

type CsvFormatter struct {
	WriteHeader bool
	items       []csvFormatterItem
}

func NewCsvFormatter(formatString string, writeHeader bool) (*CsvFormatter, error) {
	items, err := parseCsvFormatString(formatString)
	if err != nil {
		return nil, err
	}

	return &CsvFormatter{
		WriteHeader: writeHeader,
		items:       items,
	}, nil
}

type csvFormatterItem struct {
	Path     string
	Header   string
	Template *template.Template
}

func parseCsvFormatString(str string) ([]csvFormatterItem, error) {
	items := []csvFormatterItem{}
	for _, item := range strings.Split(str, ",") {
		parts := strings.SplitN(item, "|", 2)
		i := csvFormatterItem{
			Path:   parts[0],
			Header: parts[0],
		}
		if len(parts) > 1 {
			i.Header = parts[1]
		}

		i.Template = template.New(i.Path)

		i.Template.Funcs(template.FuncMap{
			"formatTime": func(v any, format string) string {
				switch vv := v.(type) {
				case schema.TeamcityTime:
					return vv.Time().Format(format)
				case optional.Option[schema.TeamcityTime]:
					if vvv, err := vv.Take(); err != nil {
						return "-"
					} else {
						return vvv.Time().Format(format)
					}
				default:
					return "-"
				}
			},
		})

		_, err := i.Template.Parse(i.Path)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}
	return items, nil
}

func escape(s string) string {
	if !strings.ContainsAny(s, "\"\n,") {
		return s
	}
	return "\"" + strings.ReplaceAll(strings.ReplaceAll(s, "\\", "\\\\"), "\"", "\\\"") + "\""
}

func (f *CsvFormatter) formatSingle(v any) (string, error) {
	items := []string{}
	for _, item := range f.items {
		b := strings.Builder{}

		if err := item.Template.Execute(&b, v); err != nil {
			b.WriteString("(?)")
		}

		items = append(items, escape(b.String()))
	}

	return strings.Join(items, ","), nil
}

func (f *CsvFormatter) Format(v any) (string, error) {
	fieldNames := []string{}
	for _, item := range f.items {
		fieldNames = append(fieldNames, item.Header)
	}

	output := ""
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		if f.WriteHeader {
			output = strings.Join(fieldNames, ",") + "\n"
		}

		line, err := f.formatSingle(v)
		if err != nil {
			return "", err
		}
		output += line + "\n"
		return output, nil
	} else {
		if f.WriteHeader {
			output = strings.Join(fieldNames, ",") + "\n"
		}

		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i).Interface()
			line, err := f.formatSingle(elem)
			if err != nil {
				return "", err
			}
			output += line + "\n"
		}
		return output, nil
	}
}
