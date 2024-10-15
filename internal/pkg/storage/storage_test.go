package storage

import "testing"

type testCase struct {
	name  string
	key   string
	value string
}

func TestSetGet(t *testing.T) {
	cases := []testCase{
		{"test1", "kluch", "string"},
		{"test2", "kluchik", "1234"},
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
				t.Errorf("Значения не равны")
			}
		})
	}
}

type testKind struct {
	name  string
	key   string
	value string
	kind  string
}

func TestSetGetWithType(t *testing.T) {
	cases := []testKind{
		{"stringtest", "key1", "stroka", KindString},
		{"inttest", "key2", "28282", KindInt},
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
				t.Errorf("Значения не равны")
			}

			if getType(*sValue) != getType(c.value) {
				t.Errorf("Типы не равны")
			}

			if getType(*sValue) != c.kind {
				t.Errorf("предпологаемое значение: %v", c.kind)
			}
		})
	}
}

// Бенчмарки
func BenchmarkGet(b *testing.B) {
	s, _ := NewStorage()
	k := []string{"a", "b", "c", "d", "e", "f", "g", "h", "k"}
	for _, c := range k {
		s.Set(c, c)
	}

	for _, c := range k {
		b.ResetTimer()
		s.Get(c)
	}
}
func BenchmarkSet(b *testing.B) {
	s, _ := NewStorage()
	k := []string{"a", "b", "c", "d", "e", "f", "g", "h", "k"}
	for _, c := range k {
		b.ResetTimer()
		s.Set(c, c)
	}
}
func BenchmarkSetGET(b *testing.B) {
	s, _ := NewStorage()
	k := []string{"a", "b", "c", "d", "e", "f", "g", "h", "k"}
	for _, c := range k {
		b.ResetTimer()
		s.Set(c, c)
		s.Get(c)
	}
}
