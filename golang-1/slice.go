package main

import "fmt"

func main() {
	slice_1 := []string{"sate", "madura"}
	slice_1 = append(slice_1, "cuih")
	slice_1 = append(slice_1, "sate")
	
	fmt.Println(slice_1)
	fmt.Println(len(slice_1))
}