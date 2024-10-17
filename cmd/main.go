package main

import (
	"BIGGO/internal/pkg/storage"
	"fmt"
	"os"
	"path/filepath"
	//"encoding/json"
)

func WriteAtomic(path string, b []byte) error {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	tmpPathName := filepath.Join(dir, filename+"tmp")

	err := os.WriteFile(tmpPathName, b, 0777)
	if err != nil {
		return err
	}

	defer func() {
		os.Remove(tmpPathName)
	}()

	return os.Rename(tmpPathName, path)

}

var path string = "/home/ippolid/Desktop/BIGGO/internal/pkg/storage/data/storage_ma.json"

func main() {
	zl, _ := storage.NewStorageMa()

	content, err := os.ReadFile(path)
	if err == nil {
		zl.UnMarshStor(content)
	}

	zl.RPUSH("s", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	zl.RPUSH("K", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	zl.LPOP("s")

	zl.RPOP("K")
	fmt.Println(zl)

	WriteAtomic(path, zl.MarshStor())

}

//err := pl.LSET("s", 6, 23)
//fmt.Print(err, pl)
