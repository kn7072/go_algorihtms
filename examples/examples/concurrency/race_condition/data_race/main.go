package main

import (
	"fmt"
	"sync"
)
//https://ubiklab.net/posts/race-condition-and-data-race/

/*
Data Race

    Data race это состояние когда разные потоки обращаются к одной ячейке памяти без какой-либо синхронизации и как минимум один из потоков осуществляет запись.
*/

type account struct {
	balance int
}

/*
Причина в том, что операция acc.balance += amount не атомарная. Она может разложиться на 3:
tmp := acc.balance
tmp = tmp + amount
acc.balance = tmp

*/

func deposit(acc *account, amount int) {
	acc.balance += amount
}

/*
Решается проблема с помощью синхронизации:

var mu sync.Mutex

func deposit(acc *account, amount int) {
   mu.Lock()
   acc.balance += amount
   mu.Unlock()
}

Иногда более эффективным решением будет использовать пакет atomic.

func deposit(acc *account, amount int64) {
   atomic.AddInt64(&acc.balance, amount)
}

*/

func main() {
	acc := account{balance: 0}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			deposit(&acc, 1)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("balance=%d\n", acc.balance)

	/*
	Решив Data Race через синхронизацию доступа к памяти (блокировки) не всегда решается race condition и logical correctness.
	*/

}