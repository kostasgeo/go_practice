// showcase of built0u-in make function that allocates a zeroed array and returns a slice that 
// referts to that array 
package making_slices 

import "fmt"

func Do(){
	//len(a)=5, cap(a)=5, all zeros
	a := make([]int, 5)
	printSlice("a", a)

	// len(b)=0, cap(b)=5, empty
	b := make([]int, 0, 5)
	printSlice("b", b)

	// len(c)=2 (b[0], b[1]), cap(c)=5
	c := b[:2]
	printSlice("c", c)

	// len(d)=3 (b[2], b[3], b[4] since there is enough capacity in c)
	// cap(d) = 3 (since start from 2 and reached full capacity) 
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)

}