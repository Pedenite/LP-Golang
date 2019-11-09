package main

import "fmt"

type Palavra struct {
	palavraOriginal string
	palavraTraducao string
	anterior        *Palavra
	proximo         *Palavra
}

type ConjuntoPalavras struct {
	tamanho int
	inicio  *Palavra
	fim     *Palavra
}

func (palavras *ConjuntoPalavras) Append(novaPalavra *Palavra) {
	if palavras.tamanho == 0 {
		palavras.inicio = novaPalavra
		palavras.fim = novaPalavra
	} else {
		ultimaPalavra := palavras.fim
		ultimaPalavra.proximo = novaPalavra
		novaPalavra.anterior = ultimaPalavra
		palavras.fim = novaPalavra
	}
	palavras.tamanho++
}

func main() {
	conjuntoP := &ConjuntoPalavras{}
	p1 := Palavra{
		palavraOriginal: "Blue",
		palavraTraducao: "Azul",
	}
	conjuntoP.Append(&p1)

	fmt.Printf("Length: %v\n", conjuntoP.tamanho)
	fmt.Printf("First: %v\n", conjuntoP.inicio)

	// p2 := Post{
	//     body: "Dolor sit amet",
	// }
	// f.Append(&p2)

	// fmt.Printf("Length: %v\n", f.length)
	// fmt.Printf("First: %v\n", f.start)
	// fmt.Printf("Second: %v\n", f.start.next)
}
