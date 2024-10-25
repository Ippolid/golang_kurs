package main

import (
	"os"
	"path/filepath"
)

func Getlocalpath() string {
	mydir, _ := os.Getwd()
	k := filepath.Dir(mydir)
	return k
}
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
