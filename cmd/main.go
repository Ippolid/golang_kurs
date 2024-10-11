package main

import (
	"BIGGO/internal/pkg/storage"
	"fmt"
)

func main() {
	r, _ := storage.NewStorage()
	r.Set("test1", "test123")
	r.Set("inttest1", 228)
	val1 := *r.Get("test1")
	val2 := *r.Get("inttest1")
	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(r.GetKind("test1"))
	fmt.Println(r.GetKind("inttest1"))

}
