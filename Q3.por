programa {
  funcao inicio() {
    inteiro a,b,c
    real n1,n2,k,m
   escreva("Digite o primeiro número: ")
	 leia(a)
   escreva("Digite o segundo número: ")
	 leia(b)
   escreva("Digite o terceiro número: ")
	 leia(c)
   k= 100 * a
   m= 10 * b
   n1= k + m + c 
   n2= (n1 * n1)
   se (a>9 ou b>9 ou c>9)
   escreva ("DIGITO INVALIDO")
   senao escreva (n1,-n2)
  }
}
