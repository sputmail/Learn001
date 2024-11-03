package main

import (
	"fmt"
	"io"
	"os"
)

type timeslot struct {
	open   int8
	max    int8
	min    int8
	volume uint8
	open1  int8
	open2  int8
	open3  int8
	answer int16
}

func loadtimeslots() []timeslot {
	file, err := os.OpenFile("sbcandlesM5_out.txt", os.O_RDONLY, 0666) // создаем файл
	if err != nil {                                                    // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	var ts timeslot
	timeslots := make([]timeslot, 0)

	defer file.Close() // закрываем файл
	for {

		_, err = fmt.Fscanf(file, "%d %d %d %d %d %d %d %d\n", &ts.open, &ts.max, &ts.min, &ts.volume, &ts.open1, &ts.open2, &ts.open3, &ts.answer)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		timeslots = append(timeslots, ts)
		//cnt := len(timeslots) - 1
		//fmt.Printf("%d %d %d %d\n", timeslots[cnt].open, timeslots[cnt].max, timeslots[cnt].min, timeslots[cnt].volume)

	}
	fmt.Printf("%d is load\n", len(timeslots))
	return timeslots
}
