package fetcher

import "github.com/jinzhu/copier"

func DeepCopy[T any](t *T) T {
	var copy T
	copier.Copy(&copy, t)
	return copy
}
