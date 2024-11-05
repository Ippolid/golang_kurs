package storage

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Value struct {
	S   string
	K   string
	EXP int64
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
		r.inner[key] = Value{S: value, K: kind, EXP: z}
	case KindString:
		r.inner[key] = Value{S: value, K: kind, EXP: z}
	case KindUndefined:
		r.inner[key] = Value{S: value, K: kind, EXP: z}
	}

}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	var k *string = nil
	if !ok {
		return k
	} else if time.Now().UnixMilli() >= res.EXP && res.EXP != 0 {
		return k
	}
	return &res.S
}
func (r Storage) GetKind(key string) string {
	res, ok := r.inner[key]
	if !ok {
		return "No value"
	} else if time.Now().UnixMilli() >= res.EXP && res.EXP != 0 {
		return "expired"
	}
	return res.K
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
	res.EXP = z
	r.inner[key] = res

}

func (s Storage) MarshStor() ([]byte, error) {

	jsonInfo, err := json.Marshal(s.inner)

	if err != nil {
		fmt.Println("Ошибка записи данных:", err)
	}

	return jsonInfo, err
}

func (s *Storage) UnMarshStor(z []byte) {
	err := json.Unmarshal([]byte(z), &s.inner)
	if err != nil {
		fmt.Println("Ошибка чтения JSON-данных:", err)
	}
}
