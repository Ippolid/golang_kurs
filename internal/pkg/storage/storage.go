package storage

import (
	"strconv"
	"sync"
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
	mu    sync.RWMutex
}

func NewStorage() (Storage, error) {
	return Storage{
		inner: make(map[string]Value),
	}, nil
}

func (r *Storage) Set(key string, value string, exp int64) {
	r.mu.RLock()
	defer r.mu.Unlock()

	var z int64
	if exp == 0 {
		z = 0
	} else {
		ttl := time.Duration(exp) * time.Second
		z = time.Now().Add(ttl).UnixMilli()
	}

	kind := getType(value)
	r.inner[key] = Value{s: value, k: kind, exp: z}

}

func (r *Storage) Get(key string) *string {
	r.mu.Lock()
	defer r.mu.Unlock()

	res, ok := r.inner[key]
	var k *string
	if !ok {
		return k
	} else if time.Now().UnixMilli() >= res.exp && res.exp != 0 {
		return k
	}
	return &res.s
}
func (r *Storage) GetKind(key string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	res, ok := r.inner[key]
	if !ok {
		return "No value"
	} else if time.Now().UnixMilli() >= res.exp && res.exp != 0 {
		r.DeleteElem(key)
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

func (r *Storage) EXPIRE(key string, sec int) {
	r.mu.RLock()
	defer r.mu.Unlock()

	res, ok := r.inner[key]
	if !ok {
		return
	}
	ttl := time.Duration(sec) * time.Second
	z := time.Now().Add(ttl).UnixMilli()
	res.exp = z
	r.inner[key] = res

}

func (r *Storage) DeleteElem(key string) {
	r.mu.RLock()
	defer r.mu.Unlock()

	delete(r.inner, key)
}
