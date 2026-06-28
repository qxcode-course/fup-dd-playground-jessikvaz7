package main

import "fmt"

func main() {

    var troco float64

    fmt.Scan(&troco)
    
    vetor := []float64{100.00, 50.00, 20.0, 10.0, 5.0, 2.00, 1.00, 0.50, 0.25, 0.10, 0.5}

    cedulas := make([]int, len(vetor))

    for troco > 0 || troco > 0.5 {
    for i := 0; i < len(vetor); i++ {
    if(troco >= vetor[i]) {
      cedulas[i] = (int) (troco/(vetor[i]))
     } else {
        cedulas[i] = 0;

     }
   }  
  }
}