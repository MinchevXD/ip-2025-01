programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
   real a,b,c,d
   escreva("Digite o primeiro número: ")
		leia(a)
    escreva("Digite o segundo número: ")
		leia(b)
    escreva("Digite o terceiro número: ")
		leia(c)
    d=b*b-(4*a*c)
    escreva ( "O VALOR DE DELTA E =",mat.arredondar(delta,2),"\n")
  }
}
