package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
)

//import "fmt"

type Storage_ma struct {
	inner map[string][]int
}

func NewStorageMa() (Storage_ma, error) {
	return Storage_ma{
		make(map[string][]int),
	}, nil
}
func (r Storage_ma) MarshStor() []byte {
	jsonInfo, err := json.Marshal(&r.inner)
	if err != nil {
		fmt.Println("Ошибка записи данных:", err)
	}
	return jsonInfo
}

func (r *Storage_ma) UnMarshStor(z []byte) {
	err := json.Unmarshal([]byte(z), &r.inner)
	if err != nil {
		fmt.Println("Ошибка чтения JSON-данных:", err)
	}
}
func (s Storage_ma) setma(key string) {
	s.inner[key] = make([]int, 0)
}

func (s Storage_ma) isma(key string) bool {
	_, ok := s.inner[key]
	if !ok {
		return false
	}
	return true
}
func (s Storage_ma) LPUSH(key string, element ...int) {
	if !(s.isma(key)) {
		s.setma(key)
	}
	for _, i := range element {
		s.inner[key] = append(s.inner[key], i)
	}
}
func (s Storage_ma) RPUSH(key string, element ...int) {
	if !(s.isma(key)) {
		s.setma(key)
	}
	s.inner[key] = append(element, s.inner[key]...)
}

func (s Storage_ma) RADDTOSET(key string, element ...int) {
	if !(s.isma(key)) {
		s.setma(key)
	}
	k := 0
	for _, i := range element {
		k = 0
		for _, l := range s.inner[key] {
			if i == l {
				k++
				break
			}
		}
		if k == 0 {
			s.inner[key] = append(s.inner[key], i)
		}
	}
}

func (s Storage_ma) LPOP(key string, element ...int) []int {
	a := make([]int, 0)
	if len(element) == 0 {
		a = append(a, s.inner[key][0])
		s.inner[key] = s.inner[key][1:]
	} else if len(element) == 1 {
		if element[0] < 0 {
			element[0] = len(s.inner[key]) + element[0] + 1
		}
		a = append(a, s.inner[key][0:min(element[0], len(s.inner[key]))]...)
		s.inner[key] = s.inner[key][min(element[0], len(s.inner[key])):]
	} else {
		if element[0] < 0 {
			element[0] = len(s.inner[key]) + element[0] + 1
		}
		if element[1] < 0 {
			element[1] = len(s.inner[key]) + element[1] + 1
		}
		a = append(a, s.inner[key][element[0]:element[1]]...)
		s.inner[key] = append(s.inner[key][:element[0]], s.inner[key][element[1]:]...)
	}
	return a

}

func (s Storage_ma) RPOP(key string, element ...int) []int {
	slices.Reverse(s.inner[key])
	v := []int{}
	if len(element) == 0 || len(element) == 1 {
		v = s.LPOP(key, element...)
		slices.Reverse(s.inner[key])
	} else {
		v = s.LPOP(key, element[0]-1, element[1]-1)
		slices.Reverse(s.inner[key])
	}
	return v
}

func (s Storage_ma) LSET(key string, index, element int) error {
	if index < 0 {
		index = len(s.inner[key]) + index
	}
	if index < len(s.inner[key]) {
		s.inner[key][index] = element
		return nil
	}
	l := errors.New("Index out of range")
	return l
}
func (s Storage_ma) LGET(key string, index int) (int, error) {
	if index < 0 {
		index = len(s.inner[key]) + index
	}
	if index < len(s.inner[key]) {
		return s.inner[key][index], nil
	}
	l := errors.New("Index out of range")
	return -1000000000000000000, l
}
