package storage

import (
	"fmt"
	"testing"
)

type testCase struct {
	name  string
	key   string
	value string
}

func TestSetGet(t *testing.T) {
	cases := []testCase{
		{"string1", "keystring", "vluestring"},
		{"int1", "keyint", "1"},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := s.Get(c.key)

			if *sValue != c.value {
				t.Errorf("values not equal")
			}
		})
	}
}

type testCaseWithKind struct {
	name  string
	key   string
	value any
	kind  string
}

func TestSetGetWithType(t *testing.T) {
	cases := []testCaseWithKind{
		{"string value", "hello", "word", "S"},
		{"int_value", "key", 49857, "D"},
		{"ne_izvestno", "kluch", 45.45, "Нет такого значения"},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)
			fmt.Print(s.GetKind(c.key))

			if s.GetKind(c.key) != c.kind {
				t.Errorf("expected value kind: %v", c.kind)
			}
		})
	}
}
func BenchmarkGet(b *testing.B) {
	cases := []testCase{
		{"string1", "keystring", "vluestring"},
		{"int1", "keyint", "1"},
		{"string1", "keystring11111", "vluestring"},
		{"int12", "keyint1", "1"},
		{"string1", "keystring1", "vluestring"},
		{"int1", "keyint12", "1"},
		{"string1", "keystring12", "vluestring"},
		{"int1", "keyint123", "1"},
		{"string1", "keystring214", "vluestring"},
		{"int1", "keyint389y4", "1"},
		{"string1", "keystring34", "vluestring"},
		{"int1", "keyint8947", "1"},
	}
	s, _ := NewStorage()

	for i := 0; i < 12; i++ {
		s.Set(cases[i].key, cases[i].value)
	}

	b.ResetTimer()

	for i := 0; i < 12; i++ {
		s.Get(cases[i].key)
	}
}
func BenchmarkSet(b *testing.B) {
	cases := []testCase{
		{"string1", "keystring", "vluestring"},
		{"int1", "keyint", "1"},
		{"string1", "keystring11111", "vluestring"},
		{"int12", "keyint1", "1"},
		{"string1", "keystring1", "vluestring"},
		{"int1", "keyint12", "1"},
		{"string1", "keystring12", "vluestring"},
		{"int1", "keyint123", "1"},
		{"string1", "keystring214", "vluestring"},
		{"int1", "keyint389y4", "1"},
		{"string1", "keystring34", "vluestring"},
		{"int1", "keyint8947", "1"},
	}
	s, _ := NewStorage()
	b.ResetTimer()

	for i := 0; i < 12; i++ {
		s.Set(cases[i].key, cases[i].value)
	}
}
func BenchmarkSetGet(b *testing.B) {
	cases := []testCase{
		{"string1", "keystring", "vluestring"},
		{"int1", "keyint", "1"},
		{"string1", "keystring11111", "vluestring"},
		{"int12", "keyint1", "1"},
		{"string1", "keystring1", "vluestring"},
		{"int1", "keyint12", "1"},
		{"string1", "keystring12", "vluestring"},
		{"int1", "keyint123", "1"},
		{"string1", "keystring214", "vluestring"},
		{"int1", "keyint389y4", "1"},
		{"string1", "keystring34", "vluestring"},
		{"int1", "keyint8947", "1"},
	}
	s, _ := NewStorage()
	b.ResetTimer()

	for i := 0; i < 12; i++ {
		s.Set(cases[i].key, cases[i].value)
	}
	for i := 0; i < 12; i++ {
		s.Get(cases[i].key)
	}
}
