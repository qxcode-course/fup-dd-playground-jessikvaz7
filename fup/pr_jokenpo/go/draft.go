package main
import "fmt"
func main() {
    const pedra int = 0 
    const papel int = 1
    const Tesoura int = 2

    var jog1, jog2 int
    var vit1, vit2 int

    for vit1 < 2 && vit2 < 2 {
        fmt.Println("placar vit1:", vit1, "vit2")


    fmt.Println("jog1:digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog1)
    fmt.Println("jog2:digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog2)
    
    if jog1 == jog2 {
        fmt.Println("empate")
    } else if (jog1 == pedra && jog2 == Tesoura) ||
    (jog1 == papel && jog2 == pedra) ||
    (jog1 == tesoura && jog2 == papel) {





}
