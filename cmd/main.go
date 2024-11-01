package main

import (
	"BIGGO/internal/pkg/storage"
	"fmt"
	"time"
)

// import (
// 	"BIGGO/internal/pkg/storage"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// var localdir string = Getlocalpath()
// var path string = filepath.Join(localdir, "/internal/pkg/storage/data/storageMa.json")

func main() {
	closeChan := make(chan struct{})

	//zl, _ := storage.NewStorageMa()
	k, _ := storage.NewStorage()
	k.Set("lk", "1", 0)
	//k.EXPIRE("lk", 3)
	k.Set("lk1", "1", 0)
	k.Set("lk2", "1", 4)
	k.Set("lk3", "1", 3)
	k.Set("lk4", "1", 2)
	k.Set("lk5", "1", 1)

	go storage.CleaningSession(k, closeChan, time.Second*10)
	// content, err := os.ReadFile(path)
	// if err != nil {
	// 	os.Create(path)
	// } else {
	// 	zl.UnMarshStor(content)
	// }

	// zl.RPUSH("s", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// zl.RPUSH("K", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// zl.LPOP("s")
	// zl.LPUSH("s", 1, 23, 3)

	// zl.RPOP("K")
	time.Sleep(6 * time.Second)
	fmt.Println(k)
	fmt.Println(k.Get("lk"))
	time.Sleep(6 * time.Second)
	fmt.Println(k.GetKind("lk"))
	//zl.RADDTOSET("s", 2, 3, 56, 56, 10, 9)
	fmt.Println(k)

	close(closeChan)
	//p, _ := zl.MarshStor()

	// WriteAtomic(path, p)
	// store, err := storage.NewStorage()
	// if err != nil {
	// 	panic(err)
	// }
	// s := server.New(":8090", store)
	// s.Start()

}
