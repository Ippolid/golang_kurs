package main

import (
	"BIGGO/internal/pkg/storage"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var localdir string = Getlocalpath()
var path string = filepath.Join(localdir, "/internal/pkg/storage/data/storageMa.json")

func main() {

	zl, _ := storage.NewStorage()
	closeChan := make(chan struct{})

	content, err := os.ReadFile(path)
	if err != nil {
		os.Create(path)
	} else {
		zl.UnMarshStor(content)
	}

	zl.Set("lk", "1", 0)
	zl.EXPIRE("lk", 3)
	zl.Set("lk1", "1", 0)
	zl.Set("lk2", "1", 4)
	zl.Set("lk3", "1", 3)
	zl.Set("lk4", "1", 2)
	zl.Set("lk5555555", "1", 0)

	go storage.CleaningSession(zl, closeChan, time.Second*10)

	// zl.RPUSH("s", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// zl.RPUSH("K", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// zl.LPOP("s")
	// zl.LPUSH("s", 1, 23, 3)

	// zl.RPOP("K")
	time.Sleep(6 * time.Second)
	// fmt.Println(zl)
	// fmt.Println(zl.Get("lk"))
	time.Sleep(6 * time.Second)
	// fmt.Println(zl.GetKind("lk"))
	//zl.RADDTOSET("s", 2, 3, 56, 56, 10, 9)
	fmt.Println(zl)

	close(closeChan)
	p, _ := zl.MarshStor()

	WriteAtomic(path, p)
	// s := server.New(":8090", store)
	// s.Start()

}
