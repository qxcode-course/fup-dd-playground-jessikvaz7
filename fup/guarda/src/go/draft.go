package main

import "fmt"

func main() {

    var wifi int
    var login int
    var admin int
    
    fmt.Scan(&wifi)
    fmt.Scan(&login)
    fmt.Scan(&admin)

  if !(wifi == 1){
    fmt.Println("you must connect to wifi")
  } else if !(login == 1) {
    fmt.Println("you need to login first")
  } else if !(admin == 1) {
    fmt.Println("you must to login as admin")
  } else {
    fmt.Println("done") 

  }

}