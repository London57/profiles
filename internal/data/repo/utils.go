package repo

import "strings"

func FieldsToExexString(fields map[string]any) (string, []string, []any) {
	values := make([]any, len(fields))
	keys := make([]string, len(fields))
	var result_string strings.Builder

	for k, v := range fields {
		values = append(values, v)
		keys = append(keys, k)
		result_string.WriteString(k)
		result_string.WriteByte('=')
		result_string.WriteByte('?')
		result_string.WriteString(", ")
	}

	s := result_string.String()
	result_string.Reset()
	result_string.WriteString(s[:len(s)-2])
	return result_string.String(), keys, values
}