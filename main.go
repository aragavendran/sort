package main

import "fmt"

func main() {

	data := []int{70, 40, 50, 10, 22, 33, 54, 77}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] > data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
	fmt.Println(data)
}
