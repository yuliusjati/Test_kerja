package main

import (
	"sort"
	"fmt"
)

func main(){
	number := [...]int{1,4,5,6,8,2}

	for i := 0; i < 6; i++ {
		
		fmt.Print(number[i])
		for j := 0; j < number[i]; j++ {
			
			fmt.Print("|")
			
		}
		fmt.Println("")
	
	}
	fmt.Println("=++++++++++++++++++++++++=")

	arrslice := number[:]
	sort.Sort(sort.IntSlice(arrslice))

	for i := 0; i < 6; i++ {
		bil := number[i]
		fmt.Print(bil)
		for j := 0; j < number[i]; j++ {
			fmt.Print("|")			
		}		
		fmt.Println("")
	} 
}