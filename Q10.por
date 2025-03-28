programa {inclua biblioteca Matematica --> mat
  funcao inicio() {
    real a,b,c,d,determinante
     escreva("Digite o primeiro número: ")
		leia(a)
    escreva("Digite o segundo número: ")
		leia(b)
    escreva("Digite o terceiro número: ")
		leia(c)
    escreva("Digite o quarto número: ")
		leia(d)
    determinante=a*d-b*c
    escreva ("O VALOR DO DETERMINANTE E=",mat.arredondar(determinante,2) )
  }
}
