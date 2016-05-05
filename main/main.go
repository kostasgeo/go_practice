package main 

import "fmt"
import "github.com/kostasgeo/slice_pointers"
import "github.com/kostasgeo/slice_literals"

func main() {
	fmt.Println("Go main Function")
	fmt.Println("slice_pointers")
	slice_pointers.Do()
	fmt.Println("slice_literals")
	slice_literals.Do()
}