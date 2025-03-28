programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real n1,n2,n3,m
    escreva("Digite o primeiro número: ")
		leia(n1)
    escreva("Digite o segundo número: ")
		leia(n2)
    escreva("Digite o terceiro número: ")
		leia(n3)
    m = (n1+n2+n3)/3
    se (m>=6)
    escreva ( "MÉDIA=",mat.arredondar(m,2),"\n","APROVADO")
    se (m<6)
    escreva ("MÉDIA=",mat.arredondar(m,2),"\n","REPROVAD0")
  }
}
