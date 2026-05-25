package main

import "fmt"

func main() {

  var P, D1, D2 int
     fmt.Scanln(&P)
     fmt.Scanln(&D1)
     fmt.Scanln(&D2)

     soma := D1 + D2
     if soma % 2 == 0 {
     if P == 0 {
     fmt.Println(0)
} else {
     fmt.Println(1)
}
 } else {                                
 
     if P == 1 {
      fmt.Println(0)
 } else {
      fmt.Println(1)   
  }   
 }
}