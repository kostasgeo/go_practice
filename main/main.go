package main 

import "fmt"
import "github.com/kostasgeo/slice_pointers"
import "github.com/kostasgeo/slice_literals"
import "github.com/kostasgeo/slice_bounds"

func main() {
	fmt.Println("Go main Function")
	fmt.Println("slice_pointers")
	slice_pointers.Do()
	fmt.Println("slice_literals")
	slice_literals.Do()
	fmt.Println("slice_bounds")
	slice_bounds.Do()
}