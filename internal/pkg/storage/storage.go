package storage

import "strconv"

type Value struct {
	s string
	d int
}
type Storage struct {
	String map[string]Value
}

func NewStorage() (Storage, error) {
	return Storage{
		make(map[string]Value),
	}, nil
}

func (r Storage) Set(key string, value any) {
	switch value.(type) {
	case int:
		l := Value{s: "", d: value.(int)}
		r.String[key] = l
	case string:
		l := Value{s: value.(string), d: 0}
		r.String[key] = l
	}

}

func (r Storage) Get(key string) *string {
	res, ok := r.String[key]
	if !ok {
		return nil
	} else {
		if res.d != 0 {
			x := strconv.Itoa(res.d)
			return &x
		} else {
			x := res.s
			return &x
		}
	}
}
func (r Storage) GetKind(key string) string {
	res, ok := r.String[key]
	if !ok {
		return "Нет такого значения"
	} else {
		if res.d != 0 {
			return "D"
		} else {
			return "S"
		}
	}
}
