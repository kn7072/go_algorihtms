go test -bench=. . все в текущей директории
go test -bench=Fast . 
go test -bench=. -benchmem . потребление памяти на операцию и алокации
go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out .

go test -bench=. -mutexprofile mutex.out

[Документация по флагам](https://golang.org/cmd/go/#hdr-Testing_flags) (go help testflag)