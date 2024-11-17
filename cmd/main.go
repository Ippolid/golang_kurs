package main

import (
	//"BIGGO/internal/pkg/server"
	"BIGGO/internal/pkg/storage"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

// import (
// 	"BIGGO/internal/pkg/storage"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

var (
	localdir string = GetLocalPath()
	path     string = filepath.Join(localdir, "/internal/pkg/storage/data/storageMa.json")
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Канал завершения
	done := make(chan bool, 1)
	// Горутина для обработки сигнала
	go func() {
		sig := <-signalChan
		fmt.Println("\nПолучен сигнал:", sig)
		done <- true
	}()

	k, _ := storage.NewStorage()
	closeChan := make(chan struct{})
	go storage.CleaningSession(k, closeChan, time.Second*10)
	content, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, createErr := os.Create(path)
			if createErr != nil {
				fmt.Printf("Ошибка при создании файла: %v\n", createErr)
				return
			}
			defer file.Close()
		} else {
			fmt.Printf("Ошибка при чтении файла: %v\n", err)
			return
		}
	} else {
		k.UnMarshStor(content)
	}

	fmt.Println(k)

	k.Set("lk", "1", 0)
	k.Set("lk1", "1", 0)
	fmt.Println(k)
	k.Set("lk2", "1", 4)
	k.Set("lk3", "1", 3)
	k.Set("lk4", "1", 2)
	k.Set("lk5", "1", 1)

	// zl.RPUSH("s", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// zl.RPUSH("K", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// zl.LPOP("s")
	// zl.LPUSH("s", 1, 23, 3)

	// zl.RPOP("K")
	time.Sleep(6 * time.Second)
	//fmt.Println(k)
	fmt.Println("sfdf")
	fmt.Println(k.Get("lk"))
	time.Sleep(6 * time.Second)
	fmt.Println(k.GetKind("lk"))
	fmt.Println(k.GetKind("lk"))
	fmt.Println(k)
	//zl.RADDTOSET("s", 2, 3, 56, 56, 10, 9)
	fmt.Println(k.Get("lk1"))

	close(closeChan)
	p, _ := k.MarshStor()

	// store, err := storage.NewStorage()
	// if err != nil {
	// 	panic(err)
	// }
	// s := server.New(":8090", store)
	// s.Start()
	fmt.Println("Приложение .")
	<-done
	close(done)
	WriteAtomic(path, p)

	fmt.Println("Приложение завершено.")

}
