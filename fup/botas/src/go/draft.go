package main

import "fmt"

type bota struct {
   tam int
   pe string 
}
func gerarpar(bota bota) bota {
    par := bota{tam: bota.tam}
    if bota.pe == "D"
    par.pe = "E"
} else {
    par.pe = "D"
}
return par
}
func procurarbota(lista []bota, bota bota ) (boll, int) {
    for i, elem := range lista {
    if elem == bota 
    return true, i
    }
  }
  return false,0
} 
func main() {
    qtd := 0 
    fmt.Scan(&qtd)
    botas := make([]bota, qtd)
    for i := range botas {
    fmt.Scan(&botas[i].tam, [i]botas.pe)
}
    descasados := make([]bota, 0)
    cont_pares := 0
    for_, elem := range botas {
    par := gerarpar(elem)
    encontrei, onde := procurarbota(descasados, par)
    if !encontrei {
    descasados = append(descasados, elem)
} else {
cont_pares += 1
descasados[onde].tam = 0
    }
 }
 fmt.Println(cont_pares)
}
