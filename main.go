package main

import (
	"fmt"
	"stamp/stamp"
)

func main() {
	foobardata := stamp.Foobar(50)
	fmt.Print(foobardata)
	fmt.Println("")
	data, err := stamp.FetchOpenWeatherMap()
	if err != nil {
		fmt.Print(err)
	} else {
		for _, v := range data {
			fmt.Println(v)
		}
	}
}
