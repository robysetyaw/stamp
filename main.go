package main

import (
	"fmt"
	"stamp/stamp"
)

func main() {
	foobardata := stamp.Foobar(100)
	fmt.Println("")
	// fmt.Print(foobardata)
	for _, v := range foobardata {
		fmt.Print(v, " ")
	}
	fmt.Println("")
	fmt.Println("")

	data, err := stamp.FetchOpenWeatherMap()
	fmt.Println("Weather Forecast :")
	if err != nil {
		fmt.Print(err)
	} else {
		for _, v := range data {
			fmt.Println(v)
		}
	}
	fmt.Println("")
}
