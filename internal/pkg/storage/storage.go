package storage

import (
	"strconv"
	"time"
)

type Value struct {
	s   string
	k   string
	exp int64
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

func (r Storage) Set(key string, value string, exp int64) {
	var z int64
	if exp == 0 {
		z = 0
	} else {
		ttl := time.Duration(exp) * time.Second
		z = time.Now().Add(ttl).UnixMilli()
	}

	switch kind := getType(value); kind {
	case KindInt:
		r.inner[key] = Value{s: value, k: kind, exp: z}
	case KindString:
		r.inner[key] = Value{s: value, k: kind, exp: z}
	case KindUndefined:
		r.inner[key] = Value{s: value, k: kind, exp: z}
	}

}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	var k *string = nil
	if !ok {
		return k
	} else if time.Now().UnixMilli() >= res.exp && res.exp != 0 {
		return k
	}
	return &res.s
}
func (r Storage) GetKind(key string) string {
	res, ok := r.inner[key]
	if !ok {
		return "No value"
	} else if time.Now().UnixMilli() >= res.exp && res.exp != 0 {
		return "expired"
	}
	return res.k
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

func (r Storage) EXPIRE(key string, sec int) {
	res, ok := r.inner[key]
	if !ok {
		return
	}
	ttl := time.Duration(sec) * time.Second
	z := time.Now().Add(ttl).UnixMilli()
	res.exp = z
	r.inner[key] = res

}
