package main

import (
	"bufio"
	"fmt"
	"strings"
	"unsafe"
	"reflect"
	"os"

)

/*
Если вы пишите на Go, то считывать данные рекомендуется, используя bufio.NewReader(os.Stdin). 
Иначе при больших входных данных могут быть проблемы с каким-то внутренним буфером в Go и данные не будут вычитаны полностью. 
Например, зачастую можно делать вот так: in := bufio.NewReader(os.Stdin) и потом fmt.Fscan(in, &T)

А вот если делать просто fmt.Scan(&t), то при больших входных данных могут быть проблемы.
*/


const (
	path = "./data_write"
)

func main() {

	file, err := os.OpenFile(path, os.O_CREATE | os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println(bufio.MaxScanTokenSize)

	//input := strings.Repeat("foo\nbar\nbaz", 3) //"foo\nbar\nbaz", 3
	input := strings.Repeat("12345 ", 1000000) + "\n" + "777" +  "\n" + "888" + "\n" 
	fmt.Println(unsafe.Sizeof(input))
	fmt.Println(reflect.TypeOf(input).Size())

	scanner := bufio.NewScanner(strings.NewReader(input))
	size := 32 * 1e6
	buf := make([]byte, 0, int(size))
	scanner.Buffer(buf, int(size))

	//scanner.Split(bufio.ScanWords)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(text)
		count := len(strings.Split(text, " "))
		fmt.Println(count)
		file.WriteString(text)
		file.WriteString("\n")
	}
}