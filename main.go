package main

import (
	"fmt"
)

func main(){
   
   var t int
   multiplier := []int{1,2,3,4,5,6,7,8,9,10,11,12}
   fmt.Print("Enter a number to see its times table: ")
  
  _, err := fmt.Scan(&t)  
   if err != nil {
	   fmt.Println("Invalid input:" , err)
	   return
   }
   for _, i := range(multiplier){
	   result := t * i
	   fmt.Printf("%d x %d = %d\n", t, i, result)
   }
}