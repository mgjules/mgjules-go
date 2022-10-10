package mapstruct

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func FromSlice[T any](ss []T) []map[string]any {
	var res []map[string]any
	for _, s := range ss {
		res = append(res, FromSingle(s))
	}

	return res
}

func FromSingle[T any](s T) map[string]any {
	var m map[string]any
	data, _ := json.Marshal(s)
	json.Unmarshal(data, &m)
	return m
}
