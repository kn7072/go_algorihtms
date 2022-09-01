package main

import (
	"fmt"
)

//https://ubiklab.net/posts/data-protected-by-confinement/

/*
Первое, что мы должны сделать — это назначить владельца канала.

    Владелец канала — это горутина которая создает, пишет и закрывает канал.


Однонаправленные каналы — это инструмент, который позволит нам различать подпрограммы, которые владеют каналами, и те, которые только используют их.
Владельцы каналов имеют доступ для записи в канал (chan или chan<-), а утилиты канала имеют доступ только для чтения (<-chan).


Процедура, которой принадлежит канал, должна:
    Создать канал
    Передать канал потребителям
    Выполнить запись
    Закрыть канал

chanOwner := func() <-chan int {
   out := make(chan int, 5) // 1. create

   go func() {
      defer close(out) // 4. close
      for i := 0; i <= 5; i++ {
         out <- i // 3. write
      }
   }()
   return out // 2. return
}


Разрешая создавать, записывать и закрывать канал только владельцу мы получаем следующие преимущества
    Поскольку мы инициализируем канал, мы исключаем риск deadlock путем записи в нулевой канал
    Поскольку мы инициализируем канал, мы исключаем риск вызвать panic, закрывая нулевой канал
    Поскольку мы сами решаем, когда канал закроется, мы исключаем риск вызвать panic, записывая в закрытый канал
    Поскольку мы сами решаем, когда канал закроется, мы исключаем риск вызвать panic, закрывая канал более одного раза
    Мы используем средство проверки типов во время компиляции, чтобы предотвратить неправильную запись в наш канал

Потребитель канала должен беспокоиться о двух вещах
    Знание, когда канал закрыт
    Обработка блокировок по любой причине
Информацию о закрытии канала можно получить из второго значения при чтении из канала. Или воспользоваться конструкцией for range.

Второй пункт определить гораздо сложнее, поскольку он зависит от нашей стратегии. 
Можно установить тайм-аут, можно прекратить чтение, когда кто-то скажет нам, 
или мы просто можем блокировать информацию на протяжении всего жизненного цикла процесса. 
Важно то, что как потребитель мы должны учитывать тот факт, что чтение может и будет блокироваться.
*/


func main(){

    chanOwner := func() <-chan int {
        out := make(chan int, 5) // 1. create
     
        go func() {
           defer close(out) // 4. close
           for i := 0; i <= 5; i++ {
              out <- i // 3. write
           }
        }()
     
        return out // 2. return
    }
     
    /*
        Потребитель имеет доступ только к каналу чтения, и поэтому ему нужно знать только, 
        как он должен обрабатывать блокировку чтения и закрытие канала.
    */
    consumer := func(in <-chan int) {
        for result := range in {
           fmt.Printf("Received: %d\n", result)
        }
        fmt.Println("Done receiving!")
    }
     
    results := chanOwner()
    consumer(results)
    /*
        В контексте лексической области можно писать синхронный код. 
        Избегая большого количества проблем синхронизаций. 
        Таким образом получается более понятный и безопасный код.
    */
}