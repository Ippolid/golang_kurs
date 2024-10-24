package main

import (
	"BIGGO/internal/pkg/storage"
	"fmt"
	"os"
	"path/filepath"
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

var localdir string = Getlocalpath()
var path string = filepath.Join(localdir, "/internal/pkg/storage/data/storageMa.json")

func main() {
	zl, _ := storage.NewStorageMa()

	content, err := os.ReadFile(path)
	if err != nil {
		os.Create(path)
	} else {
		zl.UnMarshStor(content)
	}

	zl.RPUSH("s", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	zl.RPUSH("K", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	zl.LPOP("s")
	zl.LPUSH("s", 1, 23, 3)

	zl.RPOP("K")
	fmt.Println(zl)
	zl.RADDTOSET("s", 2, 3, 56, 56, 10, 9)
	fmt.Println(zl)

	p, _ := zl.MarshStor()

	WriteAtomic(path, p)

}
