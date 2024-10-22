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
