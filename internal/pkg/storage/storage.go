package storage

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Value struct {
	S   string
	K   string
	Exp int64
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

func NewStorage() (*Storage, error) {
	return &Storage{
		inner: make(map[string]Value),
	}, nil
}

func (r *Storage) Set(key string, value string, exp int64) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var z int64
	if exp == 0 {
		z = 0
	} else {
		ttl := time.Duration(exp) * time.Second
		z = time.Now().Add(ttl).UnixMilli()
	}

	kind := getType(value)
	r.inner[key] = Value{S: value, K: kind, Exp: z}

}

func (r *Storage) Get(key string) *string {
	r.mu.Lock()
	defer r.mu.Unlock()

	res, ok := r.inner[key]
	var k *string
	if !ok {
		return k
	} else if time.Now().UnixMilli() >= res.Exp && res.Exp != 0 {
		return k
	}
	return &res.S
}
func (r *Storage) GetKind(key string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	res, ok := r.inner[key]
	if !ok {
		return "No value"
	} else if time.Now().UnixMilli() >= res.Exp && res.Exp != 0 {
		r.DeleteElem(key)
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

func (r *Storage) EXPIRE(key string, sec int) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	res, ok := r.inner[key]
	if !ok {
		return
	}
	ttl := time.Duration(sec) * time.Second
	z := time.Now().Add(ttl).UnixMilli()
	r.inner[key] = Value{S: res.S, K: res.K, Exp: z}

}

func (r *Storage) DeleteElem(key string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	delete(r.inner, key)
}

func (r *Storage) MarshStor() ([]byte, error) {
	jsonInfo, err := json.Marshal(r.inner)

	if err != nil {
		return nil, fmt.Errorf("write error: %w", err)
	}

	return jsonInfo, err
}

func (r *Storage) UnMarshStor(z []byte) error {
	err := json.Unmarshal([]byte(z), &r.inner)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	return nil
}
