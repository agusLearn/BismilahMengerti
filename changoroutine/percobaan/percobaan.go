package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type adder interface {
	AddCounter(int64)
}

type counter struct {
	counter   int64
	goCounter uint64
	called    uint64
}

func (c *counter) AddCounter(inc int64) {
	atomic.AddUint64(&c.called, 1)  // nilai dari called - 1
	pc := make([]uintptr, 1)        // membuat slice dengan panjang kapasitas 1
	count := runtime.Callers(5, pc) // nggk tau ini apa

	if count == 0 {
		atomic.AddUint64(&c.goCounter, 1)
	}
	// jika count sama dengan 0. maka menjalankan proses atomic
	// nilai dari goCounter - 1
}

func ChanGoroutine(intChan chan<- int, add adder) {

	go func() {
		add.AddCounter(64) 
	}()

	allSum := 0
	// for z := range intChan {

	// }
	intChan <- allSum
}

func main() {
	c := new(counter)       // ini instansiasi struct counter tanpa ada isi
	ch := make(chan int, 5) // ini membuat channels

	c.counter = 45      // datanya ngasal, gw cuman lihat yang ada di changeroutine_test.go
	c.goCounter = 10000 // ini juga sama

	fmt.Printf("variable CH ini memiliki tipe data %T \n", ch)
	fmt.Printf("variable C ini memiliki tipe data %T \n", c.counter)
	fmt.Println("Hasil dari c.counter >= 45 : ", c.counter >= 45)


}
