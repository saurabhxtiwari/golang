// Mutex
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func readFromMap(m map[int]int, key int, mutex *sync.Mutex) {
	go func() {
		mutex.Lock()
		x := m[key]
		fmt.Println("Read ", x)
		mutex.Unlock()
	}()
}

func writeToMap(m map[int]int, key int, value int, mutex *sync.Mutex) {
	go func() {
		mutex.Lock()
		m[key] = value
		fmt.Println("Wrote ", value)
		mutex.Unlock()
	}()
}

func main() {

	m := make(map[int]int)
	mutex := sync.Mutex{}

	go func() {
		for i := 0; i < 20; i++ {
			value := rand.Intn(20)
			key := rand.Intn(5)
			writeToMap(m, key, value, &mutex)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			key := rand.Intn(5)
			readFromMap(m, key, &mutex)
		}
	}()

	time.Sleep(time.Second * 5)
}
