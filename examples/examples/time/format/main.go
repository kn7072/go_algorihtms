package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	// layout := "Day: 02 Month: Jan Year: 2006 Hour: 24"
	// layout := "02 Jan 2006 15:04"
	// layout := "02 Jan 06 15:04"
	layout := "15:04"
	fmt.Println(label)
	str := fmt.Sprint(t.Format(layout))
	fmt.Println(str)

	tt, err := time.Parse(layout, str)
	if err == nil {
		fmt.Println(tt)
	}
}

func main() {
	t := time.Date(1995, time.September, 9, 10, 11, 12, 0, time.Local)
	PrintTime("t", &t)

	t1 := time.Unix(1433228090, 0)
	PrintTime("t1", &t1)

	t2 := time.Now()
	PrintTime("t2", &t2)
}