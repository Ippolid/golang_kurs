package storage

import "strconv"

type Value struct {
	s string
	k string
}

const (
	KindInt       string = "D"
	KindString    string = "S"
	KindUndefined string = "UN"
)

type Storage struct {
	inner map[string]Value
}

func NewStorage() (Storage, error) {
	return Storage{
		make(map[string]Value),
	}, nil
}

func (r Storage) Set(key string, value string) {
	switch kind := getType(value); kind {
	case KindInt:
		r.inner[key] = Value{s: value, k: kind}
	case KindString:
		r.inner[key] = Value{s: value, k: kind}
	case KindUndefined:
		r.inner[key] = Value{s: value, k: kind}
	}

}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	if !ok {
		return nil
	}
	return &res.s
}
func (r Storage) GetKind(key string) string {
	return r.inner[key].k
}

func getType(value string) string {
	var val any

	val, err := strconv.Atoi(value)
	if err != nil {
		val = value
	}

	switch val.(type) {
	case int:
		return KindInt
	case string:
		return KindString
	default:
		return KindUndefined
	}
}
