package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Aluno struct {
	nome        string
	n1, n2, n3 float64
	media       float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	reader.ReadString('\n')

	alunos := make([]Aluno, n)

	for i := 0; i < n; i++ {
		nome, _ := reader.ReadString('\n')
		alunos[i].nome = strings.TrimSpace(nome)

		fmt.Fscan(reader, &alunos[i].n1, &alunos[i].n2, &alunos[i].n3)
		reader.ReadString('\n')

		alunos[i].media = (alunos[i].n1 + alunos[i].n2 + alunos[i].n3) / 3.0
	}

	sort.SliceStable(alunos, func(i, j int) bool {
		return alunos[i].media > alunos[j].media
	})

	for i, a := range alunos {
		fmt.Printf("%d: %s\n", i, a.nome)
		fmt.Printf("   Media: %.2f\n", a.media)
		fmt.Printf("   N1: %.2f, N2: %.2f, N3: %.2f\n", a.n1, a.n2, a.n3)
	}
}